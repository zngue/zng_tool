package main

import (
	_ "embed"
	"fmt"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"net/http"
	"os"
	"strings"
)

var (
	omitemptyPrefix = "omitempty"
)

type MethodDesc struct {
	Name         string
	OriginalName string
	MethodIndex  int
	ServerIndex  int
	Request      string
	Reply        string
	Comment      string
	Path         string
	Method       string
}
type ServiceDesc struct {
	ServiceType string // Greeter
	ServiceName string // helloworld.Greeter
	Metadata    string // api/helloworld/helloworld.proto
	Comment     string
	Methods     []*MethodDesc
}

func buildHTTPRule(serverIndex, methodIndex int, service *protogen.Service, m *protogen.Method, rule *annotations.HttpRule) *MethodDesc {
	var (
		path   string
		method string
		body   string
	)
	switch pattern := rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		path = pattern.Get
		method = http.MethodGet
	case *annotations.HttpRule_Put:
		path = pattern.Put
		method = http.MethodPut
	case *annotations.HttpRule_Post:
		path = pattern.Post
		method = http.MethodPost
	case *annotations.HttpRule_Delete:
		path = pattern.Delete
		method = http.MethodDelete
	case *annotations.HttpRule_Patch:
		path = pattern.Patch
		method = http.MethodPatch
	case *annotations.HttpRule_Custom:
		path = pattern.Custom.Path
		method = pattern.Custom.Kind
	}
	if method == "" {
		method = http.MethodPost
	}
	if path == "" {
		path = fmt.Sprintf("%s/%s/%s", omitemptyPrefix, service.Desc.FullName(), m.Desc.Name())
	}
	body = rule.Body
	md := buildMethodDesc(serverIndex, methodIndex, m, method, path)
	if method == http.MethodGet || method == http.MethodDelete {
		if body != "" {
			_, _ = fmt.Fprintf(os.Stderr, "\u001B[31mWARN\u001B[m: %s %s body should not be declared.\n", method, path)
		}
	} else {
		if body == "" {
			_, _ = fmt.Fprintf(os.Stderr, "\u001B[31mWARN\u001B[m: %s %s does not declare a body.\n", method, path)
		}
	}
	return md
}
func buildMethodDesc(serverIndex, methodIndex int, m *protogen.Method, method, path string) *MethodDesc {
	comment := m.Comments.Leading.String() + m.Comments.Trailing.String()
	if comment != "" {
		comment = "// " + m.GoName + strings.TrimPrefix(strings.TrimSuffix(comment, "\n"), "//")
	}
	return &MethodDesc{
		Name:         m.GoName,
		OriginalName: string(m.Desc.Name()),
		MethodIndex:  methodIndex,
		ServerIndex:  serverIndex,
		Request:      m.Input.GoIdent.GoName,
		Reply:        m.Output.GoIdent.GoName,
		Comment:      m.Comments.Leading.String(),
		Path:         path,
		Method:       method,
	}
}
