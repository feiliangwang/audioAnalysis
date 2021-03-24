package header

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

/**
 * @Author: feiliang.wang
 * @Description: HTTP头获取数据定义
 * @File:  value
 * @Version: 1.0.0
 * @Date: 2020/9/1 10:51 上午
 */

var FieldError = fmt.Errorf("field err")

/**
 * @Author feiliang.wang
 * @Description 定义获取HTTP头数据接口
 * @Date 11:12 上午 2020/9/1
 **/
type HeaderValue interface {
	/**
	 * @Author feiliang.wang
	 * @Description 获取字段名
	 * @Date 11:12 上午 2020/9/1
	 * @return
	 **/
	GetFieldName() string
	/**
	 * @Author feiliang.wang
	 * @Description 获取HTTP头上键值
	 * @Date 11:13 上午 2020/9/1
	 * @return
	 **/
	GetHeadKey() string
	/**
	 * @Author feiliang.wang
	 * @Description 设置头值到字段中
	 * @Date 11:14 上午 2020/9/1
	 * @Param field 反射的字段值
	 * @Param value HTTP头内的值
	 * @return
	 **/
	SetValue(field reflect.Value, value string) error
}

/**
 * @Author feiliang.wang
 * @Description 设置HTTP头值到结构体内
 * @Date 11:15 上午 2020/9/1
 * @Param value HTTP头数据接口
 * @Param data 需要设置的结构体
 * @Param r HTTP请求
 * @return
 **/
func SetHttpHeaderValueToData(value HeaderValue, data reflect.Value, r *http.Request) error {
	fieldNames := strings.Split(value.GetFieldName(), ".")
	field := data.Elem()
	for _, fieldName := range fieldNames {
		field = field.FieldByName(fieldName)
		if !field.IsValid() {
			return FieldError
		}
	}
	headerValue := r.Header.Get(value.GetHeadKey())
	return value.SetValue(field, headerValue)
}

func DecodeHeaderPostRequest(t reflect.Type, values ...HeaderValue) func(ctx context.Context, r *http.Request) (interface{}, error) {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		v := reflect.New(t)
		if r.ContentLength != 0 {
			if err := json.NewDecoder(r.Body).Decode(v.Interface()); err != nil {
				return nil, err
			}
		}
		fieldMap := make(map[string]*depthFiled)
		MapField(v.Elem(), fieldMap, 0)
		//添加统一处理http头
		for k, _ := range r.Header {
			if field, ok := fieldMap[strings.ToUpper(k)]; ok {
				if err := SetField(field.field, r.Header.Get(k)); err != nil {
					return nil, err
				}
			}
		}
		for _, value := range values {
			if err := SetHttpHeaderValueToData(value, v, r); err != nil {
				return nil, err
			}
		}

		d := v.Elem().Interface()
		return d, nil
	}
}

func DecodeHeaderRequest(t reflect.Type, values []HeaderValue) func(ctx context.Context, r *http.Request) (interface{}, error) {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		v := reflect.New(t)
		fieldMap := make(map[string]*depthFiled)
		MapField(v.Elem(), fieldMap, 0)
		//添加统一处理http头
		for k, _ := range r.Header {
			if field, ok := fieldMap[strings.ToUpper(k)]; ok {
				if err := SetField(field.field, r.Header.Get(k)); err != nil {
					return nil, err
				}
			}
		}
		for _, value := range values {
			if err := SetHttpHeaderValueToData(value, v, r); err != nil {
				return nil, err
			}
		}
		d := v.Elem().Interface()
		return d, nil
	}
}

func SetField(field reflect.Value, value string) error {
	if !field.CanSet() {
		return FieldError
	}
	switch field.Kind() {
	case reflect.Bool:
		if t, err := strconv.ParseBool(value); err != nil {
			return err
		} else {
			field.SetBool(t)
			return nil
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if t, err := strconv.ParseInt(value, 10, 64); err != nil {
			return err
		} else {
			field.SetInt(t)
			return nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if t, err := strconv.ParseUint(value, 10, 64); err != nil {
			return err
		} else {
			field.SetUint(t)
			return nil
		}
	case reflect.Float32, reflect.Float64:
		if t, err := strconv.ParseFloat(value, 64); err != nil {
			return err
		} else {
			field.SetFloat(t)
			return nil
		}
	case reflect.Complex64, reflect.Complex128:
		//复数部分到时候再处理
		return FieldError
	case reflect.String:
		field.SetString(value)
		return nil
	default:
		return FieldError
	}
}

func CheckFieldSet(field reflect.Value) bool {
	if !field.IsValid() {
		return false
	}
	if !field.CanSet() {
		return false
	}
	switch field.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.String:
		return true
	default:
		return false
	}
}

type depthFiled struct {
	depth int
	field reflect.Value
}

func MapField(v reflect.Value, fieldMap map[string]*depthFiled, depth int) {
	if v.Kind() != reflect.Struct {
		return
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Struct {
			MapField(field, fieldMap, depth+1)
		} else if CheckFieldSet(field) {
			fieldName := strings.ToUpper(t.Field(i).Name)
			if f, ok := fieldMap[fieldName]; !ok || f.depth > depth {
				fieldMap[fieldName] = &depthFiled{depth: depth, field: field}
			}
		}
	}
}

func FindField(v reflect.Value, fieldName string) (reflect.Value, bool) {
	if v.Kind() != reflect.Struct {
		return reflect.Value{}, false
	}
	field := v.FieldByName(fieldName)
	if CheckFieldSet(field) {
		return field, true
	}
	for i := 0; i < v.NumField(); i++ {
		if field, ok := FindField(v.Field(i), fieldName); ok {
			return field, true
		}
	}
	return reflect.Value{}, false
}

type BaseHeaderValue struct {
	FieldName string
	HeadKey   string
}

func (v BaseHeaderValue) GetFieldName() string {
	return v.FieldName
}
func (v BaseHeaderValue) GetHeadKey() string {
	return v.HeadKey
}

type StringHeaderValue struct {
	BaseHeaderValue
}

func (v StringHeaderValue) SetValue(field reflect.Value, value string) error {
	field.SetString(value)
	return nil
}

type IntHeaderValue struct {
	BaseHeaderValue
}

func (v IntHeaderValue) SetValue(field reflect.Value, value string) error {
	if t, err := strconv.ParseInt(value, 10, 64); err != nil {
		return err
	} else {
		field.SetInt(t)
		return nil
	}
}

type UIntHeaderValue struct {
	BaseHeaderValue
}

func (v UIntHeaderValue) SetValue(field reflect.Value, value string) error {
	if t, err := strconv.ParseUint(value, 10, 64); err != nil {
		return err
	} else {
		field.SetUint(t)
		return nil
	}
}

type FloatHeaderValue struct {
	BaseHeaderValue
}

func (v FloatHeaderValue) SetValue(field reflect.Value, value string) error {
	if t, err := strconv.ParseFloat(value, 64); err != nil {
		return err
	} else {
		field.SetFloat(t)
		return nil
	}
}

type BoolHeaderValue struct {
	BaseHeaderValue
}

func (v BoolHeaderValue) SetValue(field reflect.Value, value string) error {
	if t, err := strconv.ParseBool(value); err != nil {
		return err
	} else {
		field.SetBool(t)
		return nil
	}
}
