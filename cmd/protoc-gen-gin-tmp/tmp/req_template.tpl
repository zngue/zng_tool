{{- MessageFileType .Message  -}}
{{- range .Message.Fields -}}
	{{- if AutoNormal . -}}
		{{- $itemVal := AutoFiled . }}
		if req.{{$itemVal.GoName}}!=nil{
			{{ $itemVal.Name }}=&biz.{{ $itemVal.MessageType }}{
        	   {{AutoFileStruct   $itemVal.MessageType  $itemVal.NormalName }}
        	}
		}
	{{- end -}}
	{{- if AutoRepeated . -}}
		{{- $itemVal := AutoFiled . }}
		if len(req.{{$itemVal.GoName}})>0{
			for _,val:=range req.{{$itemVal.GoName}}{
				{{ $itemVal.Name }}=append({{ $itemVal.Name }},&biz.{{ $itemVal.MessageType }}{
					{{AutoFileStruct   $itemVal.MessageType  "val" }}
				})
			}
		}
	{{- end -}}
{{- end }}
var reqData = &biz.{{ RequestName .Message}}{
	{{RequestStruct .Message}}
}