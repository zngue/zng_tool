package main

import (
	"fmt"
	"github.com/emicklei/proto"
	"log"
	"os"
)

type ServiceDesc struct {
	ServiceName string
	Methods     []*MethodDesc
	MessageMap  map[string]*proto.Message
}
type MethodDesc struct {
	Name        string
	RequestType string
	ReturnType  string
	Comment     string
}

func main() {
	path := "api/test/v1/test.proto"
	reader, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(reader *os.File) {
		err = reader.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(reader)
	var (
		definition *proto.Proto
		parser     *proto.Parser
	)
	parser = proto.NewParser(reader)
	definition, err = parser.Parse()
	if err != nil {
		log.Fatal(err)
		return
	}
	elements := definition.Elements
	var (
		messageVal = make(map[string]*proto.Message)
		serviceVal []*proto.Service
		enumVal    = make(map[string]*proto.Enum)
	)

	for _, element := range elements {
		switch e := element.(type) {
		case *proto.Service:
			serviceVal = append(serviceVal, e)
		case *proto.Message:
			messageVal[e.Name] = e
		case *proto.Enum:
			enumVal[e.Name] = e
		default:
			continue
		}
	}
	for _, service := range serviceVal {
		ServerTmp(service, messageVal)
	}

	fmt.Println(enumVal, messageVal, serviceVal)
}

func ServerTmp(service *proto.Service, messageMap map[string]*proto.Message) {
	var sc = &ServiceDesc{
		ServiceName: service.Name,
		MessageMap:  messageMap,
	}
	var methods []*MethodDesc
	for _, element := range service.Elements {
		val, ok := element.(*proto.RPC)
		if !ok {
			continue
		}
		methods = append(methods, &MethodDesc{
			Name:        val.Name,
			RequestType: val.RequestType,
			ReturnType:  val.ReturnsType,
			Comment:     val.Comment.Message(),
		})
		fmt.Println(element)
	}
	sc.Methods = methods
	fmt.Println(sc)

}
