package logging

import (
	"context"
	"github.com/feiliangwang/audioAnalysis/api"
	"github.com/feiliangwang/audioAnalysis/communicate"
	"github.com/feiliangwang/audioAnalysis/dao"
	"github.com/feiliangwang/audioAnalysis/orm"
	"github.com/feiliangwang/audioAnalysis/service"
	"log"
	"time"
)

/**
 * @Author: feiliang.wang
 * @Description: 音频服务日志中间件
 * @File:  audio
 * @Version: 1.0.0
 * @Date: 2021/3/29 21:05
 */
type skAppLoggingAudioServiceMiddleware struct {
	Service service.AudioService
	logger  *log.Logger
	Name    string
}

func SkAppLoggingAudioServiceMiddleware(logger *log.Logger) service.AudioServiceMiddleware {
	return func(next service.AudioService) service.AudioService {
		return skAppLoggingAudioServiceMiddleware{next, logger, "AudioService"}
	}
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
func (m skAppLoggingAudioServiceMiddleware) Add(ctx context.Context, logger *log.Logger, request api.AduioInfoAddRequest) (id int32, err error) {
	defer func(begin time.Time) {
		m.logger.Printf("service:%s,function:%s,request:%v,result:%v,took:%s\n", m.Name, "Add", []interface{}{request}, []interface{}{id, err}, time.Since(begin))
	}(time.Now())
	id, err = m.Service.Add(ctx, logger, request)
	return
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
func (m skAppLoggingAudioServiceMiddleware) Delete(ctx context.Context, logger *log.Logger, id int32) (ok bool, err error) {
	defer func(begin time.Time) {
		m.logger.Printf("service:%s,function:%s,request:%v,result:%v,took:%s\n", m.Name, "Delete", []interface{}{id}, []interface{}{ok, err}, time.Since(begin))
	}(time.Now())
	ok, err = m.Service.Delete(ctx, logger, id)
	return
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
func (m skAppLoggingAudioServiceMiddleware) Update(ctx context.Context, logger *log.Logger, request api.AduioInfoUpdateRequest) (ok bool, err error) {
	defer func(begin time.Time) {
		m.logger.Printf("service:%s,function:%s,request:%v,result:%v,took:%s\n", m.Name, "Update", []interface{}{request}, []interface{}{ok, err}, time.Since(begin))
	}(time.Now())
	ok, err = m.Service.Update(ctx, logger, request)
	return
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
func (m skAppLoggingAudioServiceMiddleware) Detail(ctx context.Context, logger *log.Logger, id int32) (data dao.AduioDao, err error) {
	defer func(begin time.Time) {
		m.logger.Printf("service:%s,function:%s,request:%v,result:%v,took:%s\n", m.Name, "Detail", []interface{}{id}, []interface{}{data, err}, time.Since(begin))
	}(time.Now())
	data, err = m.Service.Detail(ctx, logger, id)
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
func (m skAppLoggingAudioServiceMiddleware) List(ctx context.Context, logger *log.Logger, page communicate.PageRequest, filters orm.SqlFilterMap) (list []dao.AduioDao, pageResp communicate.PageResponseDao, err error) {
	defer func(begin time.Time) {
		m.logger.Printf("service:%s,function:%s,request:%v,result:%v,took:%s\n", m.Name, "List", []interface{}{page, filters}, []interface{}{list, pageResp, err}, time.Since(begin))
	}(time.Now())
	list, pageResp, err = m.Service.List(ctx, logger, page, filters)
	return
}
