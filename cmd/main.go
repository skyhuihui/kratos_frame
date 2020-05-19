package main

import (
	"flag"
	"kratos_frame/internal/di"
	"kratos_frame/internal/err_code"
	"kratos_frame/internal/pkg/oss"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bilibili/kratos/pkg/conf/env"
	"github.com/bilibili/kratos/pkg/net/trace/zipkin"
	xtime "github.com/bilibili/kratos/pkg/time"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
)

func main() {
	flag.Set("conf", "./configs")
	flag.Parse()

	log.Init(&log.Config{
		Stdout: true,
		Dir:    os.Getenv("GOPATH") + "/src/kratos_frame/runtime/log/default",
	})
	defer log.Close()
	log.Info("kratos_frame start")

	// 注册配置文件模块
	paladin.Init()

	// 注册code码
	err_code.NewECode()

	// OSS
	if err := oss.New(); err != nil {
		panic(err)
	}

	// 注册链路跟踪 http 默认有链路跟踪的拦截 参考文档：https://book.eddycjy.com/golang/grpc/zipkin.html
	env.AppID = "kratos_frame"
	zipkin.Init(&zipkin.Config{
		Endpoint:      "http://127.0.0.1:9411/api/v2/spans", // docker run -d -p 9411:9411 openzipkin/zipkin
		BatchSize:     1000,
		Timeout:       xtime.Duration(2 * time.Second),
		DisableSample: true,
	})

	//di依赖注入操作
	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	stop(closeFunc)
}

func stop(closeFunc func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Info("kratos_frame exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
