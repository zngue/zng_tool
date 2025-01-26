package main

import (
	"google.golang.org/protobuf/compiler/protogen"
)

type Service struct {
	gen           *protogen.Plugin
	generatedFile *protogen.GeneratedFile
	messageMap    map[string]*protogen.Message
}

func (s *Service) Files() (err error) {
	for _, file := range s.gen.Files {
		if !file.Generate {
			continue
		}
		fileName := file.GeneratedFilenamePrefix + ".gin_service.pb.go"
		s.generatedFile = s.gen.NewGeneratedFile(fileName, "service")
		if err = Services(s.gen); err != nil {
			return
		}
		var messageMap = make(map[string]*protogen.Message)
		for _, message := range file.Messages {
			messageMap[message.GoIdent.GoName] = message
		}
		s.messageMap = messageMap

	}
	return
}
