package service

import (
	"context"
	"database/sql"
	"github.com/feiliangwang/audioAnalysis/api"
	"github.com/feiliangwang/audioAnalysis/communicate"
	"github.com/feiliangwang/audioAnalysis/dao"
	"github.com/feiliangwang/audioAnalysis/orm"
	"log"
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
	//todo 计算一些数据 然后添加到数据库 保存文件
	return 0, nil
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
	//todo 删除文件
	return false, nil
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
	//todo 更新信息
	return false, nil
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
	//todo 数据库中查询文件
	return dao.AduioDao{}, nil
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
	//todo 查询文件列表
	return nil, communicate.PageResponseDao{}, nil
}
