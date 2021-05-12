package etcdx

import (
	"context"
	"github.com/Tooooommy/go-one/core/zapx"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"google.golang.org/grpc"
	"io"
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
	}, zapx.KitL())
	register.Register()
	defer register.Deregister()
	time.Sleep(1000 * time.Second)

}

// 应该在Factory 函数调用 grpc.NewClient().Endpoint ->
// 最后的调用调用 grpctransport.NewClient().Endpoint --> 拿到实例 Server
// 最后直接
func TestClient(t *testing.T) {
	options := etcdv3.ClientOptions{
		DialTimeout:   5 * time.Second,
		DialKeepAlive: 5 * time.Second,
	}
	var host = "0.0.0.0:2479"
	client, err := etcdv3.NewClient(context.Background(), []string{host}, options)
	if err != nil {
		panic(err)
	}

	// 创建一个ir
	// 可以共用
	instancer, err := etcdv3.NewInstancer(client, "go-one", zapx.KitL())
	if err != nil {
		panic(err)
	}
	instancer.Stop()

	// 无法共用
	sd.NewEndpointer(instancer, factory, zapx.KitL())
}

//func MakeEndpoint(conn *grpc.ClientConn) endpoint.Endpoint {
//	return grpctransport.NewClient(
//		conn,
//		"",
//		"",
//		,
//
//	).Endpoint()
//}

func factory(instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc.Dial(instance, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	return nil, conn, err
}
