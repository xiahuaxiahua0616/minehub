package grpc

import (
	"context"
	"time"

	apiv1 "github.com/xiahuaxiahua0616/minehub/pkg/api/apiserver/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Healthz 服务健康检查.
func (h *Handler) Healthz(ctx context.Context, rq *emptypb.Empty) (*apiv1.HealthzResponse, error) {
	return &apiv1.HealthzResponse{
		Status:    apiv1.ServiceStatus_Healthy,
		Timestamp: time.Now().Format(time.DateTime),
	}, nil
}
