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
	fmt.Println(rs)
