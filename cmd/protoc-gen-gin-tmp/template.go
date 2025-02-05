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

//go:embed template.tpl
var httpTemplate string

//go:embed biz_template.tpl
var bizTemplate string

//go:embed model_template.tpl
var modelTemplate string

type ServiceDesc struct {
	LowerServiceType string
	LowerIndex       string
	ServiceType      string
	ServiceName      string
	ServiceTypeName  string
	Metadata         string
	Methods          []*MethodDesc
	GoPackageName    string
	GoImportPath     string
	MessageMap       map[string]*protogen.Message
	GeneratedFile    *protogen.GeneratedFile
	LowerServiceName string
	MessageLessThree []string
	UseMessage       []*protogen.Message
}
type MethodDesc struct {
	Name           string
	OriginalName   string
	MethodIndex    int
	ServerIndex    int
	Request        protogen.GoIdent
	RequestMessage *protogen.Message
	RequestDefault string
	Reply          protogen.GoIdent
	ReplyDefault   string
	ReplyMessage   *protogen.Message
	GoPackageName  string
	ReplyLent      int
	RequestLent    int
}

func (s *ServiceDesc) ParamsTypeDel(def string, req *protogen.Message, pkg string) (params []string) {
	if len(req.Fields) > 3 {
		if pkg != "" {
			params = append(params, fmt.Sprintf("%s *%s.%s", def, pkg, req.GoIdent.GoName))
		} else {
			params = append(params, fmt.Sprintf("%s *%s", def, req.GoIdent.GoName))
		}

	} else {
		for _, element := range req.Fields {
			msgType := util.MsgType(element)
			var kind = element.Desc.Kind().String()
			if kind == "message" {
				kind = element.GoName
			}
			var val string
			if pkg != "" {
				switch msgType {
				case util.AutoRepeated:
					val = fmt.Sprintf("%s []*%s.%s", util.LowerFirst(element.GoName), pkg, kind)
				case util.AutoNormal:
					val = fmt.Sprintf("%s *%s.%s", util.LowerFirst(element.GoName), pkg, kind)
				case util.SystemRepeated:
					val = fmt.Sprintf("%s []%s", util.LowerFirst(element.GoName), kind)
				case util.SystemNormal:
					val = fmt.Sprintf("%s %s", util.LowerFirst(element.GoName), kind)
				default:
					val = fmt.Sprintf("%s %s", util.LowerFirst(element.GoName), kind)
				}
			} else {
				switch msgType {
				case util.AutoRepeated:
					val = fmt.Sprintf("%s []*%s", util.LowerFirst(element.GoName), kind)
				case util.AutoNormal:
					val = fmt.Sprintf("%s *%s", util.LowerFirst(element.GoName), kind)
				case util.SystemRepeated:
					val = fmt.Sprintf("%s []%s", util.LowerFirst(element.GoName), kind)
				case util.SystemNormal:
					val = fmt.Sprintf("%s %s", util.LowerFirst(element.GoName), kind)
				default:
					val = fmt.Sprintf("%s %s", util.LowerFirst(element.GoName), kind)
				}
			}

			params = append(params, val)
		}
	}
	return
}
func (s *ServiceDesc) MapFn() template.FuncMap {
	return template.FuncMap{
		"StructName": func(req *protogen.Message) string {
			return string(req.Desc.FullName())
		},
		"StructType": func(req *protogen.Field) string {
			msgType := util.MsgType(req)
			var kind = req.Desc.Kind().String()
			if kind == "message" {
				kind = req.GoName
			}

			if req.Extendee != nil {
				kind = fmt.Sprintf("%s_%s", req.Extendee.GoIdent.GoName, kind)
			}

			var val string
			switch msgType {
			case util.AutoRepeated:
				val = fmt.Sprintf("%s []*%s", req.GoName, kind)
			case util.AutoNormal:
				val = fmt.Sprintf("%s *%s", req.GoName, kind)
			case util.SystemRepeated:
				val = fmt.Sprintf("%s []%s", req.GoName, kind)
			case util.SystemNormal:
				val = fmt.Sprintf("%s %s", req.GoName, kind)
			default:
				val = fmt.Sprintf("%s %s", req.GoName, kind)
			}
			return val
		},
		"NameTo": func(useType protogen.GoIdent) string {
			return s.GeneratedFile.QualifiedGoIdent(useType)
		},
		"LowerFirst": util.LowerFirst,
		"SetReqParams": func(req protogen.GoIdent) string {
			var params []string
			var useType = req.GoName
			if val, ok := s.MessageMap[useType]; ok {
				if len(val.Fields) > 3 {
					params = append(params, "reqData")
				} else {
					for _, element := range val.Fields {
						params = append(params, fmt.Sprintf("req.%s", element.GoName))
					}
				}
			} else {
				params = append(params, "reqData")
			}
			return strings.Join(params, ",")
		},
		"OutParamsType": func(def string, req *protogen.Message) string {
			var params = s.ParamsTypeDel(def, req, "")
			params = append(params, "err error")
			return strings.Join(params, ",")
		},
		"OutParamsTypeModel": func(def string, req *protogen.Message) string {
			var params = s.ParamsTypeDel(def, req, "biz")
			params = append(params, "err error")
			return strings.Join(params, ",")
		},
		"InParamsType": func(def string, req *protogen.Message) string {
			var params = s.ParamsTypeDel(def, req, "")
			return strings.Join(params, ",")
		},
		"InParamsTypeModel": func(def string, req *protogen.Message) string {
			var params = s.ParamsTypeDel(def, req, "biz")
			return strings.Join(params, ",")
		},
		"InParamsSet": func(def string, req *protogen.Message) string {
			var params []string
			if len(req.Fields) > 3 {
				params = append(params, def)
			} else {
				for _, element := range req.Fields {
					params = append(params, util.LowerFirst(element.GoName))
				}
			}
			return strings.Join(params, ",")
		},
		"OutParamsPrintln": func(def string, req *protogen.Message) string {
			var params []string
			if len(req.Fields) > 3 {
				params = append(params, def)
			} else {
				for _, element := range req.Fields {
					params = append(params, util.LowerFirst(element.GoName))
				}
			}
			params = append(params, "err")
			paramsContent := strings.Join(params, ", ")
			return paramsContent
		},
		"OutParams": func(def string, req *protogen.Message, is bool) string {
			var params []string
			if len(req.Fields) > 3 {
				params = append(params, def)
			} else {
				for _, element := range req.Fields {
					params = append(params, util.LowerFirst(element.GoName))
				}
			}
			params = append(params, "err")
			paramsContent := strings.Join(params, ", ")
			if is {
				paramsContent = fmt.Sprintf("%s = ", paramsContent)
			} else {
				if len(params) == 1 {
					paramsContent = fmt.Sprintf("%s = ", paramsContent)
				} else {
					paramsContent = fmt.Sprintf("%s := ", paramsContent)
				}
			}
			return paramsContent
		},
	}
}

