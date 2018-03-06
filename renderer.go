package templaterenderer

import (
    "html/template"
    "io"
)

type Renderer map[string]*template.Template

func (r Renderer) AddTemplate(tmplName string, filenames ...string) {
    r[tmplName] = template.Must(template.ParseFiles(filenames...))
}

func (r Renderer) Render(tmplName string, w io.Writer, data map[string]interface{}) error {
    return r[tmplName].Execute(w, data)
}
