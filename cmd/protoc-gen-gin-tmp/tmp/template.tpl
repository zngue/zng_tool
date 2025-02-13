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
		{{- AutoRequest .RequestMessage }}
	{{- end }}
	{{OutParams .ReplyDefault .ReplyMessage false}} s.{{LowerFirst $svrType}}.{{.Name}}(ctx, {{SetReqParams .Request}})
	if err != nil {
       return
    }
    {{- $content:=ServiceReplyContent .ReplyMessage -}}
    {{- if $content }}
        {{ $content }}
    {{- end }}
    return
}
{{ end }}
{{- $itemVal := IsItem "" }}
{{ if $itemVal.Flag }}
func (s *{{$svrType}}Service) ChangeItem(req *biz.{{$svrType}}Item) *{{$itemVal.MessageType}} {
	return &{{$itemVal.MessageType}}{
		{{ $itemVal.StructSetContent }}
	}
}
{{- end -}}
