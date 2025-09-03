package apiserver

import (
	"net"
	"time"

	handler "github.com/xiahuaxiahua0616/minehub/internal/apiserver/handler/grpc"
	"github.com/xiahuaxiahua0616/minehub/internal/pkg/log"
	apiv1 "github.com/xiahuaxiahua0616/minehub/pkg/api/apiserver/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	genericoptions "github.com/onexstack/onexstack/pkg/options"
)

const (
	// GRPCServerMode 定义 gRPC 服务模式.
	// 使用 gRPC 框架启动一个 gRPC 服务器.
	GRPCServerMode = "grpc"
	// GRPCServerMode 定义 gRPC + HTTP 服务模式.
	// 使用 gRPC 框架启动一个 gRPC 服务器 + HTTP 反向代理服务器.
	GRPCGatewayServerMode = "grpc-gateway"
	// GinServerMode 定义 Gin 服务模式.
	// 使用 Gin Web 框架启动一个 HTTP 服务器.
	GinServerMode = "gin"
)

// Config 配置结构体，用于存储应用相关的配置。
// 不用 viper.Get，是因为这种方式能更加清晰的知道应用提供了哪些配置项。
type Config struct {
	ServerMode  string
	JWTKey      string
	Expiration  time.Duration
	GRPCOptions *genericoptions.GRPCOptions
}

// UnionServer 定义一个联合服务器。 根据 ServerMode 决定要启动的服务器类型。
type UnionServer struct {
	cfg *Config
	srv *grpc.Server
	lis net.Listener
}

// NewUnionServer 根据配置创建联合服务器。
func (cfg *Config) NewUnionServer() (*UnionServer, error) {
	lis, err := net.Listen("tcp", cfg.GRPCOptions.Addr)
	if err != nil {
		log.Fatalw("Failed to listen", "err", err)
		return nil, err
	}

	// 创建 GRPC Server 实例
	grpcsrv := grpc.NewServer()
	apiv1.RegisterMineHubServer(grpcsrv, handler.NewHandler())
	reflection.Register(grpcsrv)

	return &UnionServer{cfg: cfg, srv: grpcsrv, lis: lis}, nil
}

// Run 运行应用。
func (s *UnionServer) Run() error {
	// 打印一条日志，用来提示 GRPC 服务已经起来，方便排障
	log.Infow("Start to listening the incoming requests on grpc address", "addr", s.cfg.GRPCOptions.Addr)

	s.srv.Serve(s.lis)

	select {}
}
