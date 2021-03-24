package orm

import (
	"encoding/json"
	"reflect"
	"testing"
)

/**
 * @Author: feiliang.wang
 * @Description: JSON测试
 * @File:  filters_test
 * @Version: 1.0.0
 * @Date: 2021/1/4 13:58
 */

func TestEnumFilter_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args EnumFilter
		want string
		err  bool
	}{
		{
			"单元素", EnumFilter{1}, "1", false,
		},
		{
			"对元素", EnumFilter{1, 2}, "[1,2]", false,
		},
		{
			"无元素", EnumFilter{}, "[]", false,
		},
		{
			"空", nil, "null", false,
		},
	}
	for _, item := range tests {
		d, err := item.args.MarshalJSON()
		if string(d) != item.want || (err != nil) != item.err {
			t.Fatal(item.name)
		}

	}
}

func TestEnumFilter_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args string
		want EnumFilter
		err  bool
	}{
		{
			"单元素", "1", EnumFilter{1}, false,
		},
		{
			"错误单元素", "1.2", nil, true,
		},
		{
			"多元素", "[1,2]", EnumFilter{1, 2}, false,
		},
		{
			"错误多元素", "[1.2,2]", nil, true,
		},
		{
			"无元素", "[]", EnumFilter{}, false,
		},
		{
			"null", "null", nil, false,
		},
		{
			"NULL", "NULL", nil, true,
		},
		{
			"空字符串", "", nil, true,
		},
		{
			"错误字符串", "{1,2}", nil, true,
		},
	}
	for _, item := range tests {
		var v EnumFilter
		err := json.Unmarshal([]byte(item.args), &v)
		if !reflect.DeepEqual(v, item.want) {
			t.Fatal(item.name)
		} else if (err != nil) != item.err {
			t.Fatal(item.name)
		}

	}
}

func TestEnumFilter_WhereSql(t *testing.T) {
	fieldName := "`abc`"
	tests := []struct {
		name string
		args EnumFilter
		want struct {
			s string
			v []interface{}
		}
	}{
		{
			name: "单元素",
			args: EnumFilter{1},
			want: struct {
				s string
				v []interface{}
			}{
				s: " AND `abc` = ? ",
				v: []interface{}{int32(1)},
			},
		},
		{
			name: "多元素",
			args: EnumFilter{1, 2},
			want: struct {
				s string
				v []interface{}
			}{
				s: " AND `abc` IN (?,?) ",
				v: []interface{}{int32(1), int32(2)},
			},
		},
		{
			name: "无元素",
			args: EnumFilter{},
			want: struct {
				s string
				v []interface{}
			}{
				s: "",
				v: nil,
			},
		},
		{
			name: "null",
			args: nil,
			want: struct {
				s string
				v []interface{}
			}{
				s: "",
				v: nil,
			},
		},
	}
	for _, item := range tests {
		s, v := item.args.WhereSql(fieldName)
		if s != item.want.s {
			t.Fatal(item.name)
		} else if !reflect.DeepEqual(v, item.want.v) {
			t.Fatal(item.name)
		}
	}
}

func TestFloat64Filter_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args *Float64Filter
		want string
		err  bool
	}{
		{
			"最大最小相等", &Float64Filter{1, 1}, "1", false,
		},
		{
			"最大最小不相等", &Float64Filter{1, 2}, "{\"min\":1,\"max\":2}", false,
		},
		{
			"空", nil, "null", false,
		},
	}
	for _, item := range tests {
		d, err := item.args.MarshalJSON()
		if string(d) != item.want || (err != nil) != item.err {
			t.Fatal(item.name)
		}

	}
}

