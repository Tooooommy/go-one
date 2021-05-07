package user

import (
	"context"
)

type User struct {
}

func (u *User) Ping(ctx context.Context, in *Request) (*Response, error) {
	return &Response{
		Pong: in.Ping,
	}, nil
}

func (u *User) Pong(ctx context.Context, in *Request) (*Response, error) {
	return &Response{
		Pong: in.Ping,
	}, nil
}
