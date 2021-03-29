package api

import (
	"github.com/feiliangwang/audioAnalysis/communicate"
	"github.com/feiliangwang/audioAnalysis/orm"
	"io"
)

/**
 * @Author: feiliang.wang
 * @Description: 音频服务API请求数据结构定义
 * @File:  aduio
 * @Version: 1.0.0
 * @Date: 2021/3/23 22:29
 */

//音频信息添加
type AduioInfoAddRequest struct {
	//手机型号
	TelephoneKind string `json:"telephoneKind"`
	//账户类型
	FromTelephoneNumber string `json:"fromTelephoneNumber"`
	//本位币
	ToTelephoneNumber string `json:"toTelephoneNumber"`
	//发生时间
	HappenTimestamp int64 `json:"happenTimestamp"`
	//时长
	TotalDuration int64 `json:"totalDuration"`
	//文件名
	FileName string `json:"fileName"`
	//文件流
	Stream io.ReadCloser `json:"-"`
}

type AduioInfoUpdateRequest struct {
	BaseAnchorRequest
	//手机型号
	TelephoneKind *string `json:"telephoneKind"`
	//账户类型
	FromTelephoneNumber *string `json:"fromTelephoneNumber"`
	//本位币
	ToTelephoneNumber *string `json:"toTelephoneNumber"`
	//发生时间
	HappenTimestamp *int64 `json:"happenTimestamp"`
	//时长
	TotalDuration *int64 `json:"totalDuration"`
	////文件名
	//FileName *string `json:"fileName"`
}

//音频信息列表查询请求
type AduioInfoListRequest struct {
	communicate.PageRequest
	//手机型号
	TelephoneKind *orm.StringFilter `json:"telephoneKind"`
	//账户类型
	FromTelephoneNumber *orm.StringFilter `json:"fromTelephoneNumber"`
	//本位币
	ToTelephoneNumber *orm.StringFilter `json:"toTelephoneNumber"`
	//发生时间
	HappenTimestamp *orm.Int64Filter `json:"happenTimestamp"`
	//时长
	TotalDuration *orm.Int64Filter `json:"totalDuration"`
	//文件名
	FileName *orm.StringFilter `json:"fileName"`
	//文件大小
	Size *orm.Int64Filter `json:"size"`
	//文件类型
	FileType *orm.StringFilter `json:"fileType"`
	//文件MD5值
	Md5 *orm.StringFilter `json:"md5"`
	//文件下载路径
	FilePath *orm.StringFilter `json:"filePath"`
}
