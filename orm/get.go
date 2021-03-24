package orm

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"
)

/**
 * @Author: feiliang.wang
 * @Description: 数据库操作
 * @File:  get
 * @Version: 1.0.0
 * @Date: 2020/8/11 1:32 下午
 */

type queryRow = func(string, ...interface{}) *sql.Row
type exec = func(string, ...interface{}) (sql.Result, error)
type query = func(string, ...interface{}) (*sql.Rows, error)

func getCount(qr queryRow, tableName string, filters SqlFilterMap) (count int, err error) {
	sb := &strings.Builder{}
	where, args := filters.WhereSql()
	sb.WriteString(fmt.Sprintf("select count(*) from %s where 1 = 1 %s", tableName, where))
	row := qr(sb.String(), args...)
	err = row.Scan(&count)
	return
}

func GetCount(db *sql.DB, tableName string, filters SqlFilterMap) (count int, err error) {
	return getCount(db.QueryRow, tableName, filters)
}

func GetCountTx(tx *sql.Tx, tableName string, filters SqlFilterMap) (count int, err error) {
	return getCount(tx.QueryRow, tableName, filters)
}

func getInfo(qr queryRow, tableName string, id int32, data interface{}) error {
	sqlStr, args := fmt.Sprintf("select %s from %s where %s = ?", MakeSqlColumns(data), tableName, SnakeField("Id")), []interface{}{id}
	row := qr(sqlStr, args...)
	return ScanDao(data, row.Scan)
}

func GetInfo(db *sql.DB, tableName string, id int32, data interface{}) error {
	return getInfo(db.QueryRow, tableName, id, data)
}

func GetInfoTx(tx *sql.Tx, tableName string, id int32, data interface{}) error {
	return getInfo(tx.QueryRow, tableName, id, data)
}

func deleteInfo(exec exec, tableName string, id int32) (ok bool, err error) {
	sqlStr, args := fmt.Sprintf("delete from %s where `id` = ?", tableName), []interface{}{id}
	if _, err = exec(sqlStr, args...); err == nil {
		ok = true
	}
	return
}

func DeleteInfo(db *sql.DB, tableName string, id int32) (ok bool, err error) {
	return deleteInfo(db.Exec, tableName, id)
}

func DeleteInfoTx(tx *sql.Tx, tableName string, id int32) (ok bool, err error) {
	return deleteInfo(tx.Exec, tableName, id)
}

func insertInfo(exec exec, tableName string, info interface{}) (id int32, err error) {
	sqlStr, args := MakeSqlInsertColumnsAndArgs(info, tableName)
	var result sql.Result
	if result, err = exec(sqlStr, args...); err == nil {
		var insertId int64
		if insertId, err = result.LastInsertId(); err == nil {
			id = int32(insertId)
		}
	}
	return
}

func InsertInfo(db *sql.DB, tableName string, info interface{}) (id int32, err error) {
	return insertInfo(db.Exec, tableName, info)
}

func InsertInfoTx(tx *sql.Tx, tableName string, info interface{}) (id int32, err error) {
	return insertInfo(tx.Exec, tableName, info)
}

