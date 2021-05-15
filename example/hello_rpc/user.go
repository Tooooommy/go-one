package user

import (
	"context"
	"fmt"
)

type User struct {
}

func (u *User) Ping(ctx context.Context, in *Request) (*Response, error) {
	fmt.Printf("%+v\n", in)
	return &Response{
		Pong: in.Ping,
	}, nil
}

func (u *User) Pong(ctx context.Context, in *Request) (*Response, error) {
	fmt.Printf("%+v\n", in)
	return &Response{
		Pong: in.Ping,
	}, nil
}
