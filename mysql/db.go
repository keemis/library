package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/keemis/library/logs"
)

type Orm struct {
	log logs.Logger
}

// Option 配置选项
type Option func(*Orm)

// SetLogger 设置日志对象
func SetLogger(log logs.Logger) Option {
	return func(orm *Orm) {
		orm.log = log
	}
}

// New 创建对象
func New(options ...Option) Orm {
	orm := Orm{}
	for _, option := range options {
		option(&orm)
	}
	return orm
}

// DB 返回一个DB对象
func (o Orm) DB(DbName ...string) *gorm.DB {
	db := &gorm.DB{}
	lens := len(DbName)
	if lens == 0 {
		for _, v := range connects.store {
			db = v
		}
	} else if lens == 1 {
		tmp, ok := connects.store[DbName[0]]
		if !ok {
			return nil
		}
		db = tmp
	} else {
		return nil
	}
	db.SetLogger(Logger{log: o.log})
	return db
}
