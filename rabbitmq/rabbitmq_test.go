package rabbitmq

import (
	"fmt"
	"testing"
	"time"
)

func TestPublish(t *testing.T) {
	ch := Connect(nil)
	exchange := &QueueExchange{
		QueueName:    "test.queue.name",
		RoutingKey:   "test.routing.key",
		ExchangeName: "test.exchange.name",
		ExchangeType: "topic",
	}
	err := exchange.Register(ch)
	t.Logf("Register: %v", err)
	if err != nil {
		return
	}
	for i := 0; i < 10000; i++ {
		err = exchange.Publish([]byte(fmt.Sprintf("hello simon: %v", time.Now().Unix())))
		t.Logf("Send: %v", err)
		time.Sleep(time.Second)
	}
}

func TestReceive(t *testing.T) {
	ch := Connect(nil)
	exchange := &QueueExchange{
		QueueName:    "test.queue.name",
		RoutingKey:   "test.routing.key",
		ExchangeName: "test.exchange.name",
		ExchangeType: "topic",
	}
	err := exchange.Register(ch)
	t.Logf("Register: %v", err)
	if err != nil {
		return
	}
	err = exchange.Receive(func(msg []byte) error {
		t.Logf("receive data: %v", string(msg))
		return nil
	})
	t.Logf("Receive: %v", err)
}
