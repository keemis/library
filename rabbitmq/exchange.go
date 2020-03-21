package rabbitmq

import (
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

const ExchangeTypeFanout = "fanout"
const ExchangeTypeDirect = "direct" // 一条消息仅一个Queue消费
const ExchangeTypeTopic = "topic"   // 一条消息可多个Queue消费
const ExchangeTypeHeaders = "headers"

// Exchange 交换机
type Exchange struct {
	exchangeName string
	ch           *amqp.Channel
}

// ExchangeDeclare 声明交换机
func (c *Conn) ExchangeDeclare(exchangeName string, exchangeType string) (*Exchange, error) {
	if c.ch == nil {
		return nil, errors.New("channel is nil")
	}
	err := c.ch.ExchangeDeclarePassive(exchangeName, exchangeType, true, false, false, false, nil)
	if err != nil {
		// name:交换机名称; kind:交换机类型; durable:是否持久化; autoDelete:是否自动删除; internal:是否为内部; noWait:是否非阻塞; args:参数
		err = c.ch.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to declare an exchange")
		}
	}
	return &Exchange{
		exchangeName: exchangeName,
		ch:           c.ch,
	}, nil
}
