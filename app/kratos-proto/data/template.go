package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/emicklei/proto"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"
)

//go:embed template.tpl
var httpTemplate string

// LowerFirst 首字母转小写
func LowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

// UpperFirst 首字母转大写
func UpperFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}

type FileType int

const (
	AutoRepeated   FileType = 1 //自定义数组
	AutoNormal     FileType = 2 //自定义
	SystemRepeated FileType = 3 //系统数组
	SystemNormal   FileType = 4 //系统
)

func DoParamsFile(msg *proto.Message, v *proto.NormalField) FileType {
	if msg != nil {
		if v.Repeated {
			return AutoRepeated
		} else {
			return AutoNormal
		}
	} else {
		if v.Repeated {
			return SystemRepeated
		} else {
			return SystemNormal
		}
	}
}
func (s *ServiceDesc) execute() string {
	buf := new(bytes.Buffer)
	funcMap := template.FuncMap{
		"MessageFile": func(fileName, fileType string, isRepeated bool) string {
			msg := s.MessageMap[fileType]
			if msg != nil { //自定义
				if isRepeated {
					return fmt.Sprintf("%s []*biz.%s", UpperFirst(fileName), fileType)
				} else {
					return fmt.Sprintf("%s *biz.%s", UpperFirst(fileName), fileType)
				}
			} else {
				if isRepeated {
					return fmt.Sprintf("%s []%s", UpperFirst(fileName), fileType)
				} else {
					return fmt.Sprintf("%s %s", UpperFirst(fileName), fileType)
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
			return LowerFirst(s)
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
					msgType := DoParamsFile(msg, v)
					switch msgType {
					case AutoRepeated:
						params = append(params, fmt.Sprintf("%s []*biz.%s", v.Name, v.Type))
					case AutoNormal:
						params = append(params, fmt.Sprintf("%s *biz.%s", v.Name, v.Type))
					case SystemRepeated:
						params = append(params, fmt.Sprintf("%s []%s", v.Name, v.Type))
					case SystemNormal:
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