func TestFloat64Filter_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args string
		want Float64Filter
		err  bool
	}{
		{
			"最大最小相等", "1", Float64Filter{1, 1}, false,
		},
		{
			"错误最大最小相等", "\"a\"", Float64Filter{0, 0}, true,
		},
		{
			"最大最小不相等", "{\"min\":1,\"max\":2}", Float64Filter{1, 2}, false,
		},
		{
			"错误最大最小不相等", "{\"min\":\"a\",\"max\":2}", Float64Filter{0, 0}, true,
		},
		{
			"null", "null", Float64Filter{0, 0}, false,
		},
		{
			"NULL", "NULL", Float64Filter{0, 0}, true,
		},
		{
			"空字符串", "", Float64Filter{0, 0}, true,
		},
		{
			"错误字符串", "{1,2}", Float64Filter{0, 0}, true,
		},
	}
	for _, item := range tests {
		var v Float64Filter
		err := json.Unmarshal([]byte(item.args), &v)
		if !reflect.DeepEqual(v, item.want) || (err != nil) != item.err {
			t.Fatal(item.name)
		}

	}
}

func TestFloat64Filter_WhereSql(t *testing.T) {
	fieldName := "`abc`"
	tests := []struct {
		name string
		args *Float64Filter
		want struct {
			s string
			v []interface{}
		}
	}{
		{
			name: "单元素",
			args: &Float64Filter{1, 1},
			want: struct {
				s string
				v []interface{}
			}{
				s: " AND `abc` = ? ",
				v: []interface{}{float64(1)},
			},
		},
		{
			name: "多元素",
			args: &Float64Filter{1, 2},
			want: struct {
				s string
				v []interface{}
			}{
				s: " AND `abc` BETWEEN ? AND ? ",
				v: []interface{}{float64(1), float64(2)},
			},
		},
		{
			name: "null",
			args: nil,
			want: struct {
				s string
				v []interface{}
			}{
				s: "",
				v: nil,
			},
		},
	}
	for _, item := range tests {
		s, v := item.args.WhereSql(fieldName)
		if s != item.want.s {
			t.Fatal(item.name)
		} else if !reflect.DeepEqual(v, item.want.v) {
			t.Fatal(item.name)
		}
	}
}

func TestFloat32Filter_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args *Float32Filter
		want string
		err  bool
	}{
		{
			"最大最小相等", &Float32Filter{1, 1}, "1", false,
		},
		{
			"最大最小不相等", &Float32Filter{1, 2}, "{\"min\":1,\"max\":2}", false,
		},
		{
			"空", nil, "null", false,
		},
	}
	for _, item := range tests {
		d, err := item.args.MarshalJSON()
		if string(d) != item.want || (err != nil) != item.err {
			t.Fatal(item.name)
		}

	}
}

func TestFloat32Filter_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args string
		want Float32Filter
		err  bool
	}{
		{
			"最大最小相等", "1", Float32Filter{1, 1}, false,
		},
		{
			"错误最大最小相等", "\"a\"", Float32Filter{0, 0}, true,
		},
		{
			"最大最小不相等", "{\"min\":1,\"max\":2}", Float32Filter{1, 2}, false,
		},
		{
			"错误最大最小不相等", "{\"min\":\"a\",\"max\":2}", Float32Filter{0, 0}, true,
		},
		{
			"null", "null", Float32Filter{0, 0}, false,
		},
		{
			"NULL", "NULL", Float32Filter{0, 0}, true,
		},
		{
			"空字符串", "", Float32Filter{0, 0}, true,
		},
		{
			"错误字符串", "{1,2}", Float32Filter{0, 0}, true,
		},
	}
	for _, item := range tests {
		var v Float32Filter
		err := json.Unmarshal([]byte(item.args), &v)
		if !reflect.DeepEqual(v, item.want) || (err != nil) != item.err {
			t.Fatal(item.name)
		}

	}
}

func TestFloat32Filter_WhereSql(t *testing.T) {
	fieldName := "`abc`"
	tests := []struct {
		name string
		args *Float32Filter
		want struct {
			s string
			v []interface{}
		}
	}{
		{
			name: "单元素",
			args: &Float32Filter{1, 1},
			want: struct {
				s string
				v []interface{}
			}{
				s: " AND `abc` = ? ",
				v: []interface{}{float32(1)},
			},
		},
		{
			name: "多元素",
			args: &Float32Filter{1, 2},
			want: struct {
				s string
				v []interface{}
			}{
				s: " AND `abc` BETWEEN ? AND ? ",
				v: []interface{}{float32(1), float32(2)},
			},
		},
		{
			name: "null",
			args: nil,
			want: struct {
				s string
				v []interface{}
			}{
				s: "",
				v: nil,
			},
		},
	}
	for _, item := range tests {
		s, v := item.args.WhereSql(fieldName)
		if s != item.want.s {
			t.Fatal(item.name)
		} else if !reflect.DeepEqual(v, item.want.v) {
			t.Fatal(item.name)
		}
	}
}

