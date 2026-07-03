package httpx

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin/binding"
)

const tagName = "json"

var pathParamRe = regexp.MustCompile(`\{([^{}/]+?)\}`)

func Bind(r *http.Request, v any) (err error) {
	if err = bindPath(r, v); err != nil {
		return
	}
	if r.Method == http.MethodGet {
		return binding.MapFormWithTag(v, r.URL.Query(), tagName)
	}
	return binding.Default(r.Method, contentType(r)).Bind(r, v)
}

func BindQuery(r *http.Request, v any) error {
	return binding.MapFormWithTag(v, r.URL.Query(), tagName)
}

func BindJSON(r *http.Request, v any) error {
	return binding.JSON.Bind(r, v)
}

func bindPath(r *http.Request, v any) error {
	names := pathParamRe.FindAllStringSubmatch(r.Pattern, -1)
	if len(names) == 0 {
		return nil
	}
	values := make(map[string][]string, len(names))
	for _, m := range names {
		name := m[1]
		if val := r.PathValue(name); val != "" {
			values[name] = []string{val}
		}
	}
	if len(values) == 0 {
		return nil
	}
	return binding.MapFormWithTag(v, values, tagName)
}

func contentType(r *http.Request) string {
	ct := r.Header.Get("Content-Type")
	for i := 0; i < len(ct); i++ {
		if ct[i] == ';' {
			return ct[:i]
		}
	}
	return ct
}
