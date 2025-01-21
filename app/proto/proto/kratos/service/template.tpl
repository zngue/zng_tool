package data
import (
	"context"
)
{{$svrName := .ServiceName}}
type {{$svrName}}Service struct {
	pb.Unimplemented{{$svrName}}Server
	{{LowerFirst $svrName}}UseCase *biz.{{$svrName}}UseCase
}
func New{{$svrName}}Service({{LowerFirst $svrName}}UseCase *biz.{{$svrName}}UseCase) *{{$svrName}}Service {
	return &{{$svrName}}Service{
		{{LowerFirst $svrName}}UseCase: {{LowerFirst $svrName}}UseCase,
	}
}
{{- range .Methods }}
func ({{FirstIndex $svrName}} *{{$svrName}}Service) {{.Name}}(ctx context.Context,req *pb.{{.RequestType}}) (*pb.{{.ReturnType}},  error) {
	//TODO implement me
	{{if IsAutoReq .RequestType}}
	var reqData  *biz.{{.RequestType}}
	{{else}}
		{{FirstIndex $svrName}}.{{LowerFirst $svrName}}UseCase.{{.Name}}(ctx, {{ ParamsSet .RequestDefault .RequestType }})
	{{end}}
	return &pb.{{.ReturnType}}{}, nil
}
{{- end}}
