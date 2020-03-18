package messagebus

import (
	"fmt"
	"testing"
	"time"
)

var bus = New(100)

func TestRun(t *testing.T) {
	// 订阅topic
	a := NewMessage("Error_Key")
	_ = NewEmail("Error_Key")
	// 发布topic
	bus.Publish("Error_Key", "hello")
	bus.Publish("Error_Key", "world")
	bus.Publish("Error_Key", "simon")

	// 取消订阅
	_ = bus.Unsubscribe("Error_Key", a.HandleEvent)
	// 发布topic
	bus.Publish("Error_Key", "php")

	time.Sleep(time.Hour)
}

// ==============Message=====================

type Message struct {
}

func NewMessage(key string) *Message {
	o := &Message{}
	bus.Subscribe(key, o.HandleEvent)
	return o
}

func (d *Message) HandleEvent(data interface{}) {
	fmt.Println("Message HandleEvent, send message: ", data)
}

// ================Email===================

type Email struct {
}

func NewEmail(key string) *Email {
	o := &Email{}
	bus.Subscribe(key, o.HandleEvent)
	return o
}

func (d *Email) HandleEvent(data interface{}) {
	fmt.Println("Email HandleEvent, send email: ", data)
}
