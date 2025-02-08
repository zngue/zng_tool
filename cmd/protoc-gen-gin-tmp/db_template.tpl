package db

type {{.ServerType}} struct {
	{{ StructType .Message }}
}
func ({{.ServerType}}) TableName() string {
	return "{{.LowerServerType}}"
}