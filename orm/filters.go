package orm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

/**
 * @Author: feiliang.wang
 * @Description: 筛选参数结构定义
 * @File:  filters
 * @Version: 1.0.0
 * @Date: 2020/8/1 4:07 下午
 */

type SqlFilter interface {
	WhereSql(fieldName string) (string, []interface{})
}

type SqlFilterMap map[string]SqlFilter

var sqlFilterType = reflect.TypeOf((*SqlFilter)(nil)).Elem()

func MakeSqlFilterMap(req interface{}) SqlFilterMap {
	v := reflect.ValueOf(req)
	switch v.Type().Kind() {
	case reflect.Ptr:
		v = v.Elem()
	case reflect.Struct:
	default:
		panic(fmt.Sprintf("not support kind %s", v.Type().Kind()))
	}
	result := make(SqlFilterMap)
	for i := 0; i < v.NumField(); i++ {
		setSqlFilterMap(v.Type().Field(i).Name, v.Field(i), result)
	}
	return result
}

func setSqlFilterMap(fieldName string, fieldValue reflect.Value, maps SqlFilterMap) {
	if fieldValue.Type().Implements(sqlFilterType) {
		maps[fieldName] = fieldValue.Interface().(SqlFilter)
	} else if fieldValue.Type().Kind() == reflect.Struct {
		for i := 0; i < fieldValue.NumField(); i++ {
			setSqlFilterMap(fieldValue.Type().Field(i).Name, fieldValue.Field(i), maps)
		}
	}
}

func (m SqlFilterMap) WhereSql() (string, []interface{}) {
	sb := strings.Builder{}
	values := make([]interface{}, 0)
	keys := make([]string, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		sql, value := m[k].WhereSql(SnakeField(k))
		values = append(values, value...)
		sb.WriteString(sql)
	}
	return sb.String(), values
}

type Int32Filter struct {
	Min int32 `json:"min,omitempty"`
	Max int32 `json:"max,omitempty"`
}

func (f *Int32Filter) WhereSql(fieldName string) (string, []interface{}) {
	if f == nil {
		return "", nil
	} else if f.Min == f.Max {
		return fmt.Sprintf(" AND %s = ? ", fieldName), []interface{}{f.Min}
	} else {
		return fmt.Sprintf(" AND %s BETWEEN ? AND ? ", fieldName), []interface{}{f.Min, f.Max}
	}
}

func (m *Int32Filter) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, []byte("null")) == 0 {
		return nil
	}
	if len(data) > 0 && data[0] != '{' {
		if v, err := strconv.ParseInt(string(data), 10, 32); err != nil {
			return err
		} else {
			*m = Int32Filter{int32(v), int32(v)}
			return nil
		}
	} else {
		v := struct {
			Min int32 `json:"min,omitempty"`
			Max int32 `json:"max,omitempty"`
		}{}
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		} else {
			*m = v
			return nil
		}
	}
}

func (m *Int32Filter) MarshalJSON() (data []byte, err error) {
	if m == nil {
		return []byte("null"), nil
	} else if m.Min == m.Max {
		return json.Marshal(m.Min)
	} else {
		return json.Marshal(struct {
			Min int32 `json:"min,omitempty"`
			Max int32 `json:"max,omitempty"`
		}{
			m.Min,
			m.Max,
		})
	}
}

type UInt32Filter struct {
	Min uint32 `json:"min,omitempty"`
	Max uint32 `json:"max,omitempty"`
}

func (f *UInt32Filter) WhereSql(fieldName string) (string, []interface{}) {
	if f == nil {
		return "", nil
	} else if f.Min == f.Max {
		return fmt.Sprintf(" AND %s = ? ", fieldName), []interface{}{f.Min}
	} else {
		return fmt.Sprintf(" AND %s BETWEEN ? AND ? ", fieldName), []interface{}{f.Min, f.Max}
	}
}

func (m *UInt32Filter) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, []byte("null")) == 0 {
		return nil
	}
	if len(data) > 0 && data[0] != '{' {
		if v, err := strconv.ParseUint(string(data), 10, 32); err != nil {
			return err
		} else {
			*m = UInt32Filter{uint32(v), uint32(v)}
			return nil
		}
	} else {
		v := struct {
			Min uint32 `json:"min,omitempty"`
			Max uint32 `json:"max,omitempty"`
		}{}
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		} else {
			*m = v
			return nil
		}
	}
}

