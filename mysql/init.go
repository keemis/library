package mysql

import (
	"fmt"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql" //gorm mysql驱动
	"github.com/jinzhu/gorm"
)

var (
	connects struct {
		store map[string]*gorm.DB
	}
)

// Conf 链接配置
type Conf struct {
	User     string
	Password string
	Host     string
	DbName   string
	MaxConn  int
	MaxIdle  int
}

// 默认单库配置
func defaultConf() Conf {
	return Conf{
		User:     beego.AppConfig.String("MysqlUser"),
		Password: beego.AppConfig.String("MysqlPassword"),
		Host:     beego.AppConfig.String("MysqlHost"),
		DbName:   beego.AppConfig.String("MysqlDbName"),
		MaxConn:  beego.AppConfig.DefaultInt("MysqlMaxConn", 1000),
		MaxIdle:  beego.AppConfig.DefaultInt("MysqlIdleConn", 20),
	}
}

// Init 初始化
func Init(configs ...Conf) {
	if len(configs) == 0 {
		config := defaultConf()
		configs = append(configs, config)
	}
	connects.store = make(map[string]*gorm.DB)
	for _, conf := range configs {
		db, err := conn(conf)
		if err != nil {
			panic(err)
		}
		connects.store[conf.DbName] = db
	}
}

// 建立链接
func conn(conf Conf) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.Host, conf.DbName)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxOpenConns(conf.MaxConn)
	db.DB().SetMaxIdleConns(conf.MaxIdle)
	if err := db.DB().Ping(); err != nil {
		return nil, err
	}
	db.SingularTable(true)
	db.LogMode(true)
	return db, nil
}
