package transport

import (
	"context"
	"fmt"
	"github.com/feiliangwang/audioAnalysis/api"
	"github.com/feiliangwang/audioAnalysis/communicate"
	"github.com/feiliangwang/audioAnalysis/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

/**
 * @Author: feiliang.wang
 * @Description: 音频服务传输点
 * @File:  audio
 * @Version: 1.0.0
 * @Date: 2021/3/23 23:41
 */

func MakeAudioHttpHandler(ctx context.Context, logger *log.Logger, router *mux.Router, endpoints endpoint.AudioServiceEndpoints) {
	options := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(communicate.EncodeError),
	}
	r := router.PathPrefix("/audio").Subrouter()
	r.Methods("POST").Path("/add").Handler(kithttp.NewServer(
		endpoints.AddEndpoint,
		func(ctx context.Context, r *http.Request) (request interface{}, err error) {
			data := api.AduioInfoAddRequest{}
			r.ParseForm() //解析表单
			r.ParseMultipartForm(32 << 20)
			checkString := func(key string, set *string) error {
				if t := r.Form.Get(key); t == "" {
					return fmt.Errorf("get key %s is empty", key)
				} else {
					*set = t
					return nil
				}
			}
			checkInt := func(key string, set *int64) error {
				if t, err := strconv.ParseInt(r.Form.Get(key), 10, 64); err != nil {
					return fmt.Errorf("get key %s value %s is invalid int", key, r.Form.Get(key))
				} else {
					*set = t
					return nil
				}
			}
			if err := checkString("TelephoneKind", &data.TelephoneKind); err != nil {
				return data, err
			} else if err := checkString("FromTelephoneNumber", &data.FromTelephoneNumber); err != nil {
				return data, err
			} else if err := checkString("ToTelephoneNumber", &data.ToTelephoneNumber); err != nil {
				return data, err
			} else if err := checkInt("HappenTimestamp", &data.HappenTimestamp); err != nil {
				return data, err
			} else if err := checkInt("TotalDuration", &data.TotalDuration); err != nil {
				return data, err
			} else if f, h, err := r.FormFile("File"); err != nil {
				return data, fmt.Errorf("get File key fail.+%v", err)
			} else {
				data.FileName = h.Filename
				data.Stream = f
				return data, nil
			}

		},
		communicate.EncodeResponse,
		options...,
	))
	r.Methods("POST").Path("/delete").Handler(kithttp.NewServer(
		endpoints.DeleteEndpoint,
		communicate.DecodePostRequest(reflect.TypeOf(api.BaseAnchorRequest{})),
		communicate.EncodeResponse,
		options...,
	))
	r.Methods("POST").Path("/update").Handler(kithttp.NewServer(
		endpoints.UpdateEndpoint,
		communicate.DecodePostRequest(reflect.TypeOf(api.AduioInfoUpdateRequest{})),
		communicate.EncodeResponse,
		options...,
	))
	r.Methods("POST").Path("/detail").Handler(kithttp.NewServer(
		endpoints.DetailEndpoint,
		communicate.DecodePostRequest(reflect.TypeOf(api.BaseAnchorRequest{})),
		communicate.EncodeResponse,
		options...,
	))
	r.Methods("POST").Path("/list").Handler(kithttp.NewServer(
		endpoints.ListEndpoint,
		communicate.DecodePostRequest(reflect.TypeOf(api.AduioInfoListRequest{})),
		communicate.EncodeResponse,
		options...,
	))
}
