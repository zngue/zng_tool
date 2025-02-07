var where = map[string]any{
      {{ UpdateWhereOperatorMore .RequestMessage }}
    }
    var rs *db.{{.SvrType}}
    rs, err = dbConn.Content(&data.ContentRequest{
        Where: where,
    })
    fmt.Println(rs)