package util

import "google.golang.org/protobuf/compiler/protogen"

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