func TestStringFilter_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args *StringFilter
		want string
		err  bool
	}{
		{
			"空字符串", &StringFilter{""}, "\"\"", false,
		},
		{
			"正常字符串", &StringFilter{"12"}, "\"12\"", false,
		},
		{
			"空", nil, "null", false,
		},
	}
	for _, item := range tests {
		d, err := item.args.MarshalJSON()
		if string(d) != item.want || (err != nil) != item.err {
			t.Fatal(item.name)
		}

	}
}

func TestStringFilter_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args string
		want StringFilter
		err  bool
	}{
		{
			"空字符串", "\"\"", StringFilter{""}, false,
		},
		{
			"不带结构字符串", "\"a\"", StringFilter{"a"}, false,
		},
		{
			"错误不带结构字符串", "123", StringFilter{""}, true,
		},
		{
			"结构字符串", "{\"value\":\"a\"}", StringFilter{"a"}, false,
		},
		{
			"错误结构字符串", "{\"value\":123}", StringFilter{""}, true,
		},
		{
			"null", "null", StringFilter{""}, false,
		},
		{
			"NULL", "NULL", StringFilter{""}, true,
		},
	}
	for _, item := range tests {
		var v StringFilter
		err := json.Unmarshal([]byte(item.args), &v)
		if !reflect.DeepEqual(v, item.want) || (err != nil) != item.err {
			t.Fatal(item.name)
		}

	}
}

func TestStringFilter_WhereSql(t *testing.T) {
	fieldName := "`abc`"
	tests := []struct {
		name string
		args *StringFilter
		want struct {
			s string
			v []interface{}
		}
	}{
		{
			name: "正常字符串",
			args: &StringFilter{"a"},
			want: struct {
				s string
				v []interface{}
			}{
				s: " AND `abc` like ? ",
				v: []interface{}{"a"},
			},
		},
		{
			name: "空字符串",
			args: &StringFilter{""},
			want: struct {
				s string
				v []interface{}
			}{
				s: "",
				v: nil,
			},
		},
		{
			name: "null",
			args: nil,
			want: struct {
				s string
				v []interface{}
			}{
				s: "",
				v: nil,
			},
		},
	}
	for _, item := range tests {
		s, v := item.args.WhereSql(fieldName)
		if s != item.want.s {
			t.Fatal(item.name)
		} else if !reflect.DeepEqual(v, item.want.v) {
			t.Fatal(item.name)
		}
	}
}

func TestUInt64Filter_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args *UInt64Filter
		want string
		err  bool
	}{
		{
			"最大最小相等", &UInt64Filter{1, 1}, "1", false,
		},
		{
			"最大最小不相等", &UInt64Filter{1, 2}, "{\"min\":1,\"max\":2}", false,
		},
		{
			"空", nil, "null", false,
		},
	}
	for _, item := range tests {
		d, err := item.args.MarshalJSON()
		if string(d) != item.want || (err != nil) != item.err {
			t.Fatal(item.name)
		}

	}
}

func TestUInt64Filter_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args string
		want UInt64Filter
		err  bool
	}{
		{
			"最大最小相等", "1", UInt64Filter{1, 1}, false,
		},
		{
			"错误最大最小相等", "\"a\"", UInt64Filter{0, 0}, true,
		},
		{
			"最大最小不相等", "{\"min\":1,\"max\":2}", UInt64Filter{1, 2}, false,
		},
		{
			"错误最大最小不相等", "{\"min\":\"a\",\"max\":2}", UInt64Filter{0, 0}, true,
		},
		{
			"null", "null", UInt64Filter{0, 0}, false,
		},
		{
			"NULL", "NULL", UInt64Filter{0, 0}, true,
		},
		{
			"空字符串", "", UInt64Filter{0, 0}, true,
		},
		{
			"错误字符串", "{1,2}", UInt64Filter{0, 0}, true,
		},
	}
	for _, item := range tests {
		var v UInt64Filter
		err := json.Unmarshal([]byte(item.args), &v)
		if !reflect.DeepEqual(v, item.want) || (err != nil) != item.err {
			t.Fatal(item.name)
		}

	}
}

