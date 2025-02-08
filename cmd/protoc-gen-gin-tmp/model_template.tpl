package model

import (
	"context"
	"fmt"
	"github.com/zngue/zng_app/db/data"
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
// {{ .Name }} 请求方法 {{ .Action }}
func ({{$lowerIndex}} *{{$svrType}}Repo) {{.Name}}(ctx context.Context, {{InParamsTypeModel .RequestDefault  .RequestMessage}}) ({{OutParamsTypeModel .ReplyDefault  .ReplyMessage}})  {
	var conn = {{$lowerIndex}}.conn.WithContext(ctx)
	var dbConn=data.NewDB[db.{{$svrType}}](conn)
	{{ModelContent .  $svrType }}
	fmt.Println(dbConn)
	//TODO implement me
	panic("{{$svrType}}Repo->{{.Name}} implement me")
	return
}
{{- end}}
