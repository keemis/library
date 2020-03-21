package rabbitmq

import (
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

// Queue 队列
type Queue struct {
	queueName string
	ch        *amqp.Channel
}

// QueueBind 绑定到队列
func (e *Exchange) QueueBind(routingKeys []string) (*Queue, error) {
	if e.exchangeName == "" {
		return nil, errors.New("ExchangeDeclare is nil")
	}
	// name:队列名称; durable:是否持久化; autoDelete:是否自动删除; exclusive:是否设置排他; noWait:是否非阻塞; args:参数
	queue, err := e.ch.QueueDeclare("", true, false, true, false, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to declare a queue")
	}
	// 将交换机与队列通过路由键绑定
	for _, routingKey := range routingKeys {
		if err = e.ch.QueueBind(queue.Name, routingKey, e.exchangeName, false, nil); err != nil {
			return nil, errors.Wrap(err, "Failed to bind a queue")
		}
	}
	return &Queue{
		queueName: queue.Name,
		ch:        e.ch,
	}, nil
}

// Receive 接收
func (q *Queue) Receive(callback func([]byte) error) error {
	if q.queueName == "" {
		return errors.New("QueueBind is nil")
	}
	if err := q.ch.Qos(1, 0, true); err != nil {
		return err
	}
	msgs, err := q.ch.Consume(q.queueName, "", false, false, false, false, nil)
	if err != nil {
		return err
	}
	for msg := range msgs {
		if err := callback(msg.Body); err != nil {
			q.ack(msg, true, 5)
		} else {
			q.ack(msg, false, 5)
		}
	}
	return nil
}

// ack 确认答复
func (q *Queue) ack(msg amqp.Delivery, multiple bool, times int) {
	if times <= 0 {
		return
	}
	if err := msg.Ack(multiple); err != nil {
		q.ack(msg, multiple, times-1)
	}
}