func updateInfo(exec exec, tableName string, id int32, sets map[string]interface{}) (ok bool, err error) {
	if sets == nil || len(sets) == 0 {
		return false, fmt.Errorf("there are no updates value")
	}
	sb := &strings.Builder{}
	sb.WriteString(fmt.Sprintf("update %s set ", tableName))
	keys := make([]string, 0, len(sets))
	values := make([]interface{}, 0, len(sets)+1)
	for k, _ := range sets {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i, key := range keys {
		if i != 0 {
			sb.WriteString(",")
		}
		sb.WriteString(fmt.Sprintf(" %s = ? ", SnakeField(key)))
		values = append(values, sets[key])
	}
	sb.WriteString(" where `Id` = ?")
	values = append(values, id)

	if _, err = exec(sb.String(), values...); err == nil {
		ok = true
	}
	return
}

func UpdateInfo(db *sql.DB, tableName string, id int32, sets map[string]interface{}) (ok bool, err error) {
	return updateInfo(db.Exec, tableName, id, sets)
}

func UpdateInfoTx(tx *sql.Tx, tableName string, id int32, sets map[string]interface{}) (ok bool, err error) {
	return updateInfo(tx.Exec, tableName, id, sets)
}

func listInfo(query query, logger *log.Logger, tableName string, sortAsc int, sortField string, num, size int, filters SqlFilterMap, list interface{}) (count int, err error) {
	t := reflect.TypeOf(list)
	if t.Kind() != reflect.Ptr {
		panic("list kind is not slice pointer")
	}
	t = t.Elem()
	if t.Kind() != reflect.Slice {
		panic("list kind is not slice pointer")
	}
	rv := reflect.ValueOf(list).Elem()
	where, args := filters.WhereSql()
	var rows *sql.Rows
	sqlStr := fmt.Sprintf("select %s from %s Where 1 = 1 %s %s %s",
		MakeSqlColumns(reflect.New(t.Elem()).Interface()), tableName, where,
		Order(sortAsc, sortField),
		Limit(num, size))
	if rows, err = query(sqlStr, args...); err == nil {
		defer func() {
			if err := rows.Close(); err != nil {
				logger.Println(err)
			}
		}()
		for rows.Next() {
			item := reflect.New(t.Elem())
			value := item.Interface()
			err = ScanDao(value, rows.Scan)
			if err != nil {
				return
			}
			count++
			rv = reflect.Append(rv, item.Elem())
		}
	}

	reflect.ValueOf(list).Elem().Set(rv)
	return
}

func ListInfo(db *sql.DB, logger *log.Logger, tableName string, sortAsc int, sortField string, num, size int, filters SqlFilterMap, list interface{}) (count int, err error) {
	return listInfo(db.Query, logger, tableName, sortAsc, sortField, num, size, filters, list)
}

func ListInfoTx(tx *sql.Tx, logger *log.Logger, tableName string, sortAsc int, sortField string, num, size int, filters SqlFilterMap, list interface{}) (count int, err error) {
	return listInfo(tx.Query, logger, tableName, sortAsc, sortField, num, size, filters, list)
}

func filterListInfo(query query, logger *log.Logger, tableName string, filters SqlFilterMap, list interface{}) (count int, err error) {
	t := reflect.TypeOf(list)
	if t.Kind() != reflect.Ptr {
		panic("list kind is not slice pointer")
	}
	t = t.Elem()
	if t.Kind() != reflect.Slice {
		panic("list kind is not slice pointer")
	}
	rv := reflect.ValueOf(list).Elem()
	where, args := filters.WhereSql()
	var rows *sql.Rows
	sqlStr := fmt.Sprintf("select %s from %s Where 1 = 1 %s ",
		MakeSqlColumns(reflect.New(t.Elem()).Interface()), tableName, where)
	if rows, err = query(sqlStr, args...); err == nil {
		defer func() {
			if err := rows.Close(); err != nil {
				logger.Println(err)
			}
		}()
		for rows.Next() {
			item := reflect.New(t.Elem())
			value := item.Interface()
			err = ScanDao(value, rows.Scan)
			if err != nil {
				return
			}
			count++
			rv = reflect.Append(rv, item.Elem())
		}
	}

	reflect.ValueOf(list).Elem().Set(rv)
	return
}

func FilterListInfo(db *sql.DB, logger *log.Logger, tableName string, filters SqlFilterMap, list interface{}) (count int, err error) {
	return filterListInfo(db.Query, logger, tableName, filters, list)
}

func FilterListInfoTx(tx *sql.Tx, logger *log.Logger, tableName string, filters SqlFilterMap, list interface{}) (count int, err error) {
	return filterListInfo(tx.Query, logger, tableName, filters, list)
}

func listDetailInfo(query query, logger *log.Logger, tableName string, foreignKey string, foreignValue interface{}, list interface{}) (count int, err error) {
	t := reflect.TypeOf(list)
	if t.Kind() != reflect.Ptr {
		panic("list kind is not slice pointer")
	}
	t = t.Elem()
	if t.Kind() != reflect.Slice {
		panic("list kind is not slice pointer")
	}
	rv := reflect.ValueOf(list).Elem()
	var rows *sql.Rows
	sqlStr := fmt.Sprintf("select %s from %s Where %s = ? ",
		MakeSqlColumns(reflect.New(t.Elem()).Interface()), tableName, SnakeField(foreignKey))
	if rows, err = query(sqlStr, foreignValue); err == nil {
		defer func() {
			if err := rows.Close(); err != nil {
				logger.Println(err)
			}
		}()
		for rows.Next() {
			item := reflect.New(t.Elem())
			value := item.Interface()
			err = ScanDao(value, rows.Scan)
			if err != nil {
				return
			}
			count++
			rv = reflect.Append(rv, item.Elem())
		}
	}

	reflect.ValueOf(list).Elem().Set(rv)
	return
}

func ListDetailInfo(db *sql.DB, logger *log.Logger, tableName string, foreignKey string, foreignValue interface{}, list interface{}) (count int, err error) {
	return listDetailInfo(db.Query, logger, tableName, foreignKey, foreignValue, list)
}

func ListDetailInfoTx(tx *sql.Tx, logger *log.Logger, tableName string, foreignKey string, foreignValue interface{}, list interface{}) (count int, err error) {
	return listDetailInfo(tx.Query, logger, tableName, foreignKey, foreignValue, list)
}
