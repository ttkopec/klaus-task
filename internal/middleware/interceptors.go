package middleware

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func ErrorLoggingInterceptor(logger *slog.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		resp, err := handler(ctx, req)

		if err != nil {
			st, _ := status.FromError(err)
			logger.Error(
				fmt.Sprintf("error encountered: %v", st.Message()),
				slog.String("method", info.FullMethod),
				slog.String("code", st.Code().String()),
			)
		}

		return resp, err
	}
}
