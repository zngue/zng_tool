{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}
import (
	"context"
	"github.com/zngue/zng_app/pkg/validate"
	"github.com/zngue/zng_app/pkg/bind"
	"github.com/zngue/zng_app/pkg/errors_ez"
	"github.com/gin-gonic/gin"
)

// 服务操作
{{- range .Methods}}
const OperationGin{{$svrType}}{{.OriginalName}} = "{{$svrName}}.{{.OriginalName}}"
{{- end}}
// 服务url
{{- range .Methods}}
const OperationGinUrl{{$svrType}}{{.OriginalName}}="{{.Path}}"
{{- end}}

type {{$svrType}}GinClient struct {
	Unimplemented{{$svrType}}Server
	srv bind.ClientServer
}

func New{{$svrType}}GinClient(srv bind.ClientServer) {{$svrType}}Server {
	return &{{$svrType}}GinClient{
		srv: srv,
	}
}
{{- range .Methods}}
{{.Comment}}
func (c *{{$svrType}}GinClient){{.Name}}(ctx context.Context, req *{{.Request}}) (rs *{{.Reply}}, err error) {
	err = c.srv.{{FnName .Method}}(ctx, OperationGinUrl{{$svrType}}{{.OriginalName}}, req, &rs)
	return
}
{{- end}}


//服务注册 {{- .Comment }}
func Register{{$svrType}}GinRouter(router *gin.Engine, srv {{$svrType}}Server){
	{{- range .Methods }}
	router.{{FnName .Method}}(OperationGinUrl{{$svrType}}{{.OriginalName}}, _{{$svrType}}_{{.Name}}{{.ServerIndex}}_GIN_HTTP_Handler(srv))
	{{- end}}
}
{{- range .Methods }}
{{.Comment}}
func _{{$svrType}}_{{.Name}}{{.ServerIndex}}_GIN_HTTP_Handler(srv {{$svrType}}Server) gin.HandlerFunc  {
	return func(c *gin.Context) {
		var (
			in {{.Request}}
			err error
			rs  *{{.Reply}}
		)
		ctx := c.Request.Context()
		ctx = bind.NewServerContext(ctx, c, OperationGin{{$svrType}}{{.OriginalName}})
		err = bind.Bind(c, &in)
		if err != nil {
			err = errors_ez.Wrap(err, "绑定参数失败")
			bind.ApiErrorParameter(c, err, bind.DataMsg("绑定参数失败"))
			return
		}
		err = validate.Validate(&in)
		if err != nil {
			err = errors_ez.Wrap(err, "参数验证失败")
			bind.ApiErrorParameter(c, err, bind.DataCode(bind.ErrorParameter), bind.DataMsg("参数验证失败"))
			return
		}
		rs, err = bind.MiddlewareHandle[*{{.Reply}}](ctx, func(ctx context.Context) (*{{.Reply}}, error) {
			return srv.{{.Name}}(ctx, &in)
		})
		bind.ApiDataWithErr(c, err, rs)
	}
}
{{- end}}

