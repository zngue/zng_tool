package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/emicklei/proto"
	"strings"
	"text/template"
)

//go:embed template.tpl
var httpTemplate string

func (s *ServiceDesc) execute() string {
	buf := new(bytes.Buffer)
	funcMap := template.FuncMap{
		"Params": func(defaultKey, methodType string) string {
			val, ok := s.MessageMap[methodType]
			if !ok {
				return ""
			}
			if len(val.Elements) <= 3 {
				var params []string
				for _, element := range val.Elements {
					v, oks := element.(*proto.NormalField)
					if !oks {
						continue
					}
					msg := s.MessageMap[v.Type]
					if msg != nil {
						if v.Repeated {
							params = append(params, fmt.Sprintf("%s []*%s", v.Name, v.Type))
						} else {
							params = append(params, fmt.Sprintf("%s *%s", v.Name, v.Type))
						}
					} else {
						if v.Repeated {
							params = append(params, fmt.Sprintf("%s []%s", v.Name, v.Type))
						} else {
							params = append(params, fmt.Sprintf("%s %s", v.Name, v.Type))
						}
					}
				}
				return strings.Join(params, ",")
			} else {
				return fmt.Sprintf("%s *%s", defaultKey, methodType)
			}
		},
	}
	tmpl, err := template.New("biz").Funcs(funcMap).Parse(strings.TrimSpace(httpTemplate))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}
