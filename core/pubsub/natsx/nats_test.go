package natsx

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"testing"
	"time"
)

var subject = "go-one.natsx.test"
var data = "hello world"

func TestSubPub(t *testing.T) {

	conn, err := nats.Connect("natsx://127.0.0.1:4222", nats.Name(t.Name()))
	if err != nil {
		t.Error(err)
	}

	sub(t, conn, "test1")

	msg, err := conn.RequestMsgWithContext(
		context.Background(),
		&nats.Msg{
			Subject: subject,
			Header: map[string][]string{
				"hello": {"hello"},
			},
			Data: []byte(data),
		})
	if err != nil {
		t.Error(err)
	}

	fmt.Println("publish reply", string(msg.Data))
	time.Sleep(10 * time.Second)
}

func sub(t *testing.T, conn *nats.Conn, name string) {
	_, err := conn.Subscribe(subject, func(msg *nats.Msg) {
		log.Println("-------name---------", name)
		log.Println("-------msg----------", string(msg.Data))
		err := msg.Ack()
		if err != nil {
			t.Error(err)
		}
	})
	if err != nil {
		t.Error(err)
	}

}
