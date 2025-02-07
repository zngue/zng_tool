	err=dbConn.Add(&db.{{.SvrType}}{
		{{ SetFiledNew .RequestMessage }}
	})