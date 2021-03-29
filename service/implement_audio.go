package service

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/feiliangwang/audioAnalysis/api"
	"github.com/feiliangwang/audioAnalysis/communicate"
	"github.com/feiliangwang/audioAnalysis/dao"
	"github.com/feiliangwang/audioAnalysis/orm"
	"io"
	"log"
	"os"
	"path"
	"time"
)

/**
 * @Author: feiliang.wang
 * @Description: 音频文件服务实现
 * @File:  audio
 * @Version: 1.0.0
 * @Date: 2021/3/23 23:11
 */

type AudioServerImplement struct {
	db  *sql.DB
	dir string
}

func NewAudioServer(db *sql.DB, dir string) AudioService {
	return &AudioServerImplement{db: db, dir: dir}
}

const Audio_Table_Name = "audio_file"

/**
 * @Author feiliang.wang
 * @Description 添加音频文件信息
 * @Date 22:38 2021/3/23
 * @Param ctx 上下文
 * @Param log 日志
 * @Param request 请求数据
 * @Param closer 文件内容流
 * @return
 **/
func (s *AudioServerImplement) Add(ctx context.Context, logger *log.Logger, request api.AduioInfoAddRequest) (id int32, err error) {
	relativePath := path.Join("/", request.FromTelephoneNumber, request.ToTelephoneNumber, time.Unix(request.HappenTimestamp, 0).Format("20060102"), request.FileName)
	dir := path.Join(s.dir, request.FromTelephoneNumber, request.ToTelephoneNumber, time.Unix(request.HappenTimestamp, 0).Format("20060102"))
	file := path.Join(dir, request.FileName)
	hash := md5.New()
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return 0, err
	} else if f, err := os.Create(file); err != nil {
		return 0, err
	} else {
		defer f.Close()
		if size, err := io.Copy(f, request.Stream); err != nil {
			return 0, err
		} else if _, err := io.Copy(hash, f); err != nil {
			return 0, err
		} else {
			return orm.InsertInfo(s.db, Audio_Table_Name, &dao.AduioDao{
				Id:                  0,
				TelephoneKind:       request.TelephoneKind,
				FromTelephoneNumber: request.FromTelephoneNumber,
				ToTelephoneNumber:   request.ToTelephoneNumber,
				HappenTimestamp:     request.HappenTimestamp,
				TotalDuration:       request.TotalDuration,
				FileName:            request.FileName,
				Size:                size,
				FileType:            path.Ext(file),
				Md5:                 hex.EncodeToString(hash.Sum(nil)),
				FilePath:            relativePath,
			})
		}
	}

}

/**
 * @Author feiliang.wang
 * @Description 删除音频文件信息
 * @Date 23:27 2021/3/23
 * @Param ctx 上下文
 * @Param log 日志
 * @Param id 文件ID
 * @return
 **/
func (s *AudioServerImplement) Delete(ctx context.Context, logger *log.Logger, id int32) (ok bool, err error) {
	return orm.DeleteInfo(s.db, Audio_Table_Name, id)
}

/**
 * @Author feiliang.wang
 * @Description 更新音频文件信息
 * @Date 23:29 2021/3/23
 * @Param ctx 上下文
 * @Param log 日志
 * @Param request 请求数据
 * @return
 **/
func (s *AudioServerImplement) Update(ctx context.Context, logger *log.Logger, request api.AduioInfoUpdateRequest) (ok bool, err error) {
	m := make(map[string]interface{})
	if request.TelephoneKind != nil {
		m["TelephoneKind"] = *request.TelephoneKind
	}
	if request.FromTelephoneNumber != nil {
		m["FromTelephoneNumber"] = *request.FromTelephoneNumber
	}
	if request.ToTelephoneNumber != nil {
		m["ToTelephoneNumber"] = *request.ToTelephoneNumber
	}
	if request.HappenTimestamp != nil {
		m["HappenTimestamp"] = *request.HappenTimestamp
	}
	if request.TotalDuration != nil {
		m["TotalDuration"] = *request.TotalDuration
	}
	if len(m) == 0 {
		return false, fmt.Errorf("no data to update.")
	} else {

		return orm.UpdateInfo(s.db, Audio_Table_Name, request.Id, map[string]interface{}{
			"TelephoneKind":       request.TelephoneKind,
			"FromTelephoneNumber": request.FromTelephoneNumber,
			"ToTelephoneNumber":   request.ToTelephoneNumber,
			"HappenTimestamp":     request.HappenTimestamp,
			"TotalDuration":       request.TotalDuration,
		})
	}
}

/**
 * @Author feiliang.wang
 * @Description 查询音频文件明细
 * @Date 22:45 2021/3/23
 * @Param ctx 上下文
 * @Param log 日志
 * @Param id 数据ID
 * @return
 **/
func (s *AudioServerImplement) Detail(ctx context.Context, logger *log.Logger, id int32) (data dao.AduioDao, err error) {
	err = orm.GetInfo(s.db, Audio_Table_Name, id, &data)
	return
}

/**
 * @Author feiliang.wang
 * @Description 分页查询
 * @Date 22:55 2021/3/23
 * @Param ctx 上下文
 * @Param log 日志
 * @Param page 分页参数
 * @return filters 过滤参数
 **/
func (s *AudioServerImplement) List(ctx context.Context, logger *log.Logger, page communicate.PageRequest, filters orm.SqlFilterMap) (list []dao.AduioDao, pageResp communicate.PageResponseDao, err error) {
	var count int
	count, err = orm.GetCount(s.db, Audio_Table_Name, filters)
	if err != nil {
		return
	}
	maxPageNum := page.Verification(count)
	pageResp.Page = page.PageNumber
	pageResp.TotalPage = maxPageNum
	pageResp.TotalSize = count
	pageResp.Size, err = orm.ListInfo(s.db, logger, Audio_Table_Name, int(page.SortAsc), page.SortField, page.PageNumber, page.PageSize, filters, list)
	return
}
