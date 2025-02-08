package util

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
)

func MsgType(field *protogen.Field) FileType {
	if field.Desc.Kind().String() == "message" {
		if field.Desc.Cardinality().String() == "repeated" {
			return AutoRepeated
		} else {
			return AutoNormal
		}
	} else {
		if field.Desc.Cardinality().String() == "repeated" {
			return SystemRepeated
		} else {
			return SystemNormal
		}
	}
}
func StructType(req *protogen.Field) (val string) {
	msgType := MsgType(req)
	var kind = req.Desc.Kind().String()
	if kind == "message" {
		kind = req.Message.GoIdent.GoName
	}
	if req.Extendee != nil {
		kind = fmt.Sprintf("%s_%s", req.Extendee.GoIdent.GoName, kind)
	}
	switch msgType {
	case AutoRepeated:
		val = fmt.Sprintf("%s []*%s", req.GoName, kind)
	case AutoNormal:
		val = fmt.Sprintf("%s *%s ", req.GoName, kind)
	case SystemRepeated:
		val = fmt.Sprintf("%s []%s ", req.GoName, kind)
	case SystemNormal:
		val = fmt.Sprintf("%s %s ", req.GoName, kind)
	default:
		val = fmt.Sprintf("%s %s ", req.GoName, kind)
	}
	return

}
