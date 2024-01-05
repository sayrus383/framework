package validator

import (
	"context"
	"google.golang.org/grpc"
)

// UnaryValidationInterceptor - middleware для проверки валидации в GRPC (для однонаправленных запросов)
func UnaryValidationInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if v, ok := req.(allValidator); ok {
			if err := v.ValidateAll(); err != nil {
				return nil, err
			}
		}
		if v, ok := req.(validator); ok {
			if err := v.Validate(); err != nil {
				return nil, err
			}
		}

		return handler(ctx, req)
	}
}

// StreamValidationInterceptor - middleware для проверки валидации в GRPC (для двунаправленных запросов)
func StreamValidationInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		wrapper := &recvWrapper{stream}
		return handler(srv, wrapper)
	}
}

type recvWrapper struct {
	grpc.ServerStream
}

func (s *recvWrapper) RecvMsg(req interface{}) error {
	if err := s.ServerStream.RecvMsg(req); err != nil {
		return err
	}

	if v, ok := req.(allValidator); ok {
		if err := v.ValidateAll(); err != nil {
			return err
		}
	}
	if v, ok := req.(validator); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}