func (m *UInt32Filter) MarshalJSON() (data []byte, err error) {
	if m == nil {
		return []byte("null"), nil
	} else if m.Min == m.Max {
		return json.Marshal(m.Min)
	} else {
		return json.Marshal(struct {
			Min uint32 `json:"min,omitempty"`
			Max uint32 `json:"max,omitempty"`
		}{
			m.Min,
			m.Max,
		})
	}
}

type Int64Filter struct {
	Min int64 `json:"min,omitempty"`
	Max int64 `json:"max,omitempty"`
}

func (f *Int64Filter) WhereSql(fieldName string) (string, []interface{}) {
	if f == nil {
		return "", nil
	} else if f.Min == f.Max {
		return fmt.Sprintf(" AND %s = ? ", fieldName), []interface{}{f.Min}
	} else {
		return fmt.Sprintf(" AND %s BETWEEN ? AND ? ", fieldName), []interface{}{f.Min, f.Max}
	}
}

func (m *Int64Filter) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, []byte("null")) == 0 {
		return nil
	}
	if len(data) > 0 && data[0] != '{' {
		if v, err := strconv.ParseInt(string(data), 10, 64); err != nil {
			return err
		} else {
			*m = Int64Filter{int64(v), int64(v)}
			return nil
		}
	} else {
		v := struct {
			Min int64 `json:"min,omitempty"`
			Max int64 `json:"max,omitempty"`
		}{}
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		} else {
			*m = v
			return nil
		}
	}
}

func (m *Int64Filter) MarshalJSON() (data []byte, err error) {
	if m == nil {
		return []byte("null"), nil
	} else if m.Min == m.Max {
		return json.Marshal(m.Min)
	} else {
		return json.Marshal(struct {
			Min int64 `json:"min,omitempty"`
			Max int64 `json:"max,omitempty"`
		}{
			m.Min,
			m.Max,
		})
	}
}

type UInt64Filter struct {
	Min uint64 `json:"min,omitempty"`
	Max uint64 `json:"max,omitempty"`
}

func (f *UInt64Filter) WhereSql(fieldName string) (string, []interface{}) {
	if f == nil {
		return "", nil
	} else if f.Min == f.Max {
		return fmt.Sprintf(" AND %s = ? ", fieldName), []interface{}{f.Min}
	} else {
		return fmt.Sprintf(" AND %s BETWEEN ? AND ? ", fieldName), []interface{}{f.Min, f.Max}
	}
}

func (m *UInt64Filter) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, []byte("null")) == 0 {
		return nil
	}
	if len(data) > 0 && data[0] != '{' {
		if v, err := strconv.ParseUint(string(data), 10, 64); err != nil {
			return err
		} else {
			*m = UInt64Filter{uint64(v), uint64(v)}
			return nil
		}
	} else {
		v := struct {
			Min uint64 `json:"min,omitempty"`
			Max uint64 `json:"max,omitempty"`
		}{}
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		} else {
			*m = v
			return nil
		}
	}
}

func (m *UInt64Filter) MarshalJSON() (data []byte, err error) {
	if m == nil {
		return []byte("null"), nil
	} else if m.Min == m.Max {
		return json.Marshal(m.Min)
	} else {
		return json.Marshal(struct {
			Min uint64 `json:"min,omitempty"`
			Max uint64 `json:"max,omitempty"`
		}{
			m.Min,
			m.Max,
		})
	}
}

type StringFilter struct {
	Value string `json:"value,omitempty"`
}

func (f *StringFilter) WhereSql(fieldName string) (string, []interface{}) {
	if f == nil || f.Value == "" {
		return "", nil
	}
	return fmt.Sprintf(" AND %s like ? ", fieldName), []interface{}{f.Value}
}

func (m *StringFilter) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, []byte("null")) == 0 {
		return nil
	}
	if len(data) > 0 && data[0] != '{' {
		if data[0] != '"' || data[len(data)-1] != '"' {
			return &json.InvalidUnmarshalError{
				Type: reflect.TypeOf(""),
			}
		} else {
			*m = StringFilter{string(data[1 : len(data)-1])}
			return nil
		}
	} else {
		v := struct {
			Value string `json:"value,omitempty"`
		}{}
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		} else {
			*m = v
			return nil
		}
	}
}

func (m *StringFilter) MarshalJSON() (data []byte, err error) {
	if m == nil {
		return []byte("null"), nil
	} else {
		return json.Marshal(m.Value)
	}
}

type Float32Filter struct {
	Min float32 `json:"min,omitempty"`
	Max float32 `json:"max,omitempty"`
}

