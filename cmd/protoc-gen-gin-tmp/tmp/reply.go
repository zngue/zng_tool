package tmp

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/zngue/zng_tool/app/util"
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
	"text/template"
)

type Reply struct {
	Message        *protogen.Message
	ServerType     string
	DataItems      []*ReplyItem
	GeneratedFile  *protogen.GeneratedFile
	RelyMessage    string
	MessageContent string
}
type ReplyItem struct {
	MessageType string
	GoName      string
	IsRepeated  bool
	LowerGoName string
	GoKind      string
}

func (r *Reply) Init() {
	var key = fmt.Sprintf("%sItem", r.ServerType)
	var items []*ReplyItem
	r.RelyMessage = r.GeneratedFile.QualifiedGoIdent(r.Message.GoIdent)
	for _, field := range r.Message.Fields {
		msgType := util.MsgType(field)
		if msgType == util.AutoNormal || msgType == util.AutoRepeated {
			kind := field.Message.GoIdent.GoName
			if kind == key {
				items = append(items, &ReplyItem{
					MessageType: kind,
					GoName:      field.GoName,
					IsRepeated:  msgType == util.AutoRepeated,
					LowerGoName: util.LowerFirst(field.GoName),
					GoKind:      r.GeneratedFile.QualifiedGoIdent(field.Message.GoIdent),
				})
			}
		}
	}
	var params []string
	for _, field := range r.Message.Fields {
		msgType := util.MsgType(field)
		if msgType == util.AutoNormal || msgType == util.AutoRepeated {
			kind := field.Message.GoIdent.GoName
			if kind != key {
				params = append(params, fmt.Sprintf("//todo %s,", kind))
				continue
			}
			params = append(params, fmt.Sprintf("%s:%sVal,", field.GoName, util.LowerFirst(field.GoName)))
		} else {
			params = append(params, fmt.Sprintf("%s:%s,", field.GoName, util.LowerFirst(field.GoName)))
		}
	}
	r.MessageContent = strings.Join(params, "\n\t\t")
	r.DataItems = items
}

func (r *Reply) MapFn() template.FuncMap {
	return template.FuncMap{
		"MessageFileType": func(req *protogen.Message) string {
			return ""
		},
	}
}

//go:embed rely_template.tpl
var replyTemplate string

func (r *Reply) Execute() string {
	r.Init()
	buf := new(bytes.Buffer)
	tmpl, err := template.New("replyTemplate").Funcs(r.MapFn()).Parse(strings.TrimSpace(replyTemplate))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, r); err != nil {
		panic(err)
	}
	return buf.String()
}
