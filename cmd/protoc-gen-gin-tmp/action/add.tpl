	err=dbConn.Add(&db.{{.SvrType}}{
		{{ SetFiled .RequestMessage }}
	})