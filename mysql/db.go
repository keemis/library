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

// OptLogger 设置日志对象
func OptLogger(log logs.Logger) Option {
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
	if len(connects.store) == 0 {
		return nil
	}
	db := &gorm.DB{}
	lens := len(DbName)
	if lens == 0 {
		for _, v := range connects.store {
			db = v
		}
	} else if lens == 1 {
		name := DbName[0]
		if name == "" {
			return nil
		}
		connects.RLock()
		tmp, ok := connects.store[name]
		connects.RUnlock()
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
