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
	Message       *protogen.Message
	ServerType    string
	GeneratedFile *protogen.GeneratedFile
	MessageMap    map[string]*protogen.Message
	//一下是自定义内容
	RelyMessage    string
	MessageContent string
	VarContent     string
	DataItems      []*ReplyItem
	OtherItems     []*ReplyItem
}
type ReplyItem struct {
	MessageType string
	GoName      string
	IsRepeated  bool
	LowerGoName string
	GoKind      string
}

func (r *Reply) Init() {
	var (
		items      []*ReplyItem
		otherItems []*ReplyItem
		key        = fmt.Sprintf("%sItem", r.ServerType)
	)
	r.RelyMessage = r.GeneratedFile.QualifiedGoIdent(r.Message.GoIdent)
	var varVal []string
	for _, field := range r.Message.Fields {
		msgType := util.MsgType(field)
		if msgType == util.AutoNormal || msgType == util.AutoRepeated {
			var isRepeated = msgType == util.AutoRepeated
			kind := r.GeneratedFile.QualifiedGoIdent(field.Message.GoIdent)
			lowerName := util.LowerFirst(field.GoName)
			if isRepeated {
				varVal = append(varVal, fmt.Sprintf("var %sVal []*%s", lowerName, kind))
			} else {
				varVal = append(varVal, fmt.Sprintf("var %sVal *%s", lowerName, kind))
			}
		}
	}
	for _, field := range r.Message.Fields {
		msgType := util.MsgType(field)
		if msgType == util.AutoNormal || msgType == util.AutoRepeated {
			kind := field.Message.GoIdent.GoName
			val := &ReplyItem{
				MessageType: kind,
				GoName:      field.GoName,
				IsRepeated:  msgType == util.AutoRepeated,
				LowerGoName: util.LowerFirst(field.GoName),
				GoKind:      r.GeneratedFile.QualifiedGoIdent(field.Message.GoIdent),
			}
			if kind == key {
				items = append(items, val)
			} else {
				otherItems = append(otherItems, val)
			}
		}
	}
	var params []string
	for _, field := range r.Message.Fields {
		msgType := util.MsgType(field)
		if msgType == util.AutoNormal || msgType == util.AutoRepeated {
			params = append(params, fmt.Sprintf("%s:%sVal,", field.GoName, util.LowerFirst(field.GoName)))
		} else {
			params = append(params, fmt.Sprintf("%s:%s,", field.GoName, util.LowerFirst(field.GoName)))
		}
	}
	r.MessageContent = strings.Join(params, "\n\t")
	r.VarContent = strings.Join(varVal, "\n\t")
	r.OtherItems = otherItems
	r.DataItems = items
}

func (r *Reply) MapFn() template.FuncMap {
	return template.FuncMap{
		"StructContent": func(messageType string, key string) string {
			var params []string
			if message, ok := r.MessageMap[messageType]; ok {
				for _, field := range message.Fields {
					msgType := util.MsgType(field)
					if msgType == util.AutoNormal || msgType == util.AutoRepeated {
						params = append(params, fmt.Sprintf("//todo %s %s", field.GoName, field.Message.GoIdent.GoName))
					} else {
						params = append(params, fmt.Sprintf("%s:%s.%s,", field.GoName, key, field.GoName))
					}
				}
			}
			return strings.Join(params, "\n\t")
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
