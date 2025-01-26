package model

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"gitlab.cqrb.cn/shangyou_mic/protoc-gen-validate/templates"
	"path/filepath"
	"strings"
)

const (
	validatorName = "validator"
	langParam     = "lang"
	moduleParam   = "module"
)

type Module struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	// lang contains the selected language (one of 'cc', 'go', 'java').
	// It is initialized in ValidatorForLanguage.
	// If unset, it will be parsed as the 'lang' parameter.
	lang string
}

func Validator() pgs.Module { return &Module{ModuleBase: &pgs.ModuleBase{}} }

func ValidatorForLanguage(lang string) pgs.Module {
	return &Module{lang: lang, ModuleBase: &pgs.ModuleBase{}}
}

func (m *Module) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.ctx = pgsgo.InitContext(ctx.Parameters())
}

func (m *Module) Name() string { return validatorName }

func (m *Module) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	lang := m.lang
	module := m.Parameters().Str(moduleParam)
	tpls := templates.Template(m.Parameters())[lang]
	m.Assert(tpls != nil, "could not find templates for `lang`: ", lang)

	for _, f := range targets {
		m.Push(f.Name().String())

		for _, msg := range f.AllMessages() {
			m.CheckRules(msg)
		}

		for _, tpl := range tpls {
			out := templates.FilePathFor(tpl)(f, m.ctx, tpl)
			if out != nil {
				outPath := strings.TrimLeft(strings.ReplaceAll(filepath.ToSlash(out.String()), module, ""), "/")
				m.AddGeneratorTemplateFile(outPath, tpl, f)
			}
		}

		m.Pop()
	}

	return m.Artifacts()
}
