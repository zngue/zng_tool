{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}
import (
	"github.com/zngue/zng_app/db/api"
	"github.com/zngue/zng_app/pkg/validate"
	"github.com/zngue/zng_app/pkg/bind"
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
		err = bind.Bind(c, &in)
		if err != nil {
			return
		}
		err = validate.Validate(&in)
		if err != nil {
			api.DataApiWithErr(c, err, rs)
			return
		}
		ctx := c.Request.Context()
		ctx = bind.NewServerContext(ctx, c, OperationGin{{$svrType}}{{.OriginalName}})
		rs, err = srv.{{.Name}}(ctx, &in)
		api.DataApiWithErr(c, err, rs)
	}
}
{{- end}}


