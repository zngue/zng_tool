package biz
{{$svrName := .ServiceName}}
type {{$svrName}}Repo interface {

}
type {{$svrName}}UseCase struct {
   {{$svrName}}Repo {{$svrName}}Repo
}

{{- range .Methods }}
func (uc *{{$svrName}}UseCase) {{.Name}}(ctx context.Context, req any) (rs any, err error) {
	return
}
{{- end}}