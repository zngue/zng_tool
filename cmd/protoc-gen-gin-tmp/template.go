package main

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"
)

//go:embed template.tpl
var httpTemplate string

func (s *ServiceDesc) execute() string {
	buf := new(bytes.Buffer)
	funcMap := template.FuncMap{
		//判断类型是否包含包名
		"Type": func(useType string) string {
			//if strings.Contains(useType, ".") {
			//	return useType
			//}
			//return fmt.Sprintf("%s.%s", s.GoPackageName, useType)
			return useType
		},
		"EmptyVal": func() string {
			for _, method := range s.Methods {
				//判断是否包含空包 empty.Empty
				if strings.Contains(method.Request, "empty") {
					return "github.com/golang/protobuf/ptypes/empty"
				}
				if strings.Contains(method.Reply, "empty") {
					return "github.com/golang/protobuf/ptypes/empty"
				}
			}
			return ""
		},
	}
	tmpl, err := template.New("service").Funcs(funcMap).Parse(strings.TrimSpace(httpTemplate))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}