func (f *Float32Filter) WhereSql(fieldName string) (string, []interface{}) {
	if f == nil {
		return "", nil
	} else if f.Min == f.Max {
		return fmt.Sprintf(" AND %s = ? ", fieldName), []interface{}{f.Min}
	} else {
		return fmt.Sprintf(" AND %s BETWEEN ? AND ? ", fieldName), []interface{}{f.Min, f.Max}
	}
}

func (m *Float32Filter) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, []byte("null")) == 0 {
		return nil
	}
	if len(data) > 0 && data[0] != '{' {
		if v, err := strconv.ParseFloat(string(data), 32); err != nil {
			return err
		} else {
			*m = Float32Filter{float32(v), float32(v)}
			return nil
		}
	} else {
		v := struct {
			Min float32 `json:"min,omitempty"`
			Max float32 `json:"max,omitempty"`
		}{
			m.Min,
			m.Max,
		}
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		} else {
			*m = v
			return nil
		}
	}
}

func (m *Float32Filter) MarshalJSON() (data []byte, err error) {
	if m == nil {
		return []byte("null"), nil
	} else if m.Min == m.Max {
		return json.Marshal(m.Min)
	} else {
		return json.Marshal(struct {
			Min float32 `json:"min,omitempty"`
			Max float32 `json:"max,omitempty"`
		}{
			m.Min,
			m.Max,
		})
	}
}

type Float64Filter struct {
	Min float64 `json:"min,omitempty"`
	Max float64 `json:"max,omitempty"`
}

func (f *Float64Filter) WhereSql(fieldName string) (string, []interface{}) {
	if f == nil {
		return "", nil
	} else if f.Min == f.Max {
		return fmt.Sprintf(" AND %s = ? ", fieldName), []interface{}{f.Min}
	} else {
		return fmt.Sprintf(" AND %s BETWEEN ? AND ? ", fieldName), []interface{}{f.Min, f.Max}
	}
}

func (m *Float64Filter) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, []byte("null")) == 0 {
		return nil
	}
	if len(data) > 0 && data[0] != '{' {
		if v, err := strconv.ParseFloat(string(data), 64); err != nil {
			return err
		} else {
			*m = Float64Filter{v, v}
			return nil
		}
	} else {
		v := struct {
			Min float64 `json:"min,omitempty"`
			Max float64 `json:"max,omitempty"`
		}{
			m.Min,
			m.Max,
		}
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		} else {
			*m = v
			return nil
		}
	}
}

func (m *Float64Filter) MarshalJSON() (data []byte, err error) {
	if m == nil {
		return []byte("null"), nil
	} else if m.Min == m.Max {
		return json.Marshal(m.Min)
	} else {
		return json.Marshal(struct {
			Min float64 `json:"min,omitempty"`
			Max float64 `json:"max,omitempty"`
		}{
			m.Min,
			m.Max,
		})
	}
}

type EnumFilter []int32

func (f *EnumFilter) WhereSql(fieldName string) (string, []interface{}) {
	if f == nil || len(*f) == 0 {
		return "", nil
	} else if len(*f) == 1 {
		return fmt.Sprintf(" AND %s = ? ", fieldName), []interface{}{(*f)[0]}
	} else {
		build := &strings.Builder{}
		build.WriteString(fmt.Sprintf(" AND %s IN (?", fieldName))
		for i := 1; i < len(*f); i++ {
			build.WriteString(",?")
		}
		build.WriteString(") ")
		valuse := make([]interface{}, len(*f))
		for i := 0; i < len(*f); i++ {
			valuse[i] = (*f)[i]
		}
		return build.String(), valuse
	}
}

func (m *EnumFilter) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, []byte("null")) == 0 {
		*m = nil
		return nil
	}
	if len(data) > 0 && data[0] != '[' {
		if v, err := strconv.ParseInt(string(data), 10, 32); err != nil {
			*m = nil
			return err
		} else {
			*m = EnumFilter{int32(v)}
			return nil
		}
	} else {
		var v []int32
		if err := json.Unmarshal(data, &v); err != nil {
			*m = nil
			return err
		} else {
			*m = v
			return nil
		}
	}
}

func (m *EnumFilter) MarshalJSON() (data []byte, err error) {
	if m == nil || *m == nil {
		return []byte("null"), nil
	} else if len(*m) == 1 {
		return json.Marshal((*m)[0])
	} else {
		return json.Marshal([]int32(*m))
	}
}
