package main

import (
	"fmt"
	"github.com/emicklei/proto"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"os"
	"strings"
)

func main() {
	path := "api/test/v1/test.proto"
	reader, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer reader.Close()

	parser := proto.NewParser(reader)
	definition, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
		return
	}

	proto.Walk(
		definition,
		proto.WithMessage(func(message *proto.Message) {
			WithMessage(message)
		}),
		proto.WithService(WithService),
	)
	fmt.Println(serverList)
	fmt.Println("inMessage", inMessage)
	fmt.Println("outMessage", outMessage)
	for key, v := range serverList {
		for _, m := range v.Methods {
			fmt.Println(fmt.Sprintf("%s->%s", key, m.Name), m.Name, m.Request, m.Reply, m.Type)
		}
	}
	var allMessage []string
	for key, val := range messageList {
		allMessage = append(allMessage, key)
		for _, m := range val {
			fmt.Println(key, m.Type, m.Name)
		}
	}
	fmt.Println("messageList", allMessage)

}

var messageList = make(map[string][]*MessageItem)
var serverList = make(map[string]*Service)
var inMessage, outMessage []string

type Service struct {
	Package     string
	Service     string
	Methods     []*Method
	GoogleEmpty bool
	UseIO       bool
	UseContext  bool
}
type MethodType uint8

const (
	unaryType          MethodType = 1
	twoWayStreamsType  MethodType = 2
	requestStreamsType MethodType = 3
	returnsStreamsType MethodType = 4
)

type Method struct {
	Service string
	Name    string
	Request string
	Reply   string
	Type    MethodType
}

type MessageItem struct {
	Type       string
	Name       string
	IsRepeated bool
}

func serviceName(name string) string {
	return toUpperCamelCase(strings.Split(name, ".")[0])
}
func toUpperCamelCase(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = cases.Title(language.Und, cases.NoLower).String(s)
	return strings.ReplaceAll(s, " ", "")
}
func parametersName(name string) string {
	return strings.ReplaceAll(name, ".", "_")
}
func getMethodType(streamsRequest, streamsReturns bool) MethodType {
	if !streamsRequest && !streamsReturns {
		return unaryType
	} else if streamsRequest && streamsReturns {
		return twoWayStreamsType
	} else if streamsRequest {
		return requestStreamsType
	} else if streamsReturns {
		return returnsStreamsType
	}
	return unaryType
}
func WithService(service *proto.Service) {
	var srName = serviceName(service.Name)
	var cs = &Service{
		Package: "",
		Service: srName,
	}
	for _, element := range service.Elements {
		r, ok := element.(*proto.RPC)
		if !ok {
			continue
		}
		if r.RequestType == "google.protobuf.Empty" {
			continue
		}
		inMessage = append(inMessage, parametersName(r.RequestType))
		outMessage = append(outMessage, parametersName(r.ReturnsType))
		cs.Methods = append(cs.Methods, &Method{
			Service: srName,
			Name:    serviceName(r.Name),
			Request: parametersName(r.RequestType),
			Reply:   parametersName(r.ReturnsType),
			Type:    getMethodType(r.StreamsRequest, r.StreamsReturns),
		})
	}
	serverList[srName] = cs
}
func WithMessage(message *proto.Message) {
	var messageItems []*MessageItem

	for _, element := range message.Elements {

		if message.Name == "UserList" {
			fmt.Println(message.Name)
		}
		switch e := element.(type) {
		case *proto.NormalField:
			messageItems = append(messageItems, &MessageItem{
				Type:       e.Type,
				Name:       e.Name,
				IsRepeated: e.Repeated,
			})
		case *proto.MapField:
			messageItems = append(messageItems, &MessageItem{
				Type: e.KeyType,
				Name: e.Name,
			})
		}
	}
	var parentName = ""
	switch e := message.Parent.(type) {
	case *proto.NormalField:
		parentName = e.Name
	case *proto.MapField:
		parentName = e.Name
	}
	fmt.Println(parentName)
	messageList[message.Name] = messageItems
}

type optionLister struct {
	proto.NoopVisitor
}

func (l optionLister) VisitOption(o *proto.Option) {
	fmt.Println(o.Name)
}
