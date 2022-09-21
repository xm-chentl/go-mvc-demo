package grpcsvc

import (
	"context"
	"fmt"

	"github.com/xm-chentl/go-mvc-demo/proto/go/userservice"
)

type userServiceImpl struct {
	userservice.UnimplementedUserServiceServer
}

func (s *userServiceImpl) Create(ctx context.Context, req *userservice.CreateRequest) (resp *userservice.CreateResponse, err error) {
	resp = &userservice.CreateResponse{}
	fmt.Println("exec user.create ok")
	return
}

func NewUserService() userservice.UserServiceServer {
	return &userServiceImpl{}
}
