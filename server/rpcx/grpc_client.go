package rpcx

import (
	"github.com/Tooooommy/go-one/core/discov"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd/etcdv3"
	"sync"
)

// GrpcClient -> client
type GrpcClient struct {
	etcd *discov.Etcd
	insm sync.Map
}

// NewGrpcClient
func NewGrpcClient() *GrpcClient {
	return &GrpcClient{}
}

// EnableEtcd
func (c *GrpcClient) EnableEtcd(cfg discov.Config) {
	c.etcd = discov.NewEtcd(cfg)
}

// 加上读写时
func (c *GrpcClient) GetIns(prefix string) (*etcdv3.Instancer, error) {
	if val, ok := c.insm.Load(prefix); ok {
		return val.(*etcdv3.Instancer), nil
	} else {
		ins, err := c.etcd.NewInstancer(prefix)
		if err != nil {
			return nil, err
		}
		c.insm.Store(prefix, ins)
		return ins, err
	}
}

func (c *GrpcClient) Endpoints(prefix string, factory discov.EndpointFactory) endpoint.Endpoint {
	ins, err := c.GetIns(prefix)
	if err != nil {
	}
	return c.etcd.Endpoints(ins, factory, 0, 0)
}
