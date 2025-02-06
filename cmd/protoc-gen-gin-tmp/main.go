package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/zngue/zng_tool/app/util"
	"google.golang.org/protobuf/compiler/protogen"
	"os"
	"regexp"
	"strings"
	"text/template"
)

var showVersion = flag.Bool("version", false, "print the version and exit")

const ReqTemp = `
var reqData=Name{
	
}
`

func ReqTempFn() {
	tmpl, err := template.New("service").Parse(strings.TrimSpace(ReqTemp))
	if err != nil {
		panic(err)
	}
	fmt.Println(tmpl)
}

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
				//MessageLessThree: messageLessThree,
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
			//util.IsDir("./internal/biz")
			//err = util.WriteFile(fmt.Sprintf("./internal/biz/%s.go", lowerName), s)
			//if err != nil {
			//	fmt.Println("err", err)
			//	return nil
			//}
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
			//util.IsDir("./internal/model")
			//err = util.WriteFile(fmt.Sprintf("./internal/model/%s.go", lowerName), modelContent)
			//if err != nil {
			//	fmt.Println("err", err)
			//	return nil
			//}
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
				//MessageLessThree: messageLessThree,
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

const wireTemplate = `
package {{PKG}}
import (
	"github.com/google/wire"
)
var ProviderSet = wire.NewSet(
	{{CONTENT}},
)`

func ReplaceWire(dir, fileName, serverName string, pkg string) {
	fileName = fmt.Sprintf("%s/%s.go", dir, fileName)
	//文件不存在则创建文件
	if !util.FileExists(fileName) {
		tmp := strings.ReplaceAll(wireTemplate, "{{PKG}}", pkg)
		tmp = strings.ReplaceAll(tmp, "{{CONTENT}}", serverName)
		err := util.WriteFile(fileName, tmp)
		if err != nil {
			return
		}
		return
	}
	readFile, err := os.ReadFile(fileName)
	if err != nil {
		return
	}
	re := regexp.MustCompile(`wire\.NewSet\(([\s\S]*?)\)`)
	matches := re.FindStringSubmatch(string(readFile))
	if len(matches) > 1 {
		// matches[0] 是整个匹配的字符串，matches[1] 是括号内的内容
		//将字符使用逗号分隔 并且去掉空格，和换行
		var params []string
		for _, param := range strings.Split(matches[1], ",") {
			param = strings.TrimSpace(param)
			if param != "" && !util.InArray(param, params) {
				params = append(params, param)
			}
		}
		////将新的加入
		if !util.InArray(serverName, params) {
			params = append(params, serverName)
		}
		////将新的替换
		newContent := "\n\t" + strings.Join(params, ",\n\t") + ",\n"
		newContent = strings.Replace(string(readFile), matches[1], newContent, 1)
		err = util.WriteFile(fileName, newContent)
		if err != nil {
			return
		}
	} else {
		tmp := strings.ReplaceAll(wireTemplate, "{{PKG}}", pkg)
		tmp = strings.ReplaceAll(tmp, "{{CONTENT}}", serverName)
		err = util.WriteFile(fileName, tmp)
		if err != nil {
			return
		}
	}
}

func WireFile(dir, fileName, content string) (err error) {
	util.IsDir(dir)
	return util.WriteFile(fmt.Sprintf("%s/%s.go", dir, fileName), content)
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
