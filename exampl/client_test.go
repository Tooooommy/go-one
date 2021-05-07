package main

import (
	"context"
	"fmt"
	user "github.com/Tooooommy/go-one/exampl/hello_rpc"
	"google.golang.org/grpc"
	"testing"
)

func TestClient(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := user.NewUserClient(conn)
	resp, err := c.Ping(context.Background(), &user.Request{Ping: "pong"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Pong)

}
