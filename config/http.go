package config

/**
 * @Author: feiliang.wang
 * @Description: HTTP配置
 * @File:  http
 * @Version: 1.0.0
 * @Date: 2020/8/4 5:47 下午
 */

type HttpConfig struct {
	Listen  int    `yaml:"listen"`
	FileDir string `yaml:"fileDir"`
}
