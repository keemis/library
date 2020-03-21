package rabbitmq

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/streadway/amqp"
)

// ConnectConfig 链接配置
type ConnectConfig struct {
	User     string
	Password string
	Host     string
	Port     int
}

// 默认配置
func defaultConfig() *ConnectConfig {
	return &ConnectConfig{
		User:     beego.AppConfig.DefaultString("RabbitMqUser", "guest"),
		Password: beego.AppConfig.DefaultString("RabbitMqPassword", "guest"),
		Host:     beego.AppConfig.DefaultString("RabbitMqHost", "127.0.0.1"),
		Port:     beego.AppConfig.DefaultInt("RabbitMqPort", 5672),
	}
}

// Conn 链接Channel
type Conn struct {
	ch *amqp.Channel
}

// Connect 建立链接
func Connect(conf *ConnectConfig) *Conn {
	if conf == nil {
		conf = defaultConfig()
	}
	rabbitUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", conf.User, conf.Password, conf.Host, conf.Port)
	conn, err := amqp.Dial(rabbitUrl)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return &Conn{ch: ch}
}
