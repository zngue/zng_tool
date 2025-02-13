package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/zngue/zng_tool/app/util"
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
	"text/template"
)

type DbReplace struct {
	MessageMap      map[string]*protogen.Message
	ServerType      string
	LowerServerType string
	Message         *protogen.Message
	Pkg             string
}

//go:embed tmp/db_template.tpl
var dbTemplate string

// execute
func (d *DbReplace) execute() string {
	d.Message = d.getMessage()
	buf := new(bytes.Buffer)
	tmpl, err := template.New("db").Funcs(map[string]any{
		"StructType": func(req *protogen.Message) string {
			var params []string
			if req != nil && len(req.Fields) > 0 {
				for _, field := range req.Fields {
					val := util.StructType(field)
					tmp := fmt.Sprintf("%s `gorm:\"column:%s\"`", val, util.CamelToSnake(field.GoName))
					params = append(params, tmp)
				}
			}
			return strings.Join(params, "\n\t")
		},
	}).Parse(strings.TrimSpace(dbTemplate))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, d); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}

// 获取message
func (d *DbReplace) getMessage() (message *protogen.Message) {
	var ok bool
	dbKey := fmt.Sprintf("%sDB", d.ServerType)
	if message, ok = d.MessageMap[dbKey]; ok && message != nil {
		return message
	}
	itemKey := fmt.Sprintf("%sItem", d.ServerType)
	if message, ok = d.MessageMap[itemKey]; ok && message != nil {
		return message
	}
	return
}
