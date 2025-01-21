package data

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/emicklei/proto"
	"github.com/zngue/zng_tool/app/util"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"
)

//go:embed template.tpl
var httpTemplate string

func (s *Template) Execute() string {
	buf := new(bytes.Buffer)
	funcMap := template.FuncMap{
		"MessageFile": func(fileName, fileType string, isRepeated bool) string {
			msg := s.MessageMap[fileType]
			if msg != nil { //自定义
				if isRepeated {
					return fmt.Sprintf("%s []*biz.%s", util.UpperFirst(fileName), fileType)
				} else {
					return fmt.Sprintf("%s *biz.%s", util.UpperFirst(fileName), fileType)
				}
			} else {
				if isRepeated {
					return fmt.Sprintf("%s []%s", util.UpperFirst(fileName), fileType)
				} else {
					return fmt.Sprintf("%s %s", util.UpperFirst(fileName), fileType)
				}
			}
		},
		//获取首字母
		"FirstIndex": func(s string) string {
			if s == "" {
				return ""
			}
			r, _ := utf8.DecodeRuneInString(s)
			return string(unicode.ToLower(r))
		},
		"LowerFirst": func(s string) string {
			return util.LowerFirst(s)
		},
		"ParamsSet": func(defaultKey, methodType string) string {
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
					params = append(params, v.Name)
				}
				return strings.Join(params, ",")
			} else {
				return fmt.Sprintf("%s", defaultKey)
			}

		},
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
					msgType := util.DoParamsFile(msg, v)
					switch msgType {
					case util.AutoRepeated:
						params = append(params, fmt.Sprintf("%s []*biz.%s", v.Name, v.Type))
					case util.AutoNormal:
						params = append(params, fmt.Sprintf("%s *biz.%s", v.Name, v.Type))
					case util.SystemRepeated:
						params = append(params, fmt.Sprintf("%s []%s", v.Name, v.Type))
					case util.SystemNormal:
						params = append(params, fmt.Sprintf("%s %s", v.Name, v.Type))
					}
				}
				return strings.Join(params, ",")
			} else {
				return fmt.Sprintf("%s *biz.%s", defaultKey, methodType)
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
