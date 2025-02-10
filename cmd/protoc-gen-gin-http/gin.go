package main

import (
	_ "embed"
	"fmt"
	"github.com/zngue/zng_tool/third_party/google/annotations"
	"net/http"
)

var (
	omitemptyPrefix = "v1"
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

func RequestInfo(rule *annotations.HttpRule, serverName, methodName string) (path string, method string) {
	if rule != nil {
		if rule.Path != "" && rule.Method != annotations.Method_METHOD_UNKNOWN {
			path = rule.Path
			method = AutoMethod(rule)
			if method != "" {
				return
			}
		}
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
		if path != "" && method != "" {
			return
		}
	}

	if path == "" || method == "" {
		path = fmt.Sprintf("/%s/%s/%s", omitemptyPrefix, serverName, methodName)
		method = http.MethodGet
	}
	return
}

func AutoMethod(rule *annotations.HttpRule) (method string) {
	switch rule.Method {
	case annotations.Method_GET:
		method = http.MethodGet
	case annotations.Method_POST:
		method = http.MethodPost
	case annotations.Method_WEBSOCKET:
		method = "WEBSOCKET"
	case annotations.Method_EVENT_STREAM:
		method = "EVENT_STREAM"
	}
	return
}
