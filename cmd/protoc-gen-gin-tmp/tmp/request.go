package tmp

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/zngue/zng_tool/app/util"
	"google.golang.org/protobuf/compiler/protogen"
	"os"
	"strings"
	"text/template"
)

type Request struct {
	Message     *protogen.Message
	MessageName string
	MessageMap  map[string]*protogen.Message
}

type RepeatedFiled struct {
	Name        string
	MessageType string
	GoName      string
	NormalName  string //设置请求参数名称

}

func (r *Request) MapFn() template.FuncMap {
	return template.FuncMap{
		"MessageFileType": func(req *protogen.Message) string {
			var params []string
			for _, field := range req.Fields {
				msgType := util.MsgType(field)
				if msgType == util.AutoRepeated {
					params = append(params, fmt.Sprintf("var req%s []*biz.%s", field.GoName, field.Message.GoIdent.GoName))
				} else if msgType == util.AutoNormal {
					params = append(params, fmt.Sprintf("var req%s *biz.%s", field.GoName, field.Message.GoIdent.GoName))
				}
			}
			return strings.Join(params, "\n\t")
		},
		"AutoRepeated": func(req *protogen.Field) bool {
			msgType := util.MsgType(req)
			return msgType == util.AutoRepeated
		},
		"AutoNormal": func(req *protogen.Field) bool {
			msgType := util.MsgType(req)
			return msgType == util.AutoNormal
		},
		"AutoFiled": func(req *protogen.Field) *RepeatedFiled {
			return &RepeatedFiled{
				Name:        fmt.Sprintf("req%s", req.GoName),
				MessageType: req.Message.GoIdent.GoName,
				GoName:      req.GoName,
				NormalName:  fmt.Sprintf("req.%s", req.GoName),
			}
		},
		"AutoFileStruct": func(messageType string, key string) string {
			messageItem := r.MessageMap[messageType]
			if messageItem == nil {
				return ""
			}
			var params []string
			for _, field := range messageItem.Fields {
				msgType := util.MsgType(field)
				if msgType == util.AutoNormal || msgType == util.AutoRepeated {
					params = append(params, fmt.Sprintf("//todo %s %s", field.GoName, field.Message.GoIdent.GoName))
				} else {
					params = append(params, fmt.Sprintf("%s:%s.%s,", field.GoName, key, field.GoName))
				}
			}
			return strings.Join(params, "\n\t")
		},
		"RequestName": func(req *protogen.Message) string {
			return req.GoIdent.GoName
		},
		"RequestStruct": func(req *protogen.Message) string {
			var params []string
			for _, field := range req.Fields {
				msgType := util.MsgType(field)
				if msgType == util.AutoNormal || msgType == util.AutoRepeated {
					params = append(params, fmt.Sprintf("%s:req%s,", field.GoName, field.GoName))
				} else {
					params = append(params, fmt.Sprintf("%s:req.%s,", field.GoName, field.GoName))
				}
			}
			content := strings.Join(params, "\n\t\t")
			WriteContent("abcde.txt", content)
			return content
		},
	}
}
func WriteContent(fileName string, content string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	// 追加写入内容
	_, err = file.WriteString(content + "\n")
	if err != nil {
		return
	}
}

// 获取请求参数
//
//go:embed req_params_template.tpl
var reqParamsTemplate string

func (r *Request) ParamsExec() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("reqParamsTemplate").Funcs(r.MapFn()).Parse(strings.TrimSpace(reqParamsTemplate))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, r); err != nil {
		panic(err)
	}
	return buf.String()
}

//go:embed req_template.tpl
var reqTemplate string

func (r *Request) Execute() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("reqTemplate").Funcs(r.MapFn()).Parse(strings.TrimSpace(reqTemplate))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, r); err != nil {
		panic(err)
	}
	return buf.String()
}
