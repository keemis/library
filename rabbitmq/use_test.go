package rabbitmq

import (
	"fmt"
	"testing"
	"time"
)

// TestPublish 发送  (Connect -> Exchange -> Publish)
func TestPublish(t *testing.T) {
	conn := Connect(nil)
	exchange, err := conn.ExchangeDeclare("test.exchange.name", ExchangeTypeTopic)
	if err != nil {
		t.Logf(err.Error())
		return
	}
	for i := 0; i < 10000; i++ {
		msg := []byte(fmt.Sprintf("hello simon: %v", time.Now().Unix()))
		routingKey := "test.routing.key"
		err = exchange.Publish(routingKey, msg)
		t.Logf("Publish: %v,  err: %v", string(msg), err)
		time.Sleep(750 * time.Millisecond)
	}
}

// TestReceive 接收  (Connect -> Exchange -> Queue -> Receive)
func TestReceive(t *testing.T) {
	conn := Connect(nil)
	exchange, err := conn.ExchangeDeclare("test.exchange.name", ExchangeTypeTopic)
	if err != nil {
		t.Logf(err.Error())
		return
	}
	routingKeys := []string{"test.routing.key", "test.routing.key222"}
	queue, err := exchange.QueueBind(routingKeys)
	if err != nil {
		t.Logf(err.Error())
		return
	}
	err = queue.Receive(func(msg []byte) error {
		t.Logf("receive data: %v", string(msg))
		return nil
	})
	t.Logf("Receive err: %v", err)
}
