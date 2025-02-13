package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/zngue/zng_tool/app/util"
	"github.com/zngue/zng_tool/cmd/protoc-gen-gin-tmp/tmp"
	"github.com/zngue/zng_tool/third_party/google/annotations"
	"github.com/zngue/zng_tool/third_party/validate"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"strings"
	"text/template"
)

//go:embed tmp/template.tpl
var httpTemplate string

//go:embed tmp/biz_template.tpl
var bizTemplate string

//go:embed tmp/model_template.tpl
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
	Action         annotations.Action
	SvrType        string //使用的时候赋值
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
	MessageMap     map[string]*protogen.Message //使用前请先赋值
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
				kind = element.Message.GoIdent.GoName
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

//追加写入文件

func DoFieldOperator(req *protogen.Field) (operator validate.Operator, filedType string) {
	options, ok := req.Desc.Options().(*descriptorpb.FieldOptions)
	if !ok || options == nil {
		return
	}
	rules := proto.GetExtension(options, validate.E_Rules)
	fieldRules, ok := rules.(*validate.FieldRules)
	if !ok || fieldRules == nil {
		return
	}
	kind := req.Desc.Kind()
	msgType := util.MsgType(req)
	var msgKind = req.Desc.Kind().String()
	if msgType == util.SystemRepeated {
		if strings.Contains(msgKind, "int") {
			filedType = "number"
		} else {
			filedType = "string"
		}
		operator = fieldRules.GetRepeated().GetOperator()
		return
	}

	switch kind {
	case protoreflect.Int32Kind, protoreflect.Int64Kind:
		if intRules := fieldRules.GetInt32(); intRules != nil {
			operator = intRules.GetOperator()
			filedType = "number"
		}
	case protoreflect.FloatKind:
		if floatRules := fieldRules.GetFloat(); floatRules != nil {
			operator = floatRules.GetOperator()
			filedType = "number"
		}

	case protoreflect.Uint32Kind, protoreflect.Uint64Kind:
		if uintRules := fieldRules.GetUint32(); uintRules != nil {
			operator = uintRules.GetOperator()
			filedType = "number"
		}
	case protoreflect.StringKind:
		if stringRules := fieldRules.GetString_(); stringRules != nil {
			operator = stringRules.GetOperator()
			filedType = "string"
		}
	}
	return
}

type IsItem struct {
	Flag             bool
	MessageType      string
	StructSetContent string
}

var reqTemplate = `


`

func (s *ServiceDesc) MapFn() template.FuncMap {
	return template.FuncMap{
		"ServiceReq": func(req *protogen.Message) string {
			if len(req.Fields) > 3 {

			}
			return ""
		},
		"ModelContent": func(method *MethodDesc, svrType string) string {
			method.SvrType = svrType
			method.MessageMap = s.MessageMap
			return method.execute()
		},
		"IsItem": func(str string) (val *IsItem) {
			val = &IsItem{}
			key := fmt.Sprintf("%sItem", s.ServiceType)
			var (
				messageItem *protogen.Message
				params      []string
			)
			if messageItem, val.Flag = s.MessageMap[key]; val.Flag {
				val.MessageType = s.GeneratedFile.QualifiedGoIdent(messageItem.GoIdent)
				if len(messageItem.Fields) > 0 {
					for _, field := range messageItem.Fields {
						params = append(params, fmt.Sprintf("%s:req.%s,", field.GoName, field.GoName))
					}
				}
				if str != "" {
					val.StructSetContent = strings.Join(params, str)
				} else {
					val.StructSetContent = strings.Join(params, "\n\t")
				}

				return
			}
			return
		},
		"StructName": func(req *protogen.Message) string {
			return string(req.Desc.FullName())
		},
		"StructType": util.StructType,
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
						msgType := util.MsgType(element)
						if msgType == util.AutoRepeated || msgType == util.AutoNormal {
							params = append(params, fmt.Sprintf("req%s", element.GoName))
						} else {
							params = append(params, fmt.Sprintf("req.%s", element.GoName))
						}
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
		"AutoRequest": func(req *protogen.Message) string {
			var data = &tmp.Request{
				Message:    req,
				MessageMap: s.MessageMap,
			}
			return data.Execute()
		},
		"ServiceReplyContent": func(req *protogen.Message) string {
			if len(req.Fields) == 0 {
				return ""
			}
			data := &tmp.Reply{
				Message:       req,
				ServerType:    s.ServiceType,
				GeneratedFile: s.GeneratedFile,
			}
			return data.Execute()
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
		if !util.InArray[string](key, s.MessageLessThree) && key != fmt.Sprintf("%sDB", s.ServiceType) {
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
