package messagebus

import (
	"fmt"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	// 初始化MessageBus
	Init(50)
	// 订阅topic
	a := MessageSubscribe("Error_Key")
	_ = EmailSubscribe("Error_Key")
	// 发布topic
	Client().Publish("Error_Key", "hello")
	Client().Publish("Error_Key", "world")
	Client().Publish("Error_Key", "simon")

	// 取消订阅
	_ = Client().Unsubscribe("Error_Key", a.HandleEvent)
	// 发布topic
	Client().Publish("Error_Key", "php")

	time.Sleep(time.Hour)
}

// ==============Message=====================

type Message struct {
}

func MessageSubscribe(key string) *Message {
	o := &Message{}
	Client().Subscribe(key, o.HandleEvent)
	return o
}

func (d *Message) HandleEvent(data interface{}) {
	fmt.Println("Message HandleEvent, send message: ", data)
}

// ================Email===================

type Email struct {
}

func EmailSubscribe(key string) *Email {
	o := &Email{}
	Client().Subscribe(key, o.HandleEvent)
	return o
}

func (d *Email) HandleEvent(data interface{}) {
	fmt.Println("Email HandleEvent, send email: ", data)
}
