var (
	    rs    []*db.{{.SvrType}}
	    where = make(map[string]any)
	)
	{{ ListWhereOperator .RequestMessage }}
	rs, err = dbConn.List(&data.ListRequest{
	    Page: &page.Page{
	        Page:     int(req.Page),
	        PageSize: int(req.PageSize),
	    },
	    Where: where,
	    Order: []string{"id desc"},
	})
	fmt.Println(rs)
