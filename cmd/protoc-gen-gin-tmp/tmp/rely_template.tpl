{{- range .DataItems -}}
    {{- $item := . -}}
    {{- if $item.IsRepeated -}}
        var {{$item.LowerGoName}}Val []*{{$item.GoKind}}
        if len({{$item.LowerGoName}})>0{
           for _, val := range {{$item.LowerGoName}} {
                {{$item.LowerGoName}}Val= append({{$item.LowerGoName}}Val, s.ChangeItem(val))
            }
        }
    {{- else -}}
        var {{$item.LowerGoName}}Val *{{$item.GoKind}}
        if {{$item.LowerGoName}} != nil {
            {{$item.LowerGoName}}Val = s.ChangeItem({{$item.LowerGoName}})
        }
    {{- end  -}}
{{- end }}
rs = &{{ .RelyMessage }}{
    {{ .MessageContent }}
}