// Code generated by protoc-gen-gin-http. DO NOT EDIT.
package v1

import (
	empty "github.com/golang/protobuf/ptypes/empty"
)

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/zng_app/pkg/bind"
	
	"github.com/zngue/zng_app/pkg/router"
	"github.com/zngue/zng_app/pkg/bind"
)

// 服务操作
const OperationGinTestInfo = "api.test.v1.Test.Info"
const OperationGinTestList = "api.test.v1.Test.List"
const OperationGinTestCreate = "api.test.v1.Test.Create"
const OperationGinTestListUser = "api.test.v1.Test.ListUser"
const OperationGinTestListUser2 = "api.test.v1.Test.ListUser2"
const OperationGinTestUpdateUser = "api.test.v1.Test.UpdateUser"

// 服务url
const OperationGinUrlTestInfo = "/v1/test/info"
const OperationGinUrlTestList = "/v1/test/list"
const OperationGinUrlTestCreate = "/v1/test/create"
const OperationGinUrlTestListUser = "/v1/test/list_user"
const OperationGinUrlTestListUser2 = "/v1/test/list_user2"
const OperationGinUrlTestUpdateUser = "/v1/test/updateUser"

//服务接口abc

type TestGinHttpService interface {
	Info(ctx *gin.Context, req *GetTestRequest) (rs *GetTestReply, err error)
	List(ctx *gin.Context, req *GetListTestRequest) (rs *GetTestReply, err error)
	Create(ctx *gin.Context, req *GetTestRequest) (rs *GetTestReply, err error)
	ListUser(ctx *gin.Context, req *GetListTestRequest) (rs *UserList, err error)
	ListUser2(ctx *gin.Context, req *GetList2TestRequest) (rs *UserList2Rely, err error)
	UpdateUser(ctx *gin.Context, req *UpdateUserRequest) (rs *empty.Empty, err error)
}
type TestGinHttpRouterService struct {
	srv    TestGinHttpService
	router *gin.RouterGroup
}

//服务注册abc

func (s *TestGinHttpRouterService) Register() []router.IRouter {
	return router.ApiServiceFn(
		router.ApiGetFn(s.router, OperationGinUrlTestInfo, s.Info),
		router.ApiGetFn(s.router, OperationGinUrlTestList, s.List),
		router.ApiPostFn(s.router, OperationGinUrlTestCreate, s.Create),
		router.ApiGetFn(s.router, OperationGinUrlTestListUser, s.ListUser),
		router.ApiGetFn(s.router, OperationGinUrlTestListUser2, s.ListUser2),
		router.ApiPostFn(s.router, OperationGinUrlTestUpdateUser, s.UpdateUser),
	)
}

// Info 获取用户信息
func (s *TestGinHttpRouterService) Info(ctx *gin.Context) (rs any, err error) {
	var in GetTestRequest
	err = bind.Bind(ctx, &in)
	if err != nil {
		return
	}
	err = validate.Validate(&in)
	if err != nil {
		return
	}
	ctx.Set("operation", OperationGinTestInfo)
	rs, err = s.srv.Info(ctx, &in)
	return
}

// List 获取用户列表
func (s *TestGinHttpRouterService) List(ctx *gin.Context) (rs any, err error) {
	var in GetListTestRequest
	err = bind.Bind(ctx, &in)
	if err != nil {
		return
	}
	err = validate.Validate(&in)
	if err != nil {
		return
	}
	ctx.Set("operation", OperationGinTestList)
	rs, err = s.srv.List(ctx, &in)
	return
}

// Create 创建用户
func (s *TestGinHttpRouterService) Create(ctx *gin.Context) (rs any, err error) {
	var in GetTestRequest
	err = bind.Bind(ctx, &in)
	if err != nil {
		return
	}
	err = validate.Validate(&in)
	if err != nil {
		return
	}
	ctx.Set("operation", OperationGinTestCreate)
	rs, err = s.srv.Create(ctx, &in)
	return
}

func (s *TestGinHttpRouterService) ListUser(ctx *gin.Context) (rs any, err error) {
	var in GetListTestRequest
	err = bind.Bind(ctx, &in)
	if err != nil {
		return
	}
	err = validate.Validate(&in)
	if err != nil {
		return
	}
	ctx.Set("operation", OperationGinTestListUser)
	rs, err = s.srv.ListUser(ctx, &in)
	return
}

func (s *TestGinHttpRouterService) ListUser2(ctx *gin.Context) (rs any, err error) {
	var in GetList2TestRequest
	err = bind.Bind(ctx, &in)
	if err != nil {
		return
	}
	err = validate.Validate(&in)
	if err != nil {
		return
	}
	ctx.Set("operation", OperationGinTestListUser2)
	rs, err = s.srv.ListUser2(ctx, &in)
	return
}

func (s *TestGinHttpRouterService) UpdateUser(ctx *gin.Context) (rs any, err error) {
	var in UpdateUserRequest
	err = bind.Bind(ctx, &in)
	if err != nil {
		return
	}
	err = validate.Validate(&in)
	if err != nil {
		return
	}
	ctx.Set("operation", OperationGinTestUpdateUser)
	rs, err = s.srv.UpdateUser(ctx, &in)
	return
}
func NewTestGinHttpRouterService(router *gin.RouterGroup, srv TestGinHttpService) *TestGinHttpRouterService {
	return &TestGinHttpRouterService{
		srv:    srv,
		router: router,
	}
}
