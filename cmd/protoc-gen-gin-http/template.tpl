{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/zng_app/pkg/router"
	"github.com/zngue/zng_app/pkg/validate"
)
{{- range .Methods}}
const OperationGin{{$svrType}}{{.OriginalName}} = "{{$svrName}}.{{.OriginalName}}"
{{- end}}
type {{$svrType}}GinHttpRouterService struct {
	srv    {{$svrType}}GinHttpService
	router *gin.RouterGroup
}
//{{.Comment}} 服务注册
func (s *{{$svrType}}GinHttpRouterService) Register() []router.IRouter {
	return router.ApiServiceFn(
	{{- range .Methods }}
		router.{{FnName .Method}}(s.router, "{{.Path}}", _{{$svrType}}_{{.Name}}{{.ServerIndex}}_GIN_HTTP_Handler(s.srv)),
	{{- end}}
	)
}
func New{{$svrType}}GinHttpRouterService(router *gin.RouterGroup,srv {{$svrType}}GinHttpService)  *{{$svrType}}GinHttpRouterService {
	return  &{{$svrType}}GinHttpRouterService{
		srv:   srv,
		router: router,
	}
}

//{{.Comment}} 服务接口
type {{$svrType}}GinHttpService interface {
	{{- range .Methods }}
	{{.Name}}(ctx *gin.Context, req *{{.Request}}) (rs *{{.Reply}}, err error)
	{{- end}}
}
{{- range .Methods }}
func _{{$svrType}}_{{.Name}}{{.ServerIndex}}_GIN_HTTP_Handler(srv {{$svrType}}GinHttpService) router.Fn {
	return func(ctx *gin.Context) (rs any, err error) {
		var in *{{.Request}}
		if err = ctx.{{ BindType .Method }}(&in); err != nil {
            return
        }
		err = validate.Validate(in)
		if err != nil {
			return
		}
		ctx.Set("operation", OperationGin{{$svrType}}{{.OriginalName}})
		rs, err = srv.{{.Name}}(ctx, in)
		return
	}
}
{{- end}}

