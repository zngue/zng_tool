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

type HTTPRuleData struct {
	GeneratedFile *protogen.GeneratedFile
	ServerIndex,
	MethodIndex int
	Service *protogen.Service
	Method  *protogen.Method
	Rule    *annotations.HttpRule
}

func buildHTTPRule(req *HTTPRuleData) *MethodDesc {
	var (
		path   string
		method string
		body   string
	)
	switch pattern := req.Rule.Pattern.(type) {
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
		path = fmt.Sprintf("%s/%s/%s", omitemptyPrefix, req.Service.Desc.FullName(), req.Method.Desc.Name())
	}
	body = req.Rule.Body
	md := buildMethodDesc(&MethodDescReq{
		GeneratedFile: req.GeneratedFile,
		Method:        req.Method,
		MethodType:    method,
		Path:          path,
		MethodIndex:   req.MethodIndex,
		ServerIndex:   req.ServerIndex,
	})
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

type MethodDescReq struct {
	GeneratedFile *protogen.GeneratedFile
	Method        *protogen.Method
	MethodType,
	Path string
	MethodIndex,
	ServerIndex int
}

func buildMethodDesc(req *MethodDescReq) *MethodDesc {
	comment := req.Method.Comments.Leading.String() + req.Method.Comments.Trailing.String()
	if comment != "" {
		comment = "// " + req.Method.GoName + strings.TrimPrefix(strings.TrimSuffix(comment, "\n"), "//")
	}
	return &MethodDesc{
		Name:         req.Method.GoName,
		OriginalName: string(req.Method.Desc.Name()),
		MethodIndex:  req.MethodIndex,
		ServerIndex:  req.ServerIndex,
		Request:      req.GeneratedFile.QualifiedGoIdent(req.Method.Input.GoIdent),
		Reply:        req.GeneratedFile.QualifiedGoIdent(req.Method.Output.GoIdent),
		Comment:      comment,
		Path:         req.Path,
		Method:       req.MethodType,
	}
}
