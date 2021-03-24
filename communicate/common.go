package communicate

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

/**
 * @Author: feiliang.wang
 * @Description: 通用的kit序列化定义
 * @File:  common
 * @Version: 1.0.0
 * @Date: 2020/8/7 4:26 下午
 */

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//w.WriteHeader(http.StatusInternalServerError)
	switch data := err.(type) {
	case CodeError:
		if err := json.NewEncoder(w).Encode(&BaseResponseDao{ErrorNo: data.Code(), ErrorMsg: data.Error()}); err != nil {
			log.Println(err)
		}
	default:
		if err := json.NewEncoder(w).Encode(&BaseResponseDao{ErrorNo: 999, ErrorMsg: data.Error()}); err != nil {
			log.Println(err)
		}
	}
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	base := &BaseResponseDao{}
	switch data := response.(type) {
	case PageResponse:
		base.Page = data.GetPage()
	}
	switch data := response.(type) {
	case DataResponse:
		bs, _ := json.Marshal(data.GetData())
		base.Data = bs
	}
	switch data := response.(type) {
	case ListResponse:
		bs, _ := json.Marshal(data.GetList())
		base.List = bs
	}
	return json.NewEncoder(w).Encode(base)
}

func DecodePostRequest(t reflect.Type) func(ctx context.Context, r *http.Request) (interface{}, error) {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		v := reflect.New(t)
		if r.ContentLength != 0 {
			if err := json.NewDecoder(r.Body).Decode(v.Interface()); err != nil {
				return nil, err
			}
		}
		d := v.Elem().Interface()
		return d, nil
	}
}

func DecodeGetRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}
