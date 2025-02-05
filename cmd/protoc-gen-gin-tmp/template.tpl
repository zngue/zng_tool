package api
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
	{{- if gt .RequestLent 3 -}}
		var reqData=&biz.{{.Request.GoName}}{}
	{{- end }}
	{{OutParams .ReplyDefault .ReplyMessage false}} s.{{LowerFirst $svrType}}.{{.Name}}(ctx, {{SetReqParams .Request}})
	if err != nil {
       return
    }
    fmt.Println("{{$svrType}}Service->{{.Name}}",{{OutParamsPrintln .ReplyDefault .ReplyMessage}})
    return
}
{{- end -}}