func (s *ServiceDesc) execute() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("api").Funcs(s.MapFn()).Parse(strings.TrimSpace(httpTemplate))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}

// bizExecute
func (s *ServiceDesc) bizExecute() string {
	s.LowerServiceType = util.LowerFirst(s.ServiceType)
	//小于三个的
	var useMessage []*protogen.Message
	for key, v := range s.MessageMap {
		if !util.InArray[string](key, s.MessageLessThree) {
			useMessage = append(useMessage, v)
		}
	}
	s.UseMessage = useMessage
	s.LowerIndex = util.LowerIndex(s.ServiceType)
	buf := new(bytes.Buffer)
	tmpl, err := template.New("biz").Funcs(s.MapFn()).Parse(strings.TrimSpace(bizTemplate))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}

// modelExecute
func (s *ServiceDesc) modelExecute() string {
	s.LowerServiceType = util.LowerFirst(s.ServiceType)
	//小于三个的
	var useMessage []*protogen.Message
	for key, v := range s.MessageMap {
		if !util.InArray[string](key, s.MessageLessThree) {
			useMessage = append(useMessage, v)
		}
	}
	s.UseMessage = useMessage
	s.LowerIndex = util.LowerIndex(s.ServiceType)
	buf := new(bytes.Buffer)
	tmpl, err := template.New("model").Funcs(s.MapFn()).Parse(strings.TrimSpace(modelTemplate))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}
