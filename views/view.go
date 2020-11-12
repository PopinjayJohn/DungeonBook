package views

import (
  "html/template"
  "net/http"
  "path/filepath"
)

var LayoutDir string = "views/layouts"

func NewView(layout string, files ...string) *View {
  files = append(layoutFiles(), files...)
  t, err := template.ParseFiles(files...)
  if err != nil {
    panic(err)
  }

  return &View{
    Template: t,
    Layout:   layout,
  }
}

type View struct {
  Template *template.Template
  Layout   string
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
  vd := ViewData{
    Navigation: navigation(),
    Flashes: flashes(),
    Data:    data,
  }
  return v.Template.ExecuteTemplate(w, v.Layout, vd)
}

func layoutFiles() []string {
  files, err := filepath.Glob(LayoutDir + "/*.gohtml")
  if err != nil {
    panic(err)
  }
  return files
}

type ViewData struct {
  Navigation map[string]string
  Flashes map[string]string
  Data    interface{}
}

func flashes() map[string]string {
    return map[string]string{
      "is-warning": "Payment method about to expire!",
    }
}

func navigation() map[string]string {
  return map[string]string {
    "/": "Index",
    "/static/": "Static",
  }
}
