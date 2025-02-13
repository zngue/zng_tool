{{- range .Message.Fields }}
	{{ if AutoRepeated . }}
		{{- $itemVal := AutoFiled . }}
		var req{{ $itemVal.Name }} []*biz.{{ $itemVal.MessageType }}
	{{- end}}
	{{ if AutoNormal . }}
		{{- $itemVal := AutoFiled . }}
		var {{ $itemVal.Name }} *biz.{{ $itemVal.MessageType }}
	{{- end}}

{{- end}}