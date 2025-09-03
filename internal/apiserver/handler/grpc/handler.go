package grpc

import (
	apiv1 "github.com/xiahuaxiahua0616/minehub/pkg/api/apiserver/v1"
)

// Handler 负责处理博客模块的请求.
type Handler struct {
	apiv1.UnimplementedMineHubServer
}

// NewHandler 创建一个新的 Handler 实例.
func NewHandler() *Handler {
	return &Handler{}
}
