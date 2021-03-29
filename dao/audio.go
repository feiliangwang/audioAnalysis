package dao

/**
 * @Author: feiliang.wang
 * @Description: 音频文件数据
 * @File:  aduio
 * @Version: 1.0.0
 * @Date: 2021/3/23 22:42
 */

type AduioDao struct {
	//文件ID
	Id int64 `json:"id"`
	//手机型号
	TelephoneKind string `json:"telephoneKind"`
	//拨打手机号
	FromTelephoneNumber string `json:"fromTelephoneNumber"`
	//接收手机号
	ToTelephoneNumber string `json:"toTelephoneNumber"`
	//发生时间
	HappenTimestamp int64 `json:"happenTimestamp"`
	//时长
	TotalDuration int64 `json:"totalDuration"`
	//文件名
	FileName string `json:"fileName"`
	//文件大小
	Size int64 `json:"size"`
	//文件类型
	FileType string `json:"fileType"`
	//文件MD5值
	Md5 string `json:"md5"`
	//文件下载路径
	FilePath string `json:"filePath"`
}
