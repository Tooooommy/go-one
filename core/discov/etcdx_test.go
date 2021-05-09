package discov

import (
	"context"
	"fmt"
	"github.com/Tooooommy/go-one/core/logx"
	"github.com/go-kit/kit/sd/etcdv3"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	options := etcdv3.ClientOptions{
		DialTimeout:   5 * time.Second,
		DialKeepAlive: 5 * time.Second,
	}
	var host = "0.0.0.0:2479"
	client, err := etcdv3.NewClient(context.Background(), []string{host}, options)
	if err != nil {
		panic(err)
	}
	register := etcdv3.NewRegistrar(client, etcdv3.Service{
		Key:   "go-one",
		Value: host,
		TTL:   nil,
	}, logx.KitL())
	register.Register()
	defer register.Deregister()
	fmt.Println(client.LeaseID())
	time.Sleep(1000 * time.Second)
}

func TestClient(t *testing.T) {

}
