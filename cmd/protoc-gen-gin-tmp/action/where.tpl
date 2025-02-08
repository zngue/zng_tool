{{ if IsString .FiledType }}
if req.{{.FiledName}} != "" {
	   {{ .Where }}
	}
{{- end }}
{{ if IsNumber .FiledType }}
if req.{{.FiledName}} > 0 {
	    {{ .Where }}
	}
{{- end }}
{{ if IsRepeated .FiledType }}
if len(req.{{.FiledName}}) > 0 {
		{{ .Where }}
	}
{{- end }}