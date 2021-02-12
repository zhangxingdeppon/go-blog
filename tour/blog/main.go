package main

import (
	"blog/pkg/tracer"
	"blog/global"
	"blog/internal/model"
	"blog/internal/routers"
	"blog/pkg/logger"
	"blog/pkg/setting"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	err := setupSetting()
	log.Printf("hello world")
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}
// @title 博客系统
// @version 1.0
// @description Go Language travel
// @termsOfService https://github.com
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("config: ", global.ServerSetting, global.AppSetting, global.DatabaseSetting)
	global.Logger.Infof("%s: go-programming-tour-book/%s", "bryan", "blog")
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil

}
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	filename := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  filename,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true}, "", log.LstdFlags).WithCaller(2)
	return nil
}
func setupTracer() error {
     jaegerTracer, _, err := tracer.NewJaegerTracer(
		 "blog-service",
		 "192.168.0.107:6831",
	 )
	 if err != nil {
		 return  err
	 }
	 global.Tracer = jaegerTracer
	 return nil	
}