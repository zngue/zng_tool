package biz

import (
	"context"
)

{{- $svrType := .ServiceType -}}
{{- $svrName := .ServiceName -}}
{{- $lowerServiceType := .LowerServiceType -}}
{{$lowerIndex := .LowerIndex}}
type {{$svrType}}Repo interface {
	{{- range .Methods }}
	{{.Name}}(ctx context.Context, {{InParamsType .RequestDefault  .RequestMessage}}) ({{OutParamsType .ReplyDefault  .ReplyMessage}})
	{{- end}}
}

{{- range .UseMessage }}
type {{.GoIdent.GoName}} struct {
{{- range .Fields }}
	{{StructType .}}
{{- end}}
}
{{- end}}
type {{$svrType}}UseCase struct {
	{{$lowerServiceType}} {{$svrType}}Repo
}
func New{{$svrType}}UseCase({{$lowerServiceType}} {{$svrType}}Repo) *{{$svrType}}UseCase {
	return &{{$svrType}}UseCase{
		{{$lowerServiceType}}: {{$lowerServiceType}},
	}
}
{{- range .Methods }}
func ({{$lowerIndex}} *{{$svrType}}UseCase) {{.Name}}(ctx context.Context, {{InParamsType .RequestDefault  .RequestMessage}}) ({{OutParamsType .ReplyDefault  .ReplyMessage}}) {
	{{OutParams .ReplyDefault .ReplyMessage true}} {{$lowerIndex}}.{{$lowerServiceType}}.{{.Name}}(ctx, {{InParamsSet .RequestDefault .RequestMessage}})
	return
}
{{- end}}