package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/zngue/zng_tool/app/util"
	"github.com/zngue/zng_tool/third_party/google/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

var showVersion = flag.Bool("version", false, "print the version and exit")

func Biz(gen *protogen.Plugin) (err error) {
	for _, file := range gen.Files {
		if !file.Generate {
			continue
		}
		fileName := file.GeneratedFilenamePrefix + ".gin_biz.pb.go"
		g := gen.NewGeneratedFile(fileName, "biz")
		//g.P("// Code generated by protoc-gen-gin-tmp. DO NOT EDIT.")
		g.P()
		var messageMap = make(map[string]*protogen.Message)
		for _, message := range file.Messages {
			messageMap[message.GoIdent.GoName] = message
		}

		//获取入参大于字段小于3个的message
		for serverIndex, service := range file.Services {
			var messageLessThree []string
			sd := &ServiceDesc{
				ServiceType: service.GoName,
				ServiceTypeName: g.QualifiedGoIdent(protogen.GoIdent{
					GoName:       service.GoName,
					GoImportPath: file.GoImportPath,
				}),
				LowerServiceName: util.LowerFirst(service.GoName),
				ServiceName:      string(service.Desc.FullName()),
				Metadata:         file.Desc.Path(),
				GoPackageName:    string(file.GoPackageName),
				GoImportPath:     file.GoImportPath.String(),
				MessageMap:       messageMap,
				GeneratedFile:    g,
			}
			var methodItems []*MethodDesc
			for methodIndex, method := range service.Methods {
				if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
					continue
				}
				rule := proto.GetExtension(method.Desc.Options(), annotations.E_Http).(*annotations.HttpRule)
				var action annotations.Action
				if rule != nil {
					action = rule.Action
				}
				methodItems = append(methodItems, &MethodDesc{
					Name:           method.GoName,
					Action:         action,
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
			sd.Methods = methodItems
			sd.MessageLessThree = messageLessThree
			tmp := sd.bizExecute()
			buffer := &bytes.Buffer{}
			buffer.WriteString(tmp)
			s := buffer.String()
			lowerName := util.UpperLineToLower(sd.ServiceType)
			err = WireFile("./internal/biz", lowerName, s)
			if err != nil {
				return
			}
			bizServiceName := fmt.Sprintf("New%sUseCase", sd.ServiceType)
			ReplaceWire("./internal/biz", "biz", bizServiceName, "biz")
			modelTmp := sd.modelExecute()
			modelBuffer := &bytes.Buffer{}
			modelBuffer.WriteString(modelTmp)
			modelContent := modelBuffer.String()
			err = WireFile("./internal/model", lowerName, modelContent)
			if err != nil {
				return
			}
			modelServiceName := fmt.Sprintf("New%sRepo", sd.ServiceType)
			ReplaceWire("./internal/model", "data", modelServiceName, "model")
			//创建db
			dbConn := DbReplace{
				MessageMap:      sd.MessageMap,
				ServerType:      sd.ServiceType,
				LowerServerType: util.CamelToSnake(sd.ServiceType),
				Pkg:             "db",
			}
			dbContent := dbConn.execute()
			dbReplace("./internal/model/db", lowerName, dbContent)
		}
		g.Skip()
	}
	return
}

func Services(gen *protogen.Plugin) (err error) {
	for _, file := range gen.Files {
		if !file.Generate {
			continue
		}
		fileName := file.GeneratedFilenamePrefix + ".gin_api.pb.go"
		g := gen.NewGeneratedFile(fileName, "api")
		//g.P("// Code generated by protoc-gen-gin-tmp. DO NOT EDIT.")
		g.P()
		var messageMap = make(map[string]*protogen.Message)
		for _, message := range file.Messages {
			messageMap[message.GoIdent.GoName] = message
		}
		buffer := &bytes.Buffer{}
		//获取入参大于字段小于3个的message
		for serverIndex, service := range file.Services {
			var messageLessThree []string
			sd := &ServiceDesc{
				ServiceType: service.GoName,
				ServiceTypeName: g.QualifiedGoIdent(protogen.GoIdent{
					GoName:       service.GoName,
					GoImportPath: file.GoImportPath,
				}),
				LowerServiceName: util.LowerFirst(service.GoName),
				ServiceName:      string(service.Desc.FullName()),
				Metadata:         file.Desc.Path(),
				GoPackageName:    string(file.GoPackageName),
				GoImportPath:     file.GoImportPath.String(),
				MessageMap:       messageMap,
				GeneratedFile:    g,
			}
			var methodItems []*MethodDesc
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
			sd.MessageLessThree = messageLessThree
			sd.Methods = methodItems
			tmp := sd.execute()
			g.P(tmp)
			var content []byte
			content, err = g.Content()
			if err != nil {
				fmt.Println("err", err)
				return nil
			}
			buffer.WriteString(string(content))
			s := buffer.String()
			lowerName := util.UpperLineToLower(sd.ServiceType)
			var dir = "./internal/api"
			err = WireFile(dir, lowerName, s)
			if err != nil {
				fmt.Println("err", err)
				return nil
			}
			serverName := fmt.Sprintf("New%sService", sd.ServiceType)
			ReplaceWire(dir, "api", serverName, "api")
		}
		g.Skip()
	}
	return
}

func main() {
	opts := &protogen.Options{}
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-gin-http %v\n", "v1.0.1")
		return
	}
	opts.Run(func(gen *protogen.Plugin) error {
		err := Services(gen)
		if err != nil {
			return err
		}
		err = Biz(gen)
		if err != nil {
			return err
		}
		return nil
	})
}
