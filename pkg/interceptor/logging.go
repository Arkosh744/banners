package interceptor

import (
	"context"
	"github.com/Arkosh744/banners/internal/log"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Debug("incoming GRPC request", zap.String("method", info.FullMethod), zap.Any("request", req))

	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		if sc, ok := span.Context().(jaeger.SpanContext); ok {
			log.Debug("tracing", zap.String("traceID", sc.TraceID().String()), zap.String("spanID", sc.SpanID().String()))
		}
	}

	res, err := handler(ctx, req)
	if err != nil {
		if span = opentracing.SpanFromContext(ctx); span != nil {
			ext.Error.Set(span, true)
		}

		log.Error(ctx, "Error handling GRPC request", zap.String("method", info.FullMethod), zap.Error(err))

		return nil, err
	}

	log.Debug("GRPC response", zap.String("method", info.FullMethod), zap.Any("response", res))

	return res, nil
}
