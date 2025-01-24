package main

import "testing"

func TestTem(t *testing.T) {

	var ara = &ServiceDesc{
		ServiceType: "User",
		ServiceName: "User",
		Metadata:    "",
		Comment:     "testUser",
		Methods: []*MethodDesc{
			{
				Name:         "List",
				OriginalName: "List",
				MethodIndex:  0,
				ServerIndex:  0,
				Request:      "UserRequest",
				Reply:        "UserReply",
				Comment:      "gin-pb",
				Path:         "/v1/user/list",
				Method:       "POST",
			},
			{
				Name:         "Info",
				OriginalName: "Info",
				MethodIndex:  1,
				ServerIndex:  0,
				Request:      "UserInfoRequest",
				Reply:        "UserInfoReply",
				Comment:      "gin-pb",
				Path:         "/v1/user/info",
				Method:       "GET",
			},
		},
	}
	execute := ara.execute()
	t.Log(execute)

}
