package data
import "context"
{{$svrName := .ServiceName}}
{{- range .DoMessage }}
{{- end}}
type {{$svrName}}Repo struct {
    data  *Data
    rpc   *GRPCClient
}
func New{{$svrName}}Repo(data *Data,redis *Redis,rpc *GRPCClient) biz.{{$svrName}}Repo {
	return &{{$svrName}}Repo{
	   data:data,
	   rpc:rpc,
	}
}
{{- range .Methods }}
func ({{FirstIndex $svrName}} *{{$svrName}}Repo) {{.Name}}(ctx context.Context, {{Params .RequestDefault .RequestType}}) ({{Params .ReturnDefault .ReturnType}}, err error) {
	return
}
{{- end}}