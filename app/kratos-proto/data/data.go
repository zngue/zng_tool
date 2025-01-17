package data

import (
	"github.com/emicklei/proto"
	"github.com/zngue/zng_tool/app/kratos-proto/types"
)

type ServiceDesc struct {
	ServiceName string
	Methods     []*types.MethodDesc
	MessageMap  map[string]*proto.Message
	DoMessage   []*types.MessageDesc
}

func NewDataTemplate(data *types.ServiceDesc) string {
	sc := &ServiceDesc{
		ServiceName: data.ServiceName,
		Methods:     data.Methods,
		MessageMap:  data.MessageMap,
		DoMessage:   data.DoMessage,
	}
	return sc.execute()
}
