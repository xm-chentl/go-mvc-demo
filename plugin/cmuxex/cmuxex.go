package cmuxex

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

type IntegratedServer struct {
	GrpcS *grpc.Server
	GinS  *gin.Engine
}

func (s *IntegratedServer) Run(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	cm := cmux.New(l)
	if s.GinS != nil {
		go s.GinS.RunListener(cm.Match(cmux.HTTP1Fast()))
	}
	if s.GrpcS != nil {
		go s.GrpcS.Serve(
			cm.Match(
				cmux.HTTP2HeaderField("content-type", "application/grpc"),
			),
		)
	}

	return cm.Serve()
}
