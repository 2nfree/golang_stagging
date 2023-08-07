package core

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"golang-stagging/core/config"
	"gorm.io/gorm"
)

var (
	Config        *config.Config     //全局配置
	DB            *gorm.DB           //全局数据库
	RDB           *redis.Client      //全局redis连接池
	Logger        *zap.Logger        //全局日志
	SugaredLogger *zap.SugaredLogger //全局日志Sugared
)
