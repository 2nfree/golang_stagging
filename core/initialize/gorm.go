package initialize

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"golang-stagging/core"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// InitGormDB 初始化数据库
func InitGormDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&timeout=20s&loc=Local",
		core.Config.Database.Mysql.Username,
		core.Config.Database.Mysql.Password,
		core.Config.Database.Mysql.Hostname,
		core.Config.Database.Mysql.Database)
	mysqlConfig := mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), initGormConfig()); err != nil {
		return nil, err
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(core.Config.Database.Mysql.MaxIdleConn)
		sqlDB.SetMaxOpenConns(core.Config.Database.Mysql.MaxOpenConn)
		sqlDB.SetConnMaxLifetime(time.Duration(core.Config.Database.Mysql.MaxLifeTime) * time.Second)
		return db, nil
	}
}

// 配置gorm设置
func initGormConfig() *gorm.Config {
	return &gorm.Config{
		Logger:                                   InitGormLogger(), // 日志打印
		DisableForeignKeyConstraintWhenMigrating: true,             // 创建表时忽略外键
	}
}

// InitTables 初始化表格
func InitTables() error {
	ctx := context.Background()
	if ctx.Value("traceID") == nil {
		ctx = context.WithValue(ctx, "traceID", uuid.New().String())
	}
	return core.DB.WithContext(ctx).Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate()
}
