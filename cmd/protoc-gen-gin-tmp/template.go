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

type ServiceDesc struct {
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
}
type MethodDesc struct {
	Name           string
	OriginalName   string
	MethodIndex    int
	ServerIndex    int
	Request        protogen.GoIdent
	RequestDefault string
	Reply          protogen.GoIdent
	ReplyDefault   string
	RequestMessage *protogen.Message
	ReplyMessage   *protogen.Message
	GoPackageName  string
	ReplyLent      int
	RequestLent    int
}

func (s *ServiceDesc) execute() string {
	buf := new(bytes.Buffer)
	funcMap := template.FuncMap{
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
		"OutParams": func(def string, req *protogen.Message) string {
			var params []string
			if len(req.Fields) > 3 {
				params = append(params, def)
			} else {
				for _, element := range req.Fields {
					params = append(params, util.LowerFirst(element.GoName))
				}
			}
			params = append(params, "err")
			return strings.Join(params, ",")
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
