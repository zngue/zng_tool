package service
{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}
import (
	"github.com/gin-gonic/gin"
	{{.GoImportPath}}
)
type {{$svrType}}Service struct {
}
func New{{$svrType}}Service() {{.GoPackageName}}.{{$svrType}}GinHttpService {
	return &{{$svrType}}Service{}
}
{{- range .Methods }}
 func (s *{{$svrType}}Service){{.Name}}(ctx *gin.Context, req *{{.Request}}) (rs *{{.Reply}}, err error){

    return
 }
{{- end}}