func TestUInt64Filter_WhereSql(t *testing.T) {
	fieldName := "`abc`"
	tests := []struct {
		name string
		args *UInt64Filter
		want struct {
			s string
			v []interface{}
		}
	}{
		{
			name: "单元素",
			args: &UInt64Filter{1, 1},
			want: struct {
				s string
				v []interface{}
			}{
				s: " AND `abc` = ? ",
				v: []interface{}{uint64(1)},
			},
		},
		{
			name: "多元素",
			args: &UInt64Filter{1, 2},
			want: struct {
				s string
				v []interface{}
			}{
				s: " AND `abc` BETWEEN ? AND ? ",
				v: []interface{}{uint64(1), uint64(2)},
			},
		},
		{
			name: "null",
			args: nil,
			want: struct {
				s string
				v []interface{}
			}{
				s: "",
				v: nil,
			},
		},
	}
	for _, item := range tests {
		s, v := item.args.WhereSql(fieldName)
		if s != item.want.s {
			t.Fatal(item.name)
		} else if !reflect.DeepEqual(v, item.want.v) {
			t.Fatal(item.name)
		}
	}
}

func TestInt64Filter_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args *Int64Filter
		want string
		err  bool
	}{
		{
			"最大最小相等", &Int64Filter{1, 1}, "1", false,
		},
		{
			"最大最小不相等", &Int64Filter{1, 2}, "{\"min\":1,\"max\":2}", false,
		},
		{
			"空", nil, "null", false,
		},
	}
	for _, item := range tests {
		d, err := item.args.MarshalJSON()
		if string(d) != item.want || (err != nil) != item.err {
			t.Fatal(item.name)
		}

	}
}

func TestInt64Filter_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args string
		want Int64Filter
		err  bool
	}{
		{
			"最大最小相等", "1", Int64Filter{1, 1}, false,
		},
		{
			"错误最大最小相等", "\"a\"", Int64Filter{0, 0}, true,
		},
		{
			"最大最小不相等", "{\"min\":1,\"max\":2}", Int64Filter{1, 2}, false,
		},
		{
			"错误最大最小不相等", "{\"min\":\"a\",\"max\":2}", Int64Filter{0, 0}, true,
		},
		{
			"null", "null", Int64Filter{0, 0}, false,
		},
		{
			"NULL", "NULL", Int64Filter{0, 0}, true,
		},
		{
			"空字符串", "", Int64Filter{0, 0}, true,
		},
		{
			"错误字符串", "{1,2}", Int64Filter{0, 0}, true,
		},
	}
	for _, item := range tests {
		var v Int64Filter
		err := json.Unmarshal([]byte(item.args), &v)
		if !reflect.DeepEqual(v, item.want) || (err != nil) != item.err {
			t.Fatal(item.name)
		}

	}
}

func TestInt64Filter_WhereSql(t *testing.T) {
	fieldName := "`abc`"
	tests := []struct {
		name string
		args *Int64Filter
		want struct {
			s string
			v []interface{}
		}
	}{
		{
			name: "单元素",
			args: &Int64Filter{1, 1},
			want: struct {
				s string
				v []interface{}
			}{
				s: " AND `abc` = ? ",
				v: []interface{}{int64(1)},
			},
		},
		{
			name: "多元素",
			args: &Int64Filter{1, 2},
			want: struct {
				s string
				v []interface{}
			}{
				s: " AND `abc` BETWEEN ? AND ? ",
				v: []interface{}{int64(1), int64(2)},
			},
		},
		{
			name: "null",
			args: nil,
			want: struct {
				s string
				v []interface{}
			}{
				s: "",
				v: nil,
			},
		},
	}
	for _, item := range tests {
		s, v := item.args.WhereSql(fieldName)
		if s != item.want.s {
			t.Fatal(item.name)
		} else if !reflect.DeepEqual(v, item.want.v) {
			t.Fatal(item.name)
		}
	}
}

