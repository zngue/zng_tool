package pkg

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"text/template"
)

type (
	RegisterFn func(tpl *template.Template, params pgs.Parameters)
	FilePathFn func(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath
)

func MakeTemplate(fn RegisterFn, params pgs.Parameters) *template.Template {
	tpl := template.New("go")
	fn(tpl, params)
	return tpl
}
func Register(tpl *template.Template, params pgs.Parameters) {
}
