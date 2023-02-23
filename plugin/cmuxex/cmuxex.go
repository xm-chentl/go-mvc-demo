package cmuxex

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

type IntegratedServer struct {
	GinSrv  *gin.Engine
	GrpcSrv *grpc.Server
}

func (s *IntegratedServer) Run(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	cm := cmux.New(l)
	if s.GinSrv != nil {
		go s.GinSrv.RunListener(cm.Match(cmux.HTTP1Fast()))
	}
	if s.GrpcSrv != nil {
		go s.GrpcSrv.Serve(
			cm.Match(
				cmux.HTTP2HeaderField("content-type", "application/grpc"),
			),
		)
	}

	return cm.Serve()
}
