package rabbitmq

import (
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

// Publish 发送
func (e *Exchange) Publish(routingKey string, msg []byte) error {
	if e.exchangeName == "" {
		return errors.New("ExchangeDeclare is nil")
	}
	// 发送任务消息
	if err := e.ch.Publish(e.exchangeName, routingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        msg,
	}); err != nil {
		return errors.Wrap(err, "Publish")
	}
	return nil
}
