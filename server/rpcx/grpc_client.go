package rpcx

import (
	"github.com/Tooooommy/go-one/core/discov"
	"github.com/go-kit/kit/endpoint"
	grpctranspot "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"io"
	"sync"
)

// GrpcClient -> client
type (
	GrpcClient struct {
		etcd *discov.Etcd
		insm sync.Map
	}

	EncodeRequest  grpctranspot.EncodeRequestFunc
	DecodeResponse grpctranspot.DecodeResponseFunc
)

// NewGrpcClient
func NewGrpcClient(cfg discov.Config) (*GrpcClient, error) {
	cli, err := discov.NewEtcd(cfg)
	if err != nil {
		return nil, err
	}
	return &GrpcClient{etcd: cli}, nil
}

// 加上读写时
func (c *GrpcClient) GetInvoker(prefix string) (*discov.Invoker, error) {
	if val, ok := c.insm.Load(prefix); ok {
		return val.(*discov.Invoker), nil
	} else {
		ins, err := c.etcd.NewInvoker(prefix)
		if err != nil {
			return nil, err
		}
		c.insm.Store(prefix, ins)
		return ins, err
	}
}
