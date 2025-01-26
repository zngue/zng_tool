package service
{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}
import (
	"github.com/gin-gonic/gin"
)
type {{$svrType}}Service struct {
	{{LowerFirst $svrType}} *biz.TestUseCase
}
func New{{$svrType}}Service({{LowerFirst $svrType}} *biz.{{$svrType}}UseCase) {{.ServiceTypeName}}GinHttpService {
	return &{{$svrType}}Service{
		{{LowerFirst $svrType}}: {{LowerFirst $svrType}},
	}
}
{{- range .Methods }}
func (s *{{$svrType}}Service){{.Name}}(ctx *gin.Context, req *{{NameTo .Request}}) (rs *{{NameTo .Reply}}, err error){

	//判断RequestLent 大于 3
	{{ if gt .RequestLent 3 }}
		var reqData {{NameTo .Request}}
	{{ end }}
	{{OutParams .ReplyDefault .ReplyMessage}}:=s.{{LowerFirst $svrType}}.{{.Name}}(ctx, {{SetReqParams .Request}})
    return
}
{{- end}}