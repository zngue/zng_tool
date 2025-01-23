{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/zng_app/pkg/router"
	"github.com/zngue/zng_app/pkg/validate"
	"github.com/zngue/zng_app/pkg/bind"
)
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
	{{.Name}}(ctx *gin.Context, req *{{.Request}}) (rs *{{.Reply}}, err error)
	{{- end}}
}
type {{$svrType}}GinHttpRouterService struct {
	srv    {{$svrType}}GinHttpService
	router *gin.RouterGroup
}
//服务注册 {{- .Comment }}
func (s *{{$svrType}}GinHttpRouterService) Register() []router.IRouter {
	return router.ApiServiceFn(
	{{- range .Methods }}
		router.{{FnName .Method}}(s.router, OperationGinUrl{{$svrType}}{{.OriginalName}}, s.{{.Name}}),
	{{- end}}
	)
}
{{- range .Methods }}
{{.Comment}}
func (s *{{$svrType}}GinHttpRouterService) {{.Name}}(ctx *gin.Context)  (rs any, err error)  {
	var in {{.Request}}
    err = bind.Bind(ctx,&in)
    if err != nil {
        return
    }
    err = validate.Validate(&in)
    if err != nil {
        return
    }
    ctx.Set("operation", OperationGin{{$svrType}}{{.OriginalName}})
    rs, err = s.srv.{{.Name}}(ctx, &in)
    return
}
{{- end}}
func New{{$svrType}}GinHttpRouterService(router *gin.RouterGroup,srv {{$svrType}}GinHttpService)  *{{$svrType}}GinHttpRouterService {
	return  &{{$svrType}}GinHttpRouterService{
		srv:   srv,
		router: router,
	}
}



