{{ .VarContent }}
{{ range .DataItems -}}
    {{- $item := . -}}
    {{- if $item.IsRepeated -}}
        if len({{$item.LowerGoName}})>0{
           for _, val := range {{$item.LowerGoName}} {
                {{$item.LowerGoName}}Val= append({{$item.LowerGoName}}Val, s.ChangeItem(val))
            }
        }
    {{- else -}}
        if {{$item.LowerGoName}} != nil {
            {{$item.LowerGoName}}Val = s.ChangeItem({{$item.LowerGoName}})
        }
    {{- end  -}}
{{- end }}
{{ range .OtherItems -}}
    {{- $item := . -}}
    {{- if $item.IsRepeated -}}
        if len({{$item.LowerGoName}})>0{
           for _, val := range {{$item.LowerGoName}} {
                {{$item.LowerGoName}}Val= append({{$item.LowerGoName}}Val, &{{$item.GoKind}}{
                    {{ StructContent $item.MessageType "val" }}
                })
            }
        }
    {{- else -}}
        if {{$item.LowerGoName}} != nil {
            {{$item.LowerGoName}}Val = &{{$item.GoKind}}{
                {{ StructContent $item.MessageType "{{$item.LowerGoName}}" }}
            }
        }
    {{- end -}}
{{- end }}
rs = &{{ .RelyMessage }}{
    {{ .MessageContent }}
}