package communicate

import (
	"encoding/json"
)

/**
 * @Author: feiliang.wang
 * @Description: 通用响应定义
 * @File:  response
 * @Version: 1.0.0
 * @Date: 2020/8/1 11:25 上午
 */

//基础响应数据
type BaseResponseDao struct {
	ErrorNo  int32            `json:"error_no"`
	ErrorMsg string           `json:"error_msg"`
	Data     json.RawMessage  `json:"data,omitempty"`
	Page     *PageResponseDao `json:"page,omitempty"`
	List     json.RawMessage  `json:"list,omitempty"`
}

//分页响应数据
type PageResponseDao struct {
	Page      int `json:"page"`
	Size      int `json:"size"`
	TotalPage int `json:"totalPage"`
	TotalSize int `json:"totalSize"`
}

func (d PageResponseDao) GetPage() *PageResponseDao {
	return &d
}

type PageResponse interface {
	GetPage() *PageResponseDao
}

type DataResponse interface {
	GetData() interface{}
}

type ListResponse interface {
	GetList() interface{}
}

type DefaultDataResponseDao struct {
	Data interface{} `json:"data"`
}

func (d DefaultDataResponseDao) GetData() interface{} {
	return d.Data
}

type DefaultListResponseDao struct {
	List interface{} `json:"list"`
}

func (d DefaultListResponseDao) GetList() interface{} {
	return d.List
}

type DefaultDataListResponseDao struct {
	DefaultDataResponseDao
	DefaultListResponseDao
}

type DefaultPageListResponseDao struct {
	PageResponseDao
	DefaultListResponseDao
}

type DefaultDataPageListResponseDao struct {
	DefaultDataResponseDao
	PageResponseDao
	DefaultListResponseDao
}
