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
		tmp = s.Del()
	case annotations.Action_list:
		tmp = s.List()
	case annotations.Action_list_page:
		tmp = s.ListPage()
	case annotations.Action_update:
		tmp = s.Update()
	case annotations.Action_query:
		tmp = s.Query()
	}
	return
}

//go:embed action/add.tpl
var addTemplate string

const sigTemplate = `

`

func (s *MethodDesc) MapFn() template.FuncMap {
	return template.FuncMap{
		"StructName": func(req *protogen.Message) string {
			return string(req.Desc.FullName())
		},
		"LowerIndex": func() string {
			return util.LowerIndex(s.SvrType)
		},
		"requestName": func(req *protogen.Message) string {
			for _, field := range req.Fields {
				msgType := util.MsgType(field)
				key := fmt.Sprintf("%sItem", s.SvrType)
				if msgType == util.AutoNormal || msgType == util.AutoRepeated {
					kind := field.Message.GoIdent.GoName
					if kind == key {
						return util.LowerFirst(field.GoName)
					}
				}
			}
			return ""
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

		"ListWhereOperator": func(message *protogen.Message) string {
			var params []string
			if len(message.Fields) > 3 {
				for _, field := range message.Fields {
					operator, fileType := DoFieldOperator(field)
					if operator != validate.Operator_OPERATOR_UNKNOWN {
						name := util.CamelToSnake(field.GoName)
						msgType := util.MsgType(field)
						where := ListOperator(name, fmt.Sprintf("req.%s", field.GoName), operator)
						if msgType == util.SystemRepeated {
							fileType = "repeated"
						}
						if where != "" {
							op := &OperatorContent{
								Operator:  operator,
								FiledType: fileType,
								Where:     where,
								FiledName: field.GoName,
							}
							tmp := op.Execute()
							if tmp != "" {
								params = append(params, tmp)
							}
						}
					}
				}
			} else {
				for _, field := range message.Fields {
					operator, fileType := DoFieldOperator(field)
					if operator != validate.Operator_OPERATOR_UNKNOWN {
						name := util.CamelToSnake(field.GoName)
						where := ListOperator(name, util.LowerFirst(field.GoName), operator)
						if where != "" {
							op := &OperatorContent{
								Operator:  operator,
								FiledType: fileType,
								Where:     where,
							}
							tmp := op.Execute()
							if tmp != "" {
								params = append(params, tmp)
							}
						}
					}
				}
			}
			return strings.Join(params, "\n\t")
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
		"UpdateOperatorMore": func(message *protogen.Message) string {
			var params []string
			if len(message.Fields) > 3 {
				for _, field := range message.Fields {
					operator, _ := DoFieldOperator(field)
					if operator == validate.Operator_OPERATOR_UNKNOWN {
						var name = util.CamelToSnake(field.GoName)
						params = append(params, fmt.Sprintf(" \"%s\" : req.%s,", name, field.GoName))
					}
				}
			} else {
				for _, field := range message.Fields {
					operator, _ := DoFieldOperator(field)
					if operator == validate.Operator_OPERATOR_UNKNOWN {
						var name = util.CamelToSnake(field.GoName)
						params = append(params, fmt.Sprintf(" \"%s\" : %s,", name, util.LowerFirst(field.GoName)))

					}
				}
			}
			return strings.Join(params, "\n\t\t")
		},
	}
}
func ListOperator(name, goName string, operator validate.Operator) (op string) {
	switch operator {
	case validate.Operator_eq:
		var key = fmt.Sprintf("%s = ?", name)
		op = fmt.Sprintf("where[\"%s\"] = %s", key, goName)
		return
	case validate.Operator_neq:
		var key = fmt.Sprintf("%s != ?", name)
		op = fmt.Sprintf("where[\"%s\"] = %s", key, goName)
		return
	case validate.Operator_in:
		var key = fmt.Sprintf("%s in ?", name)
		op = fmt.Sprintf("where[\"%s\"] = %s", key, goName)
		return
	case validate.Operator_not_in:
		var key = fmt.Sprintf("%s not in ?", name)
		op = fmt.Sprintf("where[\"%s\"] = %s", key, goName)
		return
	case validate.Operator_gt:
		var key = fmt.Sprintf("%s > ?", name)
		op = fmt.Sprintf("where[\"%s\"] = %s", key, goName)
		return
	case validate.Operator_gte:
		var key = fmt.Sprintf("%s >= ?", name)
		op = fmt.Sprintf("where[\"%s\"] = %s", key, goName)
		return
	case validate.Operator_lt:
		var key = fmt.Sprintf("%s < ?", name)
		op = fmt.Sprintf("where[\"%s\"] = %s", key, goName)
		return
	case validate.Operator_lte:
		var key = fmt.Sprintf("%s <= ?", name)
		op = fmt.Sprintf("where[\"%s\"] = %s", key, goName)
		return
	case validate.Operator_like:
		var key = fmt.Sprintf("%s like ?", name)
		op = fmt.Sprintf("where[\"%s\"] = \"%%\" + %s + \"%%\" ", key, goName)
		return
	case validate.Operator_not_like:
		var key = fmt.Sprintf("%s not like ?", name)
		op = fmt.Sprintf("where[\"%s\"] = \"%%\" + %s + \"%%\" ", key, goName)
		return
	}
	return
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

//go:embed action/del.tpl
var delTemplate string

func (s *MethodDesc) Del() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("del").Funcs(s.MapFn()).Parse(strings.TrimSpace(delTemplate))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}

//go:embed action/query.tpl
var queryTemplate string

func (s *MethodDesc) Query() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("query").Funcs(s.MapFn()).Parse(strings.TrimSpace(queryTemplate))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}

//go:embed action/list_page.tpl
var listPageTemplate string

func (s *MethodDesc) ListPage() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("list_page").Funcs(s.MapFn()).Parse(strings.TrimSpace(listPageTemplate))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}

//go:embed action/list.tpl
var listTemplate string

func (s *MethodDesc) List() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("list").Funcs(s.MapFn()).Parse(strings.TrimSpace(listTemplate))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return buf.String()
}
