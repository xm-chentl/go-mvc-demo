package cmuxex

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

type IntegratedServer struct {
	GrpcS *grpc.Server
	GinS  *gin.Engine
}

func (s *IntegratedServer) Run(addr string) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	cm := cmux.New(l)
	if s.GinS != nil {
		go func() {
			if err := s.GinS.RunListener(cm.Match(cmux.HTTP1Fast())); err != nil {
				log.Fatal("gin srv err: ", err)
			}
		}()
	}
	if s.GrpcS != nil {
		go func() {
			err := s.GrpcS.Serve(
				cm.Match(
					cmux.HTTP2HeaderField("content-type", "application/grpc"),
				),
			)
			if err != nil {
				log.Fatal("grpc srv err: ", err)
			}
		}()
	}

	go func() {
		if err := cm.Serve(); err != nil {
			log.Fatal("cmx server err: ", err)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	fmt.Println("Close the service countdown...5s")

	<-ctx.Done()
	cm.Close()
	fmt.Println("Shutdown Server ...")
}
