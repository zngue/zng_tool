package main

import (
	"flag"
	"fmt"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strings"
)

var showVersion = flag.Bool("version", false, "print the version and exit")

func main() {
	opts := &protogen.Options{}
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-gin-http %v\n", "v1.0.1")
		return
	}
	opts.Run(func(gen *protogen.Plugin) error {
		files := gen.Files
		for _, file := range files {
			fileName := file.GeneratedFilenamePrefix + ".gin_http.pb.go"
			g := gen.NewGeneratedFile(fileName, file.GoImportPath)
			g.P("// Code generated by protoc-gen-gin-http. DO NOT EDIT.")
			g.P("package ", file.GoPackageName)
			g.P()
			for _, message := range file.Messages {
				if message.Desc.IsMapEntry() {
					continue
				}
				//获取参数验证
				if message.Desc.IsMapEntry() {
					continue
				}
				if message.Desc.IsMapEntry() {
					continue
				}
			}
			for serverIndex, service := range file.Services {
				sd := &ServiceDesc{
					ServiceType: service.GoName,
					Comment:     service.Comments.Leading.String(),
					ServiceName: string(service.Desc.FullName()),
					Metadata:    file.Desc.Path(),
				}
				for methodIndex, method := range service.Methods {
					if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
						continue
					}
					rule, ok := proto.GetExtension(method.Desc.Options(), annotations.E_Http).(*annotations.HttpRule)
					if rule != nil && ok {
						for _, bind := range rule.AdditionalBindings {
							sd.Methods = append(sd.Methods, buildHTTPRule(serverIndex, methodIndex, service, method, bind))
						}
						sd.Methods = append(sd.Methods, buildHTTPRule(serverIndex, methodIndex, service, method, rule))
					} else if omitemptyPrefix != "" {
						path := fmt.Sprintf("%s/%s/%s", omitemptyPrefix, service.Desc.FullName(), method.Desc.Name())
						sd.Methods = append(sd.Methods, buildMethodDesc(serverIndex, methodIndex, method, http.MethodGet, path))
					}
				}
				serverContent := sd.execute()
				g.P(serverContent)
			}
		}
		return nil
	})
}

type MethodInfo struct {
	Method      string
	Path        string
	IsPost      bool
	ServerIndex int
	ServerName  string
}

const routerTemplate = `
type {{serverName}}GinHttpRouterService struct {
	srv {{serverName}}GinHttpService
	router *gin.RouterGroup
}
`

const routerFnTemplate = `
func (s *{{serverName}}GinHttpRouterService) Register() []router.IRouter {
	return router.ApiServiceFn(
		{{FN}}
	)
}
`

const registerTemplate = `
func Register{{serverName}}GinHttpRouterService(
	api *gin.RouterGroup,
	srv {{serverName}}GinHttpService,
	) router.IApiService  {
	return &{{serverName}}GinHttpRouterService{
		srv: srv,
		router: api,
	}
}
`

func routerTemplateReplace(serverName string) string {
	tmp := strings.ReplaceAll(routerTemplate, "{{serverName}}", serverName)
	return tmp
}
func routerFnTemplateReplace(serverName string, methods []*MethodInfo) string {
	var fns []string
	for _, method := range methods {
		var fnName = FnName(method.ServerIndex, method.ServerName, method.Method)
		if method.Path == "" {
			continue
		}
		if method.IsPost {
			fnTp := fmt.Sprintf(`router.ApiPostFn(s.router,"%s", %s(s.srv)),`, method.Path, fnName)
			fns = append(fns, fnTp)
		} else {
			fnTp := fmt.Sprintf(`router.ApiGetFn(s.router,"%s", %s(s.srv)),`, method.Path, fnName)
			fns = append(fns, fnTp)
		}
	}
	fnTmp := strings.Join(fns, "\n")
	tmp := strings.ReplaceAll(routerFnTemplate, "{{serverName}}", serverName)
	tmp = strings.ReplaceAll(tmp, "{{FN}}", fnTmp)
	return tmp
}
func registerTemplateReplace(serverName string) string {
	return strings.ReplaceAll(registerTemplate, "{{serverName}}", serverName)
}

const httpRequestTemplate = `
func (s *%s) %s(ctx *gin.Context)  (req *%s,err error) {
	 req = new(%s)
	 err = ctx.%s(&req)
	 return
`
const fnTemplates = `
func {{FN_NAME}}(srv {{serverName}}GinHttpService) router.Fn  {
	return func(ctx *gin.Context) (rs any, err error) {
		var in *{{inName}}
		if err = ctx.{{BIND}}(&in); err != nil {
			return
		}
		err = validate.Validate(in)
		if err != nil {
			return
		}
		rs, err = srv.{{method}}(ctx, in)
		return
	}
}
`

func FnName(serverIndex int, serverName, method string) string {
	return fmt.Sprintf("_%s_%s%d_GIN_HTTP_Handler", serverName, method, serverIndex)
}
func fnTemplatesReplace(isPost bool, serverIndex int, serverName, method, inName string) string {
	tmp := strings.ReplaceAll(fnTemplates, "{{FN_NAME}}", FnName(serverIndex, serverName, method))
	tmp = strings.ReplaceAll(tmp, "{{serverName}}", serverName)
	tmp = strings.ReplaceAll(tmp, "{{method}}", method)
	tmp = strings.ReplaceAll(tmp, "{{inName}}", inName)
	if isPost {
		//ShouldBind post  BindJSON  get
		tmp = strings.ReplaceAll(tmp, "{{BIND}}", "ShouldBind")
	} else {
		tmp = strings.ReplaceAll(tmp, "{{BIND}}", "BindJSON")
	}
	return tmp
}
