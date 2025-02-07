package api
{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}
import (
	"context"
)
type {{$svrType}}Service struct {
	{{LowerFirst $svrType}} *biz.{{$svrType}}UseCase
}
func New{{$svrType}}Service({{LowerFirst $svrType}} *biz.{{$svrType}}UseCase) {{.ServiceTypeName}}GinHttpService {
	return &{{$svrType}}Service{
		{{LowerFirst $svrType}}: {{LowerFirst $svrType}},
	}
}
{{- range .Methods }}
func (s *{{$svrType}}Service){{.Name}}(ctx context.Context, req *{{NameTo .Request}}) (rs *{{NameTo .Reply}}, err error){
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