package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang-stagging/core"
	"golang-stagging/core/initialize"
	"golang-stagging/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var configPath string
var err error

func init() {
	time.Local = time.FixedZone("CST", 8*3600)
	flag.StringVar(&configPath, "cfg", "config.toml", "/path/to/filename")
	flag.Parse()
	initialize.InitConfig(configPath)
	initialize.InitZapLogger()
	//core.DB, err = initialize.InitGormDB()
	//if err != nil {
	//	core.SugaredLogger.Panicf("初始化数据库失败: %v", err)
	//}
	//err := initialize.InitTables()
	//if err != nil {
	//	core.SugaredLogger.Panicf("初始化数据表失败: %v", err)
	//}
	//core.RDB, err = initialize.InitRedis()
	//if err != nil {
	//	core.SugaredLogger.Panicf("初始化Redis失败: %v", err)
	//}
}

func main() {
	gin.SetMode(core.Config.Server.RunMode)
	r := router.InitRouter()
	srv := &http.Server{
		Addr:    core.Config.Server.Listen,
		Handler: r,
	}
	core.SugaredLogger.Infof("listening and serving HTTP on %v", srv.Addr)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			core.SugaredLogger.Fatalf("listening on %v with error: %v", srv.Addr, err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	core.SugaredLogger.Warn("waiting to stop server......")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		core.SugaredLogger.Error("server regular stop failed, force quit", zap.Error(err))
	}
	core.SugaredLogger.Warn("server regular stop")
}
