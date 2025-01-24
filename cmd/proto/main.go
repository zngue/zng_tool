package main

import (
	"github.com/emicklei/proto"
	"log"
	"os"
)

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

func main() {
	path := "api/gin-pb/v1/gin-pb.proto"
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
}

func ServerTmp(service *proto.Service, messageMap map[string]*proto.Message) (sc *ServiceDesc, err error) {
	sc = &ServiceDesc{
		ServiceName: service.Name,
		MessageMap:  messageMap,
	}
	var methods []*MethodDesc
	for _, element := range service.Elements {
		val, ok := element.(*proto.RPC)
		if !ok || val == nil {
			continue
		}
		methods = append(methods, &MethodDesc{
			Name:           val.Name,
			RequestType:    val.RequestType,
			ReturnType:     val.ReturnsType,
			ReturnDefault:  "rs",
			RequestDefault: "req",
		})
	}
	var inMessageName, outMessageName []string
	for _, method := range methods {
		inMessageName = append(inMessageName, method.RequestType)
		outMessageName = append(outMessageName, method.ReturnType)
	}
	var doMessageName []string
	for key, message := range messageMap {
		if !inArray(key, inMessageName) && !inArray(key, outMessageName) {
			doMessageName = append(doMessageName, key)
			continue
		}
		if inArray(key, inMessageName) {
			if len(message.Elements) > 3 {
				doMessageName = append(doMessageName, key)
				continue
			}
		}
		if inArray(key, outMessageName) {
			if len(message.Elements) > 3 {
				doMessageName = append(doMessageName, key)
				continue
			}
		}
	}
	var messageDescItems []*MessageDesc
	for _, val := range doMessageName {
		msg := messageMap[val]
		if msg != nil {
			var filedSec []*FiledSec
			for _, element := range msg.Elements {
				v, ok := element.(*proto.NormalField)
				if !ok {
					continue
				}
				if v != nil {
					file := &FiledSec{
						Name:       v.Name,
						Type:       v.Type,
						IsRepeated: v.Repeated,
					}
					filedSec = append(filedSec, file)
				}
			}
			mg := &MessageDesc{
				Name:   msg.Name,
				Fields: filedSec,
			}
			messageDescItems = append(messageDescItems, mg)
		}
	}
	sc.DoMessage = messageDescItems
	sc.Methods = methods
	return
	//tmp := sc.execute()
	//fmt.Println(tmp)
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

func inArray(val string, arr []string) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}
