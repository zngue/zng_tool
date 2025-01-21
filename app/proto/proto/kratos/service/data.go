package service

import (
	"github.com/emicklei/proto"
	"github.com/zngue/zng_tool/app/proto/proto/types"
)

type Template struct {
	ServiceName string
	Methods     []*types.MethodDesc
	MessageMap  map[string]*proto.Message
	DoMessage   []*types.MessageDesc
}

func NewDataTemplate(data *types.ServiceDesc) types.Template {
	sc := &Template{
		ServiceName: data.ServiceName,
		Methods:     data.Methods,
		MessageMap:  data.MessageMap,
		DoMessage:   data.DoMessage,
	}
	return sc
}
