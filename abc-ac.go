package service

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/zngue/zng_tool/api/test/v1"
)

import (
	"github.com/gin-gonic/gin"
)

type TestService struct {
}

func NewTestService() v1.TestGinHttpService {
	return &TestService{}
}
func (s *TestService) Info(ctx *gin.Context, req *v1.GetTestRequest) (rs *v1.GetTestReply, err error) {
	return
}
func (s *TestService) List(ctx *gin.Context, req *v1.GetListTestRequest) (rs *v1.GetTestReply, err error) {
	return
}
func (s *TestService) Create(ctx *gin.Context, req *v1.GetTestRequest) (rs *v1.GetTestReply, err error) {
	return
}
func (s *TestService) ListUser(ctx *gin.Context, req *v1.GetListTestRequest) (rs *v1.UserList, err error) {
	return
}
func (s *TestService) ListUser2(ctx *gin.Context, req *v1.GetList2TestRequest) (rs *v1.UserList2Rely, err error) {
	return
}
func (s *TestService) UpdateUser(ctx *gin.Context, req *v1.UpdateUserRequest) (rs *empty.Empty, err error) {
	return
}
