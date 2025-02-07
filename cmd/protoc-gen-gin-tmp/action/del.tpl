var where = map[string]any{
      {{ UpdateWhereOperatorMore .RequestMessage }}
    }
    err = dbConn.Delete(where)