func TestUInt32Filter_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args *UInt32Filter
		want string
		err  bool
	}{
		{
			"最大最小相等", &UInt32Filter{1, 1}, "1", false,
		},
		{
			"最大最小不相等", &UInt32Filter{1, 2}, "{\"min\":1,\"max\":2}", false,
		},
		{
			"空", nil, "null", false,
		},
	}
	for _, item := range tests {
		d, err := item.args.MarshalJSON()
		if string(d) != item.want || (err != nil) != item.err {
			t.Fatal(item.name)
		}

	}
}

func TestUInt32Filter_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args string
		want UInt32Filter
		err  bool
	}{
		{
			"最大最小相等", "1", UInt32Filter{1, 1}, false,
		},
		{
			"错误最大最小相等", "\"a\"", UInt32Filter{0, 0}, true,
		},
		{
			"最大最小不相等", "{\"min\":1,\"max\":2}", UInt32Filter{1, 2}, false,
		},
		{
			"错误最大最小不相等", "{\"min\":\"a\",\"max\":2}", UInt32Filter{0, 0}, true,
		},
		{
			"null", "null", UInt32Filter{0, 0}, false,
		},
		{
			"NULL", "NULL", UInt32Filter{0, 0}, true,
		},
		{
			"空字符串", "", UInt32Filter{0, 0}, true,
		},
		{
			"错误字符串", "{1,2}", UInt32Filter{0, 0}, true,
		},
	}
	for _, item := range tests {
		var v UInt32Filter
		err := json.Unmarshal([]byte(item.args), &v)
		if !reflect.DeepEqual(v, item.want) || (err != nil) != item.err {
			t.Fatal(item.name)
		}

	}
}

func TestUInt32Filter_WhereSql(t *testing.T) {
	fieldName := "`abc`"
	tests := []struct {
		name string
		args *UInt32Filter
		want struct {
			s string
			v []interface{}
		}
	}{
		{
			name: "单元素",
			args: &UInt32Filter{1, 1},
			want: struct {
				s string
				v []interface{}
			}{
				s: " AND `abc` = ? ",
				v: []interface{}{uint32(1)},
			},
		},
		{
			name: "多元素",
			args: &UInt32Filter{1, 2},
			want: struct {
				s string
				v []interface{}
			}{
				s: " AND `abc` BETWEEN ? AND ? ",
				v: []interface{}{uint32(1), uint32(2)},
			},
		},
		{
			name: "null",
			args: nil,
			want: struct {
				s string
				v []interface{}
			}{
				s: "",
				v: nil,
			},
		},
	}
	for _, item := range tests {
		s, v := item.args.WhereSql(fieldName)
		if s != item.want.s {
			t.Fatal(item.name)
		} else if !reflect.DeepEqual(v, item.want.v) {
			t.Fatal(item.name)
		}
	}
}

func TestInt32Filter_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args *Int32Filter
		want string
		err  bool
	}{
		{
			"最大最小相等", &Int32Filter{1, 1}, "1", false,
		},
		{
			"最大最小不相等", &Int32Filter{1, 2}, "{\"min\":1,\"max\":2}", false,
		},
		{
			"空", nil, "null", false,
		},
	}
	for _, item := range tests {
		d, err := item.args.MarshalJSON()
		if string(d) != item.want || (err != nil) != item.err {
			t.Fatal(item.name)
		}

	}
}

