package transport

import (
	"gitlab.bcc.kz/digital-banking-platform/microservices/currency-market/modules/framework.git/logger"
	"google.golang.org/grpc"
)

type grpcOptions struct {
	streamInterceptors []grpc.StreamServerInterceptor
	unaryInterceptors  []grpc.UnaryServerInterceptor
	registerServers    []RegisterGrpcServer
}

type grpcServer struct {
	addr string
	log  logger.Logger

	options *grpcOptions
	server  *grpc.Server
}

//func NewGrpcServer(addr string, log logger.Logger) GrpcServer {
//	srv := &grpcServer{
//		addr: addr,
//		log:  log.With("transport", "grpc"),
//		//options: newGrpcOptions(),
//	}
//	srv.WithOptions(
//		WithStreamInterceptors(
//			validator.StreamValidationInterceptor(),
//		),
//		WithUnaryInterceptors(
//			validator.UnaryValidationInterceptor(),
//		),
//	)
//	return srv
//}

func (s *grpcServer) GetServiceInfo() map[string]grpc.ServiceInfo {
	return s.server.GetServiceInfo()
}

func (s *grpcServer) WithOptions(options ...GrpcOption) {
	for _, opt := range options {
		opt(s.options)
	}
}
