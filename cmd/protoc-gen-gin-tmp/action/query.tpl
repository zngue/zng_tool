var where = map[string]any{
      {{ UpdateWhereOperatorMore .RequestMessage }}
    }
    var rs *db.{{.SvrType}}
    rs, err = dbConn.Content(&data.ContentRequest{
        Where: where,
    })
    {{- if requestName .ReplyMessage }}
    if rs != nil {
        {{ requestName .ReplyMessage }} = {{LowerIndex}}.ChangeItem(rs)
    }
    {{-  end }}