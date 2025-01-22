package gen

import (
	"bytes"
	"strings"
	"text/template"
)

func (s *ServiceDesc) execute(httpTemplate string) string {
	buf := new(bytes.Buffer)
	funcMap := template.FuncMap{
		"BindType": func(methodType string) string {
			if methodType == "GET" {
				return "BindJSON"
			} else {
				return "ShouldBind"
			}
		},
		"FnName": func(methodType string) string {
			if methodType == "GET" {
				return "ApiGetFn"
			} else {
				return "ApiPostFn"
			}
		},
	}
	tmpl, err := template.New("http").Funcs(funcMap).Parse(strings.TrimSpace(httpTemplate))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}
