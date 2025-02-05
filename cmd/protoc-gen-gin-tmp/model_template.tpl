package model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

{{- $svrType := .ServiceType -}}
{{- $svrName := .ServiceName -}}
{{- $lowerServiceType := .LowerServiceType -}}
{{$lowerIndex := .LowerIndex}}
func New{{$svrType}}Repo(conn *gorm.DB) biz.{{$svrType}}Repo {
	return &{{$svrType}}Repo{
		conn: conn,
	}
}
type {{$svrType}}Repo struct {
	conn *gorm.DB
}

{{- range .Methods }}
func ({{$lowerIndex}} *{{$svrType}}Repo) {{.Name}}(ctx *gin.Context, {{InParamsTypeModel .RequestDefault  .RequestMessage}}) ({{OutParamsTypeModel .ReplyDefault  .ReplyMessage}})  {
	//TODO implement me
	panic("{{$svrType}}Repo->{{.Name}} implement me")
}
{{- end}}
