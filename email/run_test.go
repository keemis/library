package email

import (
	"net/smtp"
	"net/textproto"
	"testing"
	"time"

	"github.com/jordan-wright/email"
)

func TestRun(t *testing.T) {
	e := &email.Email{
		To:      []string{"test@example.com"},
		From:    "Jordan Wright <test@gmail.com>",
		Subject: "Awesome Subject",
		Text:    []byte("Text Body is, of course, supported!"),
		HTML:    []byte("<h1>Fancy HTML is supported, too!</h1>"),
		Headers: textproto.MIMEHeader{},
	}
	err := e.Send("smtp.126.com:25", smtp.PlainAuth("", "admin@126.com", "password", "smtp.126.com"))
	t.Logf(err.Error())
}

func TestRunPool(t *testing.T) {
	// pool
	var ch chan *email.Email
	p, err := email.NewPool(
		"smtp.126.com:25",
		5,
		smtp.PlainAuth("", "admin@126.com", "password", "smtp.126.com"),
	)
	if err != nil {
		t.Logf(err.Error())
		return
	}
	for i := 0; i < 5; i++ {
		go func() {
			for e := range ch {
				err := p.Send(e, 5*time.Second)
				t.Logf(err.Error())
			}
		}()
	}
	// send
	ch <- &email.Email{
		To:      []string{"test@example.com"},
		From:    "Jordan Wright <test@gmail.com>",
		Subject: "Awesome Subject",
		Text:    []byte("Text Body is, of course, supported!"),
		HTML:    []byte("<h1>Fancy HTML is supported, too!</h1>"),
		Headers: textproto.MIMEHeader{},
	}

	time.Sleep(time.Hour)
}
