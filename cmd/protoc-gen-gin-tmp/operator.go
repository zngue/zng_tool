package main

import (
	"bytes"
	_ "embed"
	"github.com/zngue/zng_tool/third_party/validate"
	"strings"
	"text/template"
)

type OperatorContent struct {
	Operator  validate.Operator
	FiledType string
	Where     string
	FiledName string
}

//go:embed action/where.tpl
var whereTemplate string

// Execute 执行操作符
func (o *OperatorContent) Execute() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("op").Funcs(map[string]any{
		"IsNumber": func(fileType string) bool {
			if fileType == "number" {
				return true
			}
			return false
		},
		"IsString": func(fileType string) bool {
			if fileType == "string" {
				return true
			}
			return false
		},
		"IsRepeated": func(fileType string) bool {
			if fileType == "repeated" {
				return true
			}
			return false
		},
	}).Parse(strings.TrimSpace(whereTemplate))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, o); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
	return ""
}
