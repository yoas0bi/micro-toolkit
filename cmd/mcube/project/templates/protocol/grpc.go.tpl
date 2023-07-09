package protocol

import (
	"net"

	"google.golang.org/grpc"

{{ if $.EnableKeyauth -}}
	"github.com/infraboard/keyauth/client/interceptor"
{{- end }}
	"github.com/yoas0bi/micro-toolkit/app"
	"github.com/yoas0bi/micro-toolkit/grpc/middleware/recovery"
	"github.com/yoas0bi/micro-toolkit/logger"
	"github.com/yoas0bi/micro-toolkit/logger/zap"

	"{{.PKG}}/conf"
)

// NewGRPCService todo
func NewGRPCService() *GRPCService {
	log := zap.L().Named("GRPC Service")

{{ if $.EnableKeyauth -}}
	c, err := conf.C().Keyauth.Client()
	if err != nil {
	panic(err)
	}
{{- end }}

	rc := recovery.NewInterceptor(recovery.NewZapRecoveryHandler())
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		rc.UnaryServerInterceptor(),
{{ if $.EnableKeyauth -}}
		interceptor.GrpcAuthUnaryServerInterceptor(c),
{{- end }}
	))

	return &GRPCService{
		svr: grpcServer,
		l:   log,
		c:   conf.C(),
	}
}

// GRPCService grpc服务
type GRPCService struct {
	svr *grpc.Server
	l   logger.Logger
	c   *conf.Config
}

// Start 启动GRPC服务
func (s *GRPCService) Start() {
	// 装载所有GRPC服务
	app.LoadGrpcApp(s.svr)

	// 启动HTTP服务
	lis, err := net.Listen("tcp", s.c.App.GRPC.Addr())
	if err != nil {
		s.l.Errorf("listen grpc tcp conn error, %s", err)
		return
	}

	s.l.Infof("GRPC 服务监听地址: %s", s.c.App.GRPC.Addr())
	if err := s.svr.Serve(lis); err != nil {
		if err == grpc.ErrServerStopped {
			s.l.Info("service is stopped")
		}

		s.l.Error("start grpc service error, %s", err.Error())
		return
	}
}

// Stop 启动GRPC服务
func (s *GRPCService) Stop() error {
	s.svr.GracefulStop()
	return nil
}