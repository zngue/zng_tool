package types

import "github.com/emicklei/proto"

type Template interface {
	Execute() string
}

type ServiceDesc struct {
	ServiceName string
	Methods     []*MethodDesc
	MessageMap  map[string]*proto.Message
	DoMessage   []*MessageDesc
}
type MethodDesc struct {
	Name           string
	RequestType    string
	ReturnType     string
	ReturnDefault  string
	RequestDefault string
	Comment        string
}
type MessageDesc struct {
	Name   string
	Fields []*FiledSec
}
type FiledSec struct {
	Name       string
	Type       string
	IsRepeated bool
}
