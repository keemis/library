package mysql

import (
	"errors"
	"fmt"
	"sync"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql" //gorm mysql驱动
	"github.com/jinzhu/gorm"
)

var (
	// 全局链接
	connects struct {
		store map[string]*gorm.DB
		sync.RWMutex
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
		User:     beego.AppConfig.DefaultString("MysqlUser", "root"),
		Password: beego.AppConfig.DefaultString("MysqlPassword", "123456"),
		Host:     beego.AppConfig.DefaultString("MysqlHost", "127.0.0.1:3306"),
		DbName:   beego.AppConfig.DefaultString("MysqlDbName", "test"),
		MaxConn:  beego.AppConfig.DefaultInt("MysqlMaxConn", 1000),
		MaxIdle:  beego.AppConfig.DefaultInt("MysqlIdleConn", 20),
	}
}

// Init 初始化
func Init(confs ...Conf) {
	if len(confs) == 0 {
		conf := defaultConf()
		confs = append(confs, conf)
	}
	connects.Lock()
	if connects.store == nil {
		connects.store = make(map[string]*gorm.DB)
	}
	connects.Unlock()
	for _, conf := range confs {
		db, err := connect(conf)
		if err != nil {
			panic(err)
		}
		connects.Lock()
		connects.store[conf.DbName] = db
		connects.Unlock()
	}
}

// 建立链接
func connect(conf Conf) (*gorm.DB, error) {
	if conf.DbName == "" || conf.Host == "" {
		return nil, errors.New("connect params is empty")
	}
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
