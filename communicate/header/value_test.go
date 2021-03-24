package header

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

/**
 * @Author: feiliang.wang
 * @Description: //TODO
 * @File:  value_test
 * @Version: 1.0.0
 * @Date: 2020/9/1 1:48 下午
 */

func TestSetHttpHeaderValueToData(t *testing.T) {
	type Base struct {
		Name string
	}
	type Item struct {
		Base
		Value string
	}
	hv := StringHeaderValue{
		BaseHeaderValue: BaseHeaderValue{
			FieldName: "Base.Name",
			HeadKey:   "name",
		},
	}
	var data Item
	ty := reflect.TypeOf(data)
	vu := reflect.New(ty)
	r := &http.Request{}
	r.Header = make(http.Header)
	r.Header.Set("name", "abc")
	r.Header.Set("value", "123")
	err := SetHttpHeaderValueToData(hv, vu, r)
	if err != nil {
		t.Fatal(err)
	}

	hv = StringHeaderValue{
		BaseHeaderValue: BaseHeaderValue{
			FieldName: "Value",
			HeadKey:   "value",
		},
	}
	err = SetHttpHeaderValueToData(hv, vu, r)
	if err != nil {
		t.Fatal(err)
	}
	data = vu.Elem().Interface().(Item)
	if data.Name != "abc" {
		t.Fatal()
	}
	if data.Value != "123" {
		t.Fatal()
	}
	t.Log(data)
}

func TestDecodeHeaderRequest(t *testing.T) {
	type Base struct {
		Name string
	}
	type Item struct {
		Base
		Value string
	}
	values := []HeaderValue{StringHeaderValue{
		BaseHeaderValue: BaseHeaderValue{
			FieldName: "Base.Name",
			HeadKey:   "name",
		},
	},
		StringHeaderValue{
			BaseHeaderValue: BaseHeaderValue{
				FieldName: "Value",
				HeadKey:   "value",
			}},
	}
	r := &http.Request{}
	r.Header = make(http.Header)
	r.Header.Set("name", "abc")
	r.Header.Set("value", "123")

	//r.Body =  ioutil.NopCloser(bytes.NewBuffer([]byte("{}")))
	data, err := DecodeHeaderRequest(reflect.TypeOf(Item{}), values)(context.TODO(), r)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestDecodeHeaderPostRequest(t *testing.T) {
	type Base struct {
		Name         string
		BaseIntValue int
	}
	type Item struct {
		Base
		Value    string
		IntValue int
		Id       int32 `json:"id"`
	}
	values := []HeaderValue{StringHeaderValue{
		BaseHeaderValue: BaseHeaderValue{
			FieldName: "Base.Name",
			HeadKey:   "aname",
		},
	},
		StringHeaderValue{
			BaseHeaderValue: BaseHeaderValue{
				FieldName: "Value",
				HeadKey:   "avalue",
			}},
	}
	r := &http.Request{}
	r.Header = make(http.Header)
	r.Header.Set("aname", "abc")
	r.Header.Set("avalue", "jhg")
	r.Header.Set("intValue", "567")
	r.Header.Set("baseIntValue", "789")

	r.Body = ioutil.NopCloser(bytes.NewBuffer([]byte("{\"id\":1}")))
	r.ContentLength = 10
	data, err := DecodeHeaderPostRequest(reflect.TypeOf(Item{}), values...)(context.TODO(), r)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}
