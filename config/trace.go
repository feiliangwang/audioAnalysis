package config

import (
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

/**
 * @Author: feiliang.wang
 * @Description:
 * @File:  trace
 * @Version: 1.0.0
 * @Date: 2020/8/4 6:23 下午
 */

type TraceConf struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Url  string `yaml:"url"`
}

func (c TraceConf) GetTrace(serverName string, hostPort string) (reporter.Reporter, *zipkin.Tracer, error) {
	if ep, err := zipkin.NewEndpoint(serverName, hostPort); err != nil {
		return nil, nil, err
	} else {
		rep := zipkinhttp.NewReporter("http://" + c.Host + ":" + c.Port + c.Url)
		tracer, err := zipkin.NewTracer(rep, zipkin.WithLocalEndpoint(ep), zipkin.WithNoopTracer(false))
		return rep, tracer, err
	}
}
