package grpcsvc

import (
	"context"
	"fmt"

	"github.com/xm-chentl/go-mvc-demo/proto/go/userservice"
	"github.com/xm-chentl/gocore/frame"
	"google.golang.org/grpc"
)

type userServiceImpl struct {
	userservice.UnimplementedUserServiceServer
}

func (s *userServiceImpl) Create(ctx context.Context, req *userservice.CreateRequest) (resp *userservice.CreateResponse, err error) {
	resp = &userservice.CreateResponse{}
	fmt.Println("exec user.create ok")
	return
}

func NewUserService1() userservice.UserServiceServer {
	return &userServiceImpl{}
}

func NewUserService() frame.GRPCOption {
	return func(g *grpc.Server) {
		userservice.RegisterUserServiceServer(g, &userServiceImpl{})
	}
}
