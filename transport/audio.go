package transport

import (
	"context"
	"github.com/feiliangwang/audioAnalysis/api"
	"github.com/feiliangwang/audioAnalysis/communicate"
	"github.com/feiliangwang/audioAnalysis/endpoint"
	"github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"log"
	"reflect"
)

/**
 * @Author: feiliang.wang
 * @Description: 音频服务传输点
 * @File:  audio
 * @Version: 1.0.0
 * @Date: 2021/3/23 23:41
 */

func MakeAudioHttpHandler(ctx context.Context, logger *log.Logger, router *mux.Router, endpoints endpoint.AudioServiceEndpoints) {
	options := []http.ServerOption{
		http.ServerErrorEncoder(communicate.EncodeError),
	}
	r := router.PathPrefix("/audio").Subrouter()
	r.Methods("POST").Path("/add").Handler(http.NewServer(
		endpoints.AddEndpoint,
		communicate.DecodePostRequest(reflect.TypeOf(api.AduioInfoAddRequest{})),
		communicate.EncodeResponse,
		options...,
	))
	r.Methods("POST").Path("/delete").Handler(http.NewServer(
		endpoints.DeleteEndpoint,
		communicate.DecodePostRequest(reflect.TypeOf(api.BaseAnchorRequest{})),
		communicate.EncodeResponse,
		options...,
	))
	r.Methods("POST").Path("/update").Handler(http.NewServer(
		endpoints.UpdateEndpoint,
		communicate.DecodePostRequest(reflect.TypeOf(api.AduioInfoUpdateRequest{})),
		communicate.EncodeResponse,
		options...,
	))
	r.Methods("POST").Path("/detail").Handler(http.NewServer(
		endpoints.DetailEndpoint,
		communicate.DecodePostRequest(reflect.TypeOf(api.BaseAnchorRequest{})),
		communicate.EncodeResponse,
		options...,
	))
	r.Methods("POST").Path("/list").Handler(http.NewServer(
		endpoints.ListEndpoint,
		communicate.DecodePostRequest(reflect.TypeOf(api.AduioInfoListRequest{})),
		communicate.EncodeResponse,
		options...,
	))
}
