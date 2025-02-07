{{ if .IsNumber .FiledType }}
if req.Name != "" {
   where["name like ?"] = "%" + req.Name + "%"
}
{{- end }}
{{ if .IsNumber .FiledType }}
if req.Name >= 0 {
   where["name like ?"] = "%" + req.Name + "%"
}
{{- end }}