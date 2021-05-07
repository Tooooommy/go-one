package rpcx

import (
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
	address            string
	options            []grpc.ServerOption
	streamInterceptors []grpc.StreamServerInterceptor
	unaryInterceptors  []grpc.UnaryServerInterceptor
}

func NewGrpcServer(address string) *GrpcServer {
	return &GrpcServer{
		address: address,
	}
}

func (s *GrpcServer) WithOption(option ...grpc.ServerOption) {
	s.options = append(s.options, option...)
}

func (s *GrpcServer) UseStreamInterceptor(interceptor ...grpc.StreamServerInterceptor) {
	s.streamInterceptors = append(s.streamInterceptors, interceptor...)
}

func (s *GrpcServer) UseUnaryInterceptors(interceptor ...grpc.UnaryServerInterceptor) {
	s.unaryInterceptors = append(s.unaryInterceptors, interceptor...)
}

func (s *GrpcServer) Start() error {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	s.options = append(s.options, grpc.ChainUnaryInterceptor(s.unaryInterceptors...))
	s.options = append(s.options, grpc.ChainStreamInterceptor(s.streamInterceptors...))
	server := grpc.NewServer(s.options...)
	defer server.GracefulStop() // TODO: 可能优化
	return server.Serve(lis)
}
