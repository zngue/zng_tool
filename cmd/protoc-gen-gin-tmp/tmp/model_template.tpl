package model

import (
	"context"
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
	{{if ModelContent .  $svrType }}
	{{ModelContent .  $svrType }}
	{{else}}
	fmt.Println(dbConn)
	//TODO implement me
	panic("{{$svrType}}Repo->{{.Name}} implement me")
	{{end}}
	return
}
{{- end}}
{{- $itemVal := IsItem "\n\t\t" }}
{{ if $itemVal.Flag }}
func ({{$lowerIndex}} *{{$svrType}}Repo) ChangeItem(req *db.{{$svrType}}) *biz.{{$svrType}}Item {
	return &biz.{{$svrType}}Item{
		{{$itemVal.StructSetContent}}
	}
}
{{- end -}}
