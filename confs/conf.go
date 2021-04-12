package confs

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

//初始化全局DB
var DB *gorm.DB

func SetUp() {
	// 初始化配置
	cErr := initConfig()
	if cErr != nil {
		fmt.Println(cErr)
	}

	//初始化数据库
	dbErr := initDb()
	if dbErr != nil {
		fmt.Println(dbErr)
	}
}

func initConfig() error {
	viper.SetConfigFile("./app.yaml") // 如果指定了配置文件，则解析指定的配置文件
	viper.SetConfigType("yaml")       // 设置配置文件格式为YAML
	if err := viper.ReadInConfig(); err != nil {
		fmt.Print(err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改:", in.Name)
	})
	return nil
}

func initDb() (err error) {
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	dbName := viper.GetString("mysql.db")
	local := viper.GetString("mysql.host") + ":" + viper.GetString("mysql.port")
	dsn := user + ":" + password + "@tcp(" + local + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	sqlDB, _ := DB.DB()
	maxIdleConn := viper.GetInt("mysql.db")
	maxOpenConn := viper.GetInt("mysql.db")
	if maxIdleConn <= 0 {
		maxIdleConn = 100
	}
	if maxOpenConn <= 0 {
		maxOpenConn = 10
	}
	//设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(maxIdleConn)
	//设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(maxOpenConn)
	//设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}