func TestInt32Filter_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args string
		want Int32Filter
		err  bool
	}{
		{
			"最大最小相等", "1", Int32Filter{1, 1}, false,
		},
		{
			"错误最大最小相等", "\"a\"", Int32Filter{0, 0}, true,
		},
		{
			"最大最小不相等", "{\"min\":1,\"max\":2}", Int32Filter{1, 2}, false,
		},
		{
			"错误最大最小不相等", "{\"min\":\"a\",\"max\":2}", Int32Filter{0, 0}, true,
		},
		{
			"null", "null", Int32Filter{0, 0}, false,
		},
		{
			"NULL", "NULL", Int32Filter{0, 0}, true,
		},
		{
			"空字符串", "", Int32Filter{0, 0}, true,
		},
		{
			"错误字符串", "{1,2}", Int32Filter{0, 0}, true,
		},
	}
	for _, item := range tests {
		var v Int32Filter
		err := json.Unmarshal([]byte(item.args), &v)
		if !reflect.DeepEqual(v, item.want) || (err != nil) != item.err {
			t.Fatal(item.name)
		}

	}
}

func TestInt32Filter_WhereSql(t *testing.T) {
	fieldName := "`abc`"
	tests := []struct {
		name string
		args *Int32Filter
		want struct {
			s string
			v []interface{}
		}
	}{
		{
			name: "单元素",
			args: &Int32Filter{1, 1},
			want: struct {
				s string
				v []interface{}
			}{
				s: " AND `abc` = ? ",
				v: []interface{}{int32(1)},
			},
		},
		{
			name: "多元素",
			args: &Int32Filter{1, 2},
			want: struct {
				s string
				v []interface{}
			}{
				s: " AND `abc` BETWEEN ? AND ? ",
				v: []interface{}{int32(1), int32(2)},
			},
		},
		{
			name: "null",
			args: nil,
			want: struct {
				s string
				v []interface{}
			}{
				s: "",
				v: nil,
			},
		},
	}
	for _, item := range tests {
		s, v := item.args.WhereSql(fieldName)
		if s != item.want.s {
			t.Fatal(item.name)
		} else if !reflect.DeepEqual(v, item.want.v) {
			t.Fatal(item.name)
		}
	}
}

func TestMakeSqlFilterMap(t *testing.T) {
	tests := []struct {
		name string
		args interface{}
		want SqlFilterMap
	}{
		{
			name: "正常结构体",
			args: struct {
				S *StringFilter
			}{
				S: &StringFilter{"1"},
			},
			want: SqlFilterMap{
				"S": &StringFilter{"1"},
			},
		},
		{
			name: "参杂正常结构体",
			args: struct {
				S  *StringFilter
				S2 string
			}{
				S:  &StringFilter{"1"},
				S2: "2",
			},
			want: SqlFilterMap{
				"S": &StringFilter{"1"},
			},
		},
		{
			name: "正常结构体指针",
			args: &struct {
				S *StringFilter
			}{
				S: &StringFilter{"1"},
			},
			want: SqlFilterMap{
				"S": &StringFilter{"1"},
			},
		},
		{
			name: "多层正常结构体",
			args: &struct {
				S  *StringFilter
				S1 struct {
					I *Int32Filter
				}
			}{
				S: &StringFilter{"1"},
				S1: struct {
					I *Int32Filter
				}{
					I: &Int32Filter{2, 2},
				},
			},
			want: SqlFilterMap{
				"S": &StringFilter{"1"},
				"I": &Int32Filter{2, 2},
			},
		},
	}
	for _, item := range tests {
		s := MakeSqlFilterMap(item.args)
		if !reflect.DeepEqual(s, item.want) {
			t.Fatal(item.name)
		}
	}
}

func TestSqlFilterMap_WhereSql(t *testing.T) {
	tests := []struct {
		name string
		args SqlFilterMap
		want struct {
			s string
			v []interface{}
		}
	}{
		{
			name: "正常结构体",
			args: SqlFilterMap{
				"a": &StringFilter{"1"},
				"i": &Int32Filter{2, 2},
			},
			want: struct {
				s string
				v []interface{}
			}{
				s: " AND `a` like ?  AND `i` = ? ",
				v: []interface{}{"1", int32(2)},
			},
		},
	}
	for _, item := range tests {
		s, v := item.args.WhereSql()
		if s != item.want.s || !reflect.DeepEqual(v, item.want.v) {
			t.Fatal(item.name)
		}
	}
}
