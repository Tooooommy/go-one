package main

import (
	user "github.com/Tooooommy/go-one/example/hello_rpc"
	"google.golang.org/grpc"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	user.RegisterUserServer(s, &user.User{})
	s.Serve(l)
}
