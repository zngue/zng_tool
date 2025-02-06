package app

import (
	"fmt"
	"github.com/zngue/zng_tool/app/util"
	"google.golang.org/protobuf/compiler/protogen"
)

type PluginOption struct {
	GoImportPath string
	Name         string
	AutoName     string
	IsSkip       bool
	Fn           ServiceFn
}

type ServiceDesc struct {
	LowerServiceType string
	LowerIndex       string
	ServiceType      string
	ServiceName      string
	ServiceTypeName  string
	Metadata         string
	Methods          []*MethodDesc
	GoPackageName    string
	GoImportPath     string
	MessageMap       map[string]*protogen.Message
	GeneratedFile    *protogen.GeneratedFile
	LowerServiceName string
	MessageLessThree []string
	UseMessage       []*protogen.Message
}
type MethodDesc struct {
	Name           string
	OriginalName   string
	MethodIndex    int
	ServerIndex    int
	Request        protogen.GoIdent
	RequestMessage *protogen.Message
	RequestDefault string
	Reply          protogen.GoIdent
	ReplyDefault   string
	ReplyMessage   *protogen.Message
	GoPackageName  string
	ReplyLent      int
	RequestLent    int
}
type ServiceFn func(g *protogen.GeneratedFile, service *ServiceDesc) (err error)

func Plugin(gen *protogen.Plugin, options ...*PluginOption) {
	for _, file := range gen.Files {
		for _, option := range options {
			var fileName string
			if option.AutoName != "" {
				fileName = fmt.Sprintf("%s.go", option.AutoName)
			} else {
				if option.Name != "" {
					fileName = fmt.Sprintf("%s.%s.pb.go", file.GeneratedFilenamePrefix, option.Name)
				} else {
					fileName = file.GeneratedFilenamePrefix + ".api.pb.go"
				}
			}
			var g *protogen.GeneratedFile
			if option.GoImportPath != "" {
				g = gen.NewGeneratedFile(fileName, protogen.GoImportPath(option.GoImportPath))
			} else {
				g = gen.NewGeneratedFile(fileName, file.GoImportPath)
			}
			NewPluginFile(g, file, option.Fn)
			if option.IsSkip {
				g.Skip()
			}
		}
	}
	return
}

type PluginFile struct {
	MessageMap    map[string]*protogen.Message
	GeneratedFile *protogen.GeneratedFile
	File          *protogen.File
}

func NewPluginFile(generatedFile *protogen.GeneratedFile, file *protogen.File, fn ServiceFn) {
	p := &PluginFile{
		GeneratedFile: generatedFile,
		File:          file,
	}
	p.Services(fn)
}

func (p *PluginFile) Messages() {
	var messageMap = make(map[string]*protogen.Message)
	for _, message := range p.File.Messages {
		messageMap[message.GoIdent.GoName] = message
	}
	p.MessageMap = messageMap
	return
}
func (p *PluginFile) Services(fn ServiceFn) {
	p.Messages()
	data := &PluginFileServices{
		GeneratedFile: p.GeneratedFile,
		File:          p.File,
		MessageMap:    p.MessageMap,
	}
	data.Services(fn)
}

type PluginFileServices struct {
	MessageMap    map[string]*protogen.Message
	GeneratedFile *protogen.GeneratedFile
	File          *protogen.File
}

func (p *PluginFileServices) Services(fn ServiceFn) {
	for serverIndex, service := range p.File.Services {
		sd := p.Service(serverIndex, service)
		if err := fn(p.GeneratedFile, sd); err != nil {
			fmt.Println(err)
		}
	}
	return
}

func (p *PluginFileServices) Service(serverIndex int, service *protogen.Service) (sd *ServiceDesc) {
	methodItems, messageLessThree := p.Methods(serverIndex, service)
	sd = &ServiceDesc{
		ServiceType: service.GoName,
		ServiceName: string(service.Desc.FullName()),
		ServiceTypeName: p.GeneratedFile.QualifiedGoIdent(protogen.GoIdent{
			GoName:       service.GoName,
			GoImportPath: p.File.GoImportPath,
		}),
		Metadata:         p.File.Desc.Path(),
		GoPackageName:    string(p.File.GoPackageName),
		GoImportPath:     p.File.GoImportPath.String(),
		MessageMap:       p.MessageMap,
		GeneratedFile:    p.GeneratedFile,
		LowerServiceName: util.LowerFirst(service.GoName),
		Methods:          methodItems,
		MessageLessThree: messageLessThree,
	}
	return
}

func (p *PluginFileServices) Methods(serverIndex int, service *protogen.Service) (methodItems []*MethodDesc, messageLessThree []string) {
	for methodIndex, method := range service.Methods {
		if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
			continue
		}
		methodItems = append(methodItems, &MethodDesc{
			Name:           method.GoName,
			OriginalName:   string(method.Desc.Name()),
			MethodIndex:    methodIndex,
			ServerIndex:    serverIndex,
			RequestDefault: "req",
			Request:        method.Input.GoIdent,
			ReplyDefault:   "reply",
			Reply:          method.Output.GoIdent,
			ReplyMessage:   method.Output,
			RequestMessage: method.Input,
			ReplyLent:      len(method.Output.Fields),
			RequestLent:    len(method.Input.Fields),
		})
		//判断是否在
		if !util.InArray(method.Output.GoIdent.GoName, messageLessThree) && len(method.Output.Fields) < 3 {
			messageLessThree = append(messageLessThree, method.Output.GoIdent.GoName)
		}
		if !util.InArray(method.Input.GoIdent.GoName, messageLessThree) && len(method.Input.Fields) < 3 {
			messageLessThree = append(messageLessThree, method.Input.GoIdent.GoName)
		}
	}
	return
}
