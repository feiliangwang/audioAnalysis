package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/feiliangwang/audioAnalysis/config"
	"github.com/feiliangwang/audioAnalysis/endpoint"
	"github.com/feiliangwang/audioAnalysis/plugins/logging"
	"github.com/feiliangwang/audioAnalysis/service"
	"github.com/feiliangwang/audioAnalysis/transport"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

/**
 * @Author: feiliang.wang
 * @Description: 主函数
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2021/3/23 22:25
 */

var (
	help bool
	cfg  string
)

func init() {
	flag.BoolVar(&help, "h", false, "this help")
	flag.StringVar(&cfg, "c", "config.yaml", "config file path,defalut is current dir config.yaml")
}

var ctx = context.Background()
var logger = log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	c := config.Config{}
	if bs, err := ioutil.ReadFile(cfg); err != nil {
		logger.Printf("read config file %s fail.%+v\n", cfg, err)
		os.Exit(1)
	} else if err := yaml.Unmarshal(bs, &c); err != nil {
		logger.Printf("read config is %s,but unmarshal fail.%+v\n", string(bs), err)
		os.Exit(1)
	} else if db, err := c.Mysql.OpenDb(); err != nil {
		logger.Printf("open db fail.%+v\n", err)
		os.Exit(1)
	} else {
		defer db.Close()

		r := mux.NewRouter()

		//音频服务
		audioServer := service.NewAudioServer(db, c.Http.FileDir)
		audioServer = logging.SkAppLoggingAudioServiceMiddleware(logger)(audioServer)
		audioEndpts := endpoint.AudioServiceEndpoints{
			AddEndpoint:    endpoint.MakeAudioAddEndpoint(audioServer, logger),
			DeleteEndpoint: endpoint.MakeAudioDeleteEndpoint(audioServer, logger),
			UpdateEndpoint: endpoint.MakeAudioUpdateEndpoint(audioServer, logger),
			DetailEndpoint: endpoint.MakeAudioDetailEndpoint(audioServer, logger),
			ListEndpoint:   endpoint.MakeAudioListEndpoint(audioServer, logger),
		}
		transport.MakeAudioHttpHandler(ctx, logger, r.PathPrefix("/api").Subrouter(), audioEndpts)
		r.PathPrefix("/").Handler(http.FileServer(http.Dir(c.Http.FileDir)))
		r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				http.Redirect(w, r, "/index.html", http.StatusFound)
			} else {
				http.Redirect(w, r, "/404.html", http.StatusFound)
			}
		})
		errChan := make(chan error)
		go func() {
			hostPort := fmt.Sprintf(":%d", c.Http.Listen)
			logger.Printf("http server start at %s\n", hostPort)
			errChan <- http.ListenAndServe(hostPort, r)
		}()

		go func() {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			errChan <- fmt.Errorf("%s", <-c)
		}()

		err := <-errChan
		logger.Printf("server exit with %v\n", err)
		os.Exit(0)
	}
}
