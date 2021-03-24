package endpoint

import (
	"context"
	"github.com/feiliangwang/audioAnalysis/api"
	"github.com/feiliangwang/audioAnalysis/communicate"
	"github.com/feiliangwang/audioAnalysis/orm"
	"github.com/feiliangwang/audioAnalysis/service"
	"github.com/go-kit/kit/endpoint"
	"log"
)

/**
 * @Author: feiliang.wang
 * @Description: 音频服务端点
 * @File:  audio
 * @Version: 1.0.0
 * @Date: 2021/3/23 23:23
 */

type AudioServiceEndpoints struct {
	AddEndpoint    endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
	UpdateEndpoint endpoint.Endpoint
	DetailEndpoint endpoint.Endpoint
	ListEndpoint   endpoint.Endpoint
}

func MakeAudioAddEndpoint(svc service.AudioService, logger *log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if r, ok := request.(api.AduioInfoAddRequest); !ok {
			return nil, ErrorBadRequest
		} else if data, err := svc.Add(ctx, logger, r); err != nil {
			return nil, err
		} else {
			return communicate.DefaultDataResponseDao{Data: data}, nil
		}
	}
}

func MakeAudioDeleteEndpoint(svc service.AudioService, logger *log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if r, ok := request.(api.BaseAnchorRequest); !ok {
			return nil, ErrorBadRequest
		} else if data, err := svc.Delete(ctx, logger, r.Id); err != nil {
			return nil, err
		} else {
			return communicate.DefaultDataResponseDao{Data: data}, nil
		}
	}
}

func MakeAudioUpdateEndpoint(svc service.AudioService, logger *log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if r, ok := request.(api.AduioInfoUpdateRequest); !ok {
			return nil, ErrorBadRequest
		} else if data, err := svc.Update(ctx, logger, r); err != nil {
			return nil, err
		} else {
			return communicate.DefaultDataResponseDao{Data: data}, nil
		}
	}
}

func MakeAudioDetailEndpoint(svc service.AudioService, logger *log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if r, ok := request.(api.BaseAnchorRequest); !ok {
			return nil, ErrorBadRequest
		} else if data, err := svc.Detail(ctx, logger, r.Id); err != nil {
			return nil, err
		} else {
			return communicate.DefaultDataResponseDao{Data: data}, nil
		}
	}
}

func MakeAudioListEndpoint(svc service.AudioService, logger *log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if r, ok := request.(api.AduioInfoListRequest); !ok {
			return nil, ErrorBadRequest
		} else if list, pageResp, err := svc.List(ctx, logger, r.PageRequest, orm.MakeSqlFilterMap(request)); err != nil {
			return nil, err
		} else {
			return communicate.DefaultPageListResponseDao{
				PageResponseDao:        pageResp,
				DefaultListResponseDao: communicate.DefaultListResponseDao{List: list}}, nil
		}
	}
}
