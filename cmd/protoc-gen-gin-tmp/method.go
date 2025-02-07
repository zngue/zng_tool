package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/zngue/zng_tool/app/util"
	"github.com/zngue/zng_tool/third_party/google/annotations"
	"github.com/zngue/zng_tool/third_party/validate"
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
	"text/template"
)

func (s *MethodDesc) execute() (tmp string) {
	switch s.Action {
	case annotations.Action_add:
		tmp = s.Add()
	case annotations.Action_delete:
	case annotations.Action_list:
	case annotations.Action_list_page:
	case annotations.Action_update:
		tmp = s.Update()
	case annotations.Action_query:
	}
	return
}

//go:embed action/add.tpl
var addTemplate string

func (s *MethodDesc) MapFn() template.FuncMap {
	return template.FuncMap{
		"StructName": func(req *protogen.Message) string {
			return string(req.Desc.FullName())
		},
		"SetFiled": func(message *protogen.Message) string {
			var params []string
			if len(message.Fields) > 3 {
				for _, field := range message.Fields {
					kind := field.Desc.Kind().String()
					if kind == "message" || kind == "enum" {
						continue
					}
					params = append(params, fmt.Sprintf("%s: req.%s,", field.GoName, field.GoName))
				}
			} else {
				for _, field := range message.Fields {
					kind := field.Desc.Kind().String()
					if kind == "message" || kind == "enum" {
						continue
					}
					params = append(params, fmt.Sprintf("%s: req.%s,", util.LowerFirst(field.GoName), field.GoName))
				}
			}
			return strings.Join(params, "\n\t\t")
		},
		"UpdateWhereOperatorMore": func(message *protogen.Message) string {
			var params []string
			if len(message.Fields) > 3 {
				for _, field := range message.Fields {
					operator, _ := DoFieldOperator(field)
					if operator != validate.Operator_OPERATOR_UNKNOWN {
						name := util.CamelToSnake(field.GoName)
						content := Operator(name, fmt.Sprintf("req.%s", field.GoName), operator)
						if content != "" {
							params = append(params, content)
						}
					}
				}
			} else {
				for _, field := range message.Fields {
					operator, _ := DoFieldOperator(field)
					if operator != validate.Operator_OPERATOR_UNKNOWN {
						name := util.CamelToSnake(field.GoName)
						content := Operator(name, util.LowerFirst(field.GoName), operator)
						if content != "" {
							params = append(params, content)
						}
					}
				}
			}
			return strings.Join(params, ",\n\t\t")
		},
		"UpdateWhereOperator": func(req *protogen.Field) string {
			operator, filedType := DoFieldOperator(req)
			if operator != 0 && filedType != "" {
				switch filedType {
				case "number":
					//return Operator(req.GoName, operator)
				case "string":
					//return Operator(req.GoName, operator)
				}
			}
			return ""
		},
		"UpdateOperatorMore": func(message *protogen.Message) string {
			var params []string
			if len(message.Fields) > 3 {
				for _, field := range message.Fields {
					operator, _ := DoFieldOperator(field)
					if operator == validate.Operator_OPERATOR_UNKNOWN {
						var name = util.CamelToSnake(field.GoName)
						params = append(params, fmt.Sprintf(" \"%s\" : req.%s", name, field.GoName))
					}
				}
			} else {
				for _, field := range message.Fields {
					operator, _ := DoFieldOperator(field)
					if operator == validate.Operator_OPERATOR_UNKNOWN {
						var name = util.CamelToSnake(field.GoName)
						params = append(params, fmt.Sprintf(" \"%s\" : %s", name, field.GoName))
					}
				}
			}
			return strings.Join(params, ",\n\t\t")
		},
		"UpdateOperator": func(req *protogen.Field) string {
			operator, _ := DoFieldOperator(req)
			if operator == validate.Operator_OPERATOR_UNKNOWN {
				name := util.CamelToSnake(req.GoName)
				return fmt.Sprintf("\"%s\": req.%s,", name, req.GoName)
			}
			return ""
		},
	}
}

func Operator(name, goName string, operator validate.Operator) (op string) {
	switch operator {
	case validate.Operator_eq:
		var key = fmt.Sprintf("%s = ?", name)
		op = fmt.Sprintf("\"%s\" : %s,", key, goName)
		return
	case validate.Operator_neq:
		var key = fmt.Sprintf("%s != ?", name)
		op = fmt.Sprintf("\"%s\" : %s,", key, goName)
		return
	case validate.Operator_in:
		var key = fmt.Sprintf("%s in ?", name)
		op = fmt.Sprintf("\"%s\" : %s,", key, goName)
		return
	case validate.Operator_not_in:
		var key = fmt.Sprintf("%s not in ?", name)
		op = fmt.Sprintf("%s : %s", key, goName)
		return
	case validate.Operator_gt:
		var key = fmt.Sprintf("%s > ?", name)
		op = fmt.Sprintf("%s : %s", key, goName)
		return
	case validate.Operator_gte:
		var key = fmt.Sprintf("%s >= ?", name)
		op = fmt.Sprintf("%s : %s", key, goName)
		return
	case validate.Operator_lt:
		var key = fmt.Sprintf("%s < ?", name)
		op = fmt.Sprintf("%s : %s", key, goName)
		return
	case validate.Operator_lte:
		var key = fmt.Sprintf("%s <= ?", name)
		op = fmt.Sprintf("%s : %s", key, goName)
		return
	case validate.Operator_like:
		var key = fmt.Sprintf("%s like ?", name)
		op = fmt.Sprintf("%s :   \"%%\" + %s+ \"%%\" ", key, goName)
		return
	case validate.Operator_not_like:
		var key = fmt.Sprintf("%s not like ?", name)
		op = fmt.Sprintf("%s :   \"%%\" + %s+ \"%%\" ", key, goName)
		return
	}
	return
}

func (s *MethodDesc) Add() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("add").Funcs(s.MapFn()).Parse(strings.TrimSpace(addTemplate))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}

//go:embed action/update.tpl
var updateTemplate string

func (s *MethodDesc) Update() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("update").Funcs(s.MapFn()).Parse(strings.TrimSpace(updateTemplate))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}
