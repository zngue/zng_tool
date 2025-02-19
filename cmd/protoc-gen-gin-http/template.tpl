{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}
import (
	"github.com/zngue/zng_app/db/api"
	"context"
	"github.com/zngue/zng_app/pkg/validate"
	"github.com/zngue/zng_app/pkg/bind"
	"github.com/gin-gonic/gin"
)
// 注册服务
func Register{{$svrType}}GinServer(router *gin.RouterGroup ,srv {{$svrType}}GinHttpService) {
	 New{{$svrType}}GinHttpRouterService(router, srv).Register()
}
// 服务操作
{{- range .Methods}}
const OperationGin{{$svrType}}{{.OriginalName}} = "{{$svrName}}.{{.OriginalName}}"
{{- end}}
// 服务url
{{- range .Methods}}
const OperationGinUrl{{$svrType}}{{.OriginalName}}="{{.Path}}"
{{- end}}
//服务接口 {{- .Comment }}
type {{$svrType}}GinHttpService interface {
	{{- range .Methods }}
	{{.Name}}(ctx context.Context, req *{{.Request}}) (rs *{{.Reply}}, err error)
	{{- end}}
}
type {{$svrType}}GinHttpRouterService struct {
	srv    {{$svrType}}GinHttpService
	router *gin.RouterGroup
}
//服务注册 {{- .Comment }}
func (s *{{$svrType}}GinHttpRouterService) Register(){
	{{- range .Methods }}
	s.router.{{FnName .Method}}(OperationGinUrl{{$svrType}}{{.OriginalName}}, _{{$svrType}}_{{.Name}}{{.ServerIndex}}_GIN_HTTP_Handler(s.srv))
	{{- end}}
}
func New{{$svrType}}GinHttpRouterService(router *gin.RouterGroup,srv {{$svrType}}GinHttpService)  *{{$svrType}}GinHttpRouterService {
	return  &{{$svrType}}GinHttpRouterService{
		srv:   srv,
		router: router,
	}
}

{{- range .Methods }}
{{.Comment}}
func _{{$svrType}}_{{.Name}}{{.ServerIndex}}_GIN_HTTP_Handler(srv {{$svrType}}GinHttpService) gin.HandlerFunc  {
	return func(c *gin.Context) {
		var (
			in {{.Request}}
			err error
			rs  *{{.Reply}}
		)
		err = bind.Bind(c, &in)
		if err != nil {
			return
		}
		err = validate.Validate(&in)
		if err != nil {
			api.DataApiWithErr(c, err, rs)
			return
		}
		c.Set("operation", OperationGin{{$svrType}}{{.OriginalName}})
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, "operation", OperationGin{{$svrType}}{{.OriginalName}})
		ctx = context.WithValue(ctx, "gin_ctx", c)
		ctx, err = bind.GetMiddleWires(ctx)
		if err != nil {
			api.DataApiWithErr(c, err, rs)
			return
		}
		rs, err = srv.{{.Name}}(ctx, &in)
		api.DataApiWithErr(c, err, rs)
	}
}
{{- end}}


