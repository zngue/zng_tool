{{ .VarContent }}
{{ range .DataItems -}}
    {{- $item := . -}}
    {{- if $item.IsRepeated -}}
    if len(reply.{{$item.GoName}}) > 0 {
       for _, val := range reply.{{$item.GoName}} {
          {{$item.LowerGoName}}Val= append({{$item.LowerGoName}}Val, s.ChangeItem(val))
       }
    }
    {{- else -}}
    if reply.{{$item.GoName}} != nil {
       {{$item.LowerGoName}}Val = s.ChangeItem(reply.{{$item.GoName}})
    }
    {{ end }}
{{- end }}
{{ range .OtherItems -}}
    {{- $item := . -}}
    {{- if $item.IsRepeated -}}
        if len(reply.{{$item.GoName}})>0{
           for _, val := range reply.{{$item.GoName}} {
                {{$item.LowerGoName}}Val= append({{$item.LowerGoName}}Val, &{{$item.GoKind}}{
                    {{ StructContent $item.MessageType "val" }}
                })
            }
        }
    {{- else -}}
		 if reply.{{$item.GoName}} != nil {
            {{$item.LowerGoName}}Val = &{{$item.GoKind}}{
                {{ StructContent $item.MessageType $item.Key }}
            }
         }
    {{ end }}
{{- end }}
rs = &{{ .RelyMessage }}{
    {{ .MessageContent }}
}

