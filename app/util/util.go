package util

import (
	"github.com/emicklei/proto"
	"os"
	"unicode"
	"unicode/utf8"
)

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

func WriteFile(fileName, content string) (err error) {
	var file *os.File
	file, err = os.Create(fileName)
	if err != nil {
		return
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {

		}
	}(file)
	_, err = file.WriteString(content)
	return
}

// UpperLineToLower 驼峰转下划线
func UpperLineToLower(s string) string {
	if s == "" {
		return ""
	}
	var result string
	for index, v := range s {
		if index == 0 {
			if unicode.IsUpper(v) {
				result += string(unicode.ToLower(v))
			} else {
				result += string(v)
			}
		} else {
			if unicode.IsUpper(v) {
				result += "_" + string(unicode.ToLower(v))
			} else {
				result += string(v)
			}
		}

	}
	return result
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
