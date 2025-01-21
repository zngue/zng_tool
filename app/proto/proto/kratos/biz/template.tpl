package biz
import "context"
{{$svrName := .ServiceName}}
type {{$svrName}}Repo interface {
	{{- range .Methods }}
	{{.Name}}(ctx context.Context,{{Params .RequestDefault .RequestType}})({{Params .ReturnDefault .ReturnType}},err error)
	{{- end}}
}
{{- range .DoMessage }}
type {{.Name}} struct {
	{{- range .Fields}}
	{{MessageFile .Name .Type .IsRepeated}}
	{{- end}}
}
{{- end}}
type {{$svrName}}UseCase struct {
   {{LowerFirst $svrName}}Repo {{$svrName}}Repo
}
func New{{$svrName}}UseCase({{LowerFirst $svrName}}Repo {{$svrName}}Repo) *{{$svrName}}UseCase {
	return &{{$svrName}}UseCase{
		{{LowerFirst $svrName}}Repo: {{LowerFirst $svrName}}Repo,
	}
}
{{- range .Methods }}
func ({{FirstIndex $svrName}} *{{$svrName}}UseCase) {{.Name}}(ctx context.Context, {{Params .RequestDefault .RequestType}}) ({{Params .ReturnDefault .ReturnType}}, err error) {
	return {{FirstIndex $svrName}}.{{LowerFirst $svrName}}Repo.{{.Name}}(ctx, {{ParamsSet .RequestDefault .RequestType}})
}
{{- end}}