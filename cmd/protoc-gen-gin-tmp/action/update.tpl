var where = make(map[string]any)
{{- range .RequestMessage.Fields }}
  {{- if UpdateWhereOperator . }}
  {{ UpdateWhereOperator . }}
  {{- end }}
{{- end }}
	var updateData = map[string]any{
{{- range .RequestMessage.Fields }}
	{{- if UpdateOperator . }}
		{{ UpdateOperator . }}
	{{- end }}
{{- end }}
	}
	err=dbConn.Update(where, updateData)
