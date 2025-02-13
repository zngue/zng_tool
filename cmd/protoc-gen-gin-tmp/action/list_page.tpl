var (
	    rs    []*db.{{.SvrType}}
	    where = make(map[string]any)
	    total int64
	)
	{{ ListWhereOperator .RequestMessage }}
	rs, total, err = dbConn.ListPage(&data.ListRequest{
	    Page: &page.Page{
	        Page:     int(req.Page),
	        PageSize: int(req.PageSize),
	    },
	    Where: where,
	    Order: []string{"id desc"},
	})
	count = int32(total)
	{{- if requestName .ReplyMessage }}
    if len(rs) > 0 {
        for _, val := range rs {
            {{requestName .ReplyMessage}} = append({{requestName .ReplyMessage}}, {{LowerIndex}}.ChangeItem(val))
        }
    }
    {{-  end }}