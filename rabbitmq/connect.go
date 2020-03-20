package rabbitmq

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

//
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

// Connect 建立链接
func Connect(conf *ConnectConfig) *amqp.Channel {
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
	return ch
}

// QueueExchange 交换机
type QueueExchange struct {
	QueueName    string        // 队列名称
	RoutingKey   string        // 路由键
	ExchangeName string        // 交换机名称
	ExchangeType string        // 交换机类型
	ch           *amqp.Channel // 队列链接
}

// 注册Queue
func (q *QueueExchange) Register(ch *amqp.Channel) error {
	if ch == nil {
		return errors.New("channel is nil")
	}
	if q.QueueName == "" {
		return errors.New("QueueName is empty")
	}
	// 注册交换机
	err := ch.ExchangeDeclarePassive(q.ExchangeName, q.ExchangeType, true, false, false, false, nil)
	if err != nil {
		// name:交换机名称; kind:交换机类型; durable:是否持久化; autoDelete:是否自动删除; internal:是否为内部; noWait:是否非阻塞; args:参数
		err = ch.ExchangeDeclare(q.ExchangeName, q.ExchangeType, true, false, false, false, nil)
		if err != nil {
			return errors.Wrap(err, "Failed to declare an exchange")
		}
	}
	// 创建队列
	//if _, err := ch.QueueDeclarePassive(q.QueueName, true, false, true, false, nil); err != nil {
	// name:队列名称; durable:是否持久化; autoDelete:是否自动删除; exclusive:是否设置排他; noWait:是否非阻塞; args:参数
	queue, err := ch.QueueDeclare("", true, false, true, false, nil)
	if err != nil {
		return errors.Wrap(err, "Failed to declare a queue")
	}
	q.QueueName = queue.Name
	//}
	// 将交换机与队列通过路由键绑定
	if err = ch.QueueBind(q.QueueName, q.RoutingKey, q.ExchangeName, false, nil); err != nil {
		return errors.Wrap(err, "Failed to bind a queue")
	}
	q.ch = ch
	return nil
}

// Publish 发送
func (q *QueueExchange) Publish(msg []byte) error {
	if q.ch == nil {
		return errors.New("no register")
	}
	// 发送任务消息
	if err := q.ch.Publish(q.ExchangeName, q.RoutingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        msg,
	}); err != nil {
		return errors.Wrap(err, "Publish")
	}
	return nil
}

// Receive 接收
func (q *QueueExchange) Receive(callback func([]byte) error) error {
	if q.ch == nil {
		return errors.New("no register")
	}
	err := q.ch.Qos(1, 0, true)
	if err != nil {
		return err
	}
	msgs, err := q.ch.Consume(q.QueueName, "", false, false, false, false, nil)
	if err != nil {
		return err
	}
	for msg := range msgs {
		if err := callback(msg.Body); err != nil {
			msg.Ack(true)
		} else {
			msg.Ack(false)
		}
	}
	return nil
}
