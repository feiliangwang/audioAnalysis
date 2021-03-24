package orm

import (
	"fmt"
	"reflect"
	"strings"
)

/**
 * @Author: feiliang.wang
 * @Description: 数据库查询时列的处理
 * @File:  column
 * @Version: 1.0.0
 * @Date: 2020/8/5 1:54 下午
 */

func MakeSqlColumns(req interface{}) string {
	t := reflect.TypeOf(req)
	switch t.Kind() {
	case reflect.Ptr:
		t = t.Elem()
	case reflect.Struct:
	default:
		panic(fmt.Sprintf("not support kind %s", t.Kind()))
	}
	num := get_field_num(t)
	cloums := make([]string, num)
	pos := 0

	add_cloum_name(t, "", cloums, &pos)
	sb := strings.Builder{}
	for i := 0; i < len(cloums); i++ {
		if i != 0 {
			sb.WriteString(",")
		}
		sb.WriteString(cloums[i])
	}
	return sb.String()
}

var IdColoumName = SnakeField("Id")

func MakeSqlInsertColumnsAndArgs(req interface{}, tableName string) (string, []interface{}) {
	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)
	switch t.Kind() {
	case reflect.Ptr:
		t = t.Elem()
		v = v.Elem()
	case reflect.Struct:
	default:
		panic(fmt.Sprintf("not support kind %s", t.Kind()))
	}

	sb := strings.Builder{}
	sb2 := strings.Builder{}
	num := get_field_num(t)
	coloums := make([]string, num)
	argsTemp := make([]interface{}, num)
	var args []interface{}
	pos := 0

	add_cloum_name_and_value(t, v, "", coloums, argsTemp, &pos)

	first := true
	for i := 0; i < len(coloums); i++ {
		if coloums[i] == IdColoumName {
			continue
		}
		if !first {
			sb.WriteString(",")
			sb2.WriteString(",")
		}
		sb.WriteString(coloums[i])
		sb2.WriteString("?")
		args = append(args, argsTemp[i])
		first = false
	}

	return fmt.Sprintf("insert into %s (%s) values(%s)", tableName, sb.String(), sb2.String()), args
}

func get_field_num(t reflect.Type) int {
	if t.PkgPath()+"/"+t.Name() == "github.com/shopspring/decimal/Decimal" {
		return 1
	}
	switch t.Kind() {
	case reflect.Ptr:
		return get_field_num(t.Elem())
	case reflect.Struct:
		count := 0
		for i := 0; i < t.NumField(); i++ {
			count += get_field_num(t.Field(i).Type)
		}
		return count
	default:
		return 1
	}
}

func add_cloum_name_and_value(t reflect.Type, v reflect.Value, fieldName string, cloums []string, args []interface{}, pos *int) {
	if t.PkgPath()+"/"+t.Name() == "github.com/shopspring/decimal/Decimal" {
		cloums[*pos] = SnakeField(fieldName)
		args[*pos] = v.Interface()
		*pos += 1
		return
	}
	switch t.Kind() {
	case reflect.Ptr:
		add_cloum_name_and_value(t.Elem(), v.Elem(), "", cloums, args, pos)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			add_cloum_name_and_value(t.Field(i).Type, v.Field(i), t.Field(i).Name, cloums, args, pos)
		}
	default:
		cloums[*pos] = SnakeField(fieldName)
		args[*pos] = v.Interface()
		*pos += 1
	}
}

func add_cloum_name(t reflect.Type, fieldName string, cloums []string, pos *int) {
	if t.PkgPath()+"/"+t.Name() == "github.com/shopspring/decimal/Decimal" {
		cloums[*pos] = SnakeField(fieldName)
		*pos += 1
		return
	}
	switch t.Kind() {
	case reflect.Ptr:
		add_cloum_name(t.Elem(), "", cloums, pos)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			add_cloum_name(t.Field(i).Type, t.Field(i).Name, cloums, pos)
		}
	default:
		cloums[*pos] = SnakeField(fieldName)
		*pos += 1
	}
}

func ScanDao(dao interface{}, scan func(dest ...interface{}) error) error {
	v := reflect.ValueOf(dao)
	switch v.Kind() {
	case reflect.Ptr:
		v = v.Elem()
	default:
		panic(fmt.Sprintf("not support kind %s", v.Kind()))
	}
	count := get_field_num(v.Type())
	dest := make([]interface{}, count)
	pos := 0
	set_column_value(v, dest, &pos)
	if err := scan(dest...); err != nil {
		return err
	}
	return nil
}

func set_column_value(v reflect.Value, dest []interface{}, pos *int) {
	if v.Type().PkgPath()+"/"+v.Type().Name() == "github.com/shopspring/decimal/Decimal" {
		dest[*pos] = v.Addr().Interface()
		*pos += 1
		return
	}
	switch v.Kind() {
	case reflect.Ptr:
		set_column_value(v.Elem(), dest, pos)
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			set_column_value(v.Field(i), dest, pos)
		}
	default:
		dest[*pos] = v.Addr().Interface()
		*pos += 1
	}

}
