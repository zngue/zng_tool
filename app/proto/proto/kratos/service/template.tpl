package data
import (
	"context"
    pb "github.com/zngue/zng_tool/api/test/v1"
    "github.com/zngue/zng_tool/demo/biz"
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
	{{if IsAutoReq .RequestType -}}
	var reqData  *biz.{{.RequestType}}
	{{FirstIndex $svrName}}.{{LowerFirst $svrName}}UseCase.{{.Name}}(ctx,reqData)
	{{- else -}}
		err:={{FirstIndex $svrName}}.{{LowerFirst $svrName}}UseCase.{{.Name}}(ctx, {{ ParamsSet .RequestDefault .RequestType }})
	{{- end}}
	return &pb.{{.ReturnType}}{}, nil
}
{{- end}}
