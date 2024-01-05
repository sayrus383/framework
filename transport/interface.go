package transport

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type GrpcOption func(o *grpcOptions)

type RegisterGrpcServer func(s *grpc.Server)
type RegisterHttpHandler func(m *runtime.ServeMux, c *grpc.ClientConn)

type GrpcServer interface {
	WithOptions(options ...GrpcOption)
	GetServiceInfo() map[string]grpc.ServiceInfo
}
