package config

/**
 * @Author: feiliang.wang
 * @Description: 配置
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2020/8/4 5:37 下午
 */

type Config struct {
	Mysql MysqlConfig `yaml:"mysql"`
	Http  HttpConfig  `yaml:"http"`
}
