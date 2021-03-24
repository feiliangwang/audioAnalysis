package config

/**
 * @Author: feiliang.wang
 * @Description: 数据库配置
 * @File:  mysql
 * @Version: 1.0.0
 * @Date: 2020/8/4 5:26 下午
 */

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	User     string `yaml:"user"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	Charset  string `yaml:"charset"`
}

func (c MysqlConfig) OpenDb() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", c.User, c.Password, c.Host, c.Port, c.Name, c.Charset)
	return sql.Open("mysql", dsn)
}
