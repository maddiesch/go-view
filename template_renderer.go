package view

import (
	"html/template"
	"io"
	"io/fs"
	"strings"
)

type TemplateRenderer struct {
	fs       fs.FS
	Layout   TemplateLayout
	ViewPath string
}

type TemplateLayout struct {
	Glob []string
	Name string
}

func NewTemplateRenderer(fs fs.FS) *TemplateRenderer {
	return &TemplateRenderer{
		fs:       fs,
		ViewPath: "{{name}}.template",
		Layout: TemplateLayout{
			Glob: []string{"layout/*.template"},
			Name: "layout",
		},
	}
}

type RenderContext struct {
	LayoutName   string
	TemplateFile string
	Data         any
}

func (r *TemplateRenderer) RenderWithName(w io.Writer, name string, data any, include ...string) error {
	t, err := template.ParseFS(r.fs, include...)
	if err != nil {
		return err
	}

	return t.ExecuteTemplate(w, name, RenderContext{
		LayoutName:   "",
		TemplateFile: "",
		Data:         data,
	})
}

func (r *TemplateRenderer) RenderWithLayout(w io.Writer, name string, data any) error {
	name = r.createTemplateName(name)
	patterns := append(r.Layout.Glob, name)
	t, err := template.ParseFS(r.fs, patterns...)
	if err != nil {
		return err
	}

	return t.ExecuteTemplate(w, r.Layout.Name, RenderContext{
		TemplateFile: name,
		LayoutName:   r.Layout.Name,
		Data:         data,
	})
}

func (r *TemplateRenderer) Render(w io.Writer, name string, data any) error {
	name = r.createTemplateName(name)
	content, err := fs.ReadFile(r.fs, name)
	if err != nil {
		return err
	}
	t, err := template.New(name).Parse(string(content))
	if err != nil {
		return err
	}

	return t.Execute(w, RenderContext{
		TemplateFile: name,
		LayoutName:   "",
		Data:         data,
	})
}

func (r *TemplateRenderer) createTemplateName(name string) string {
	return strings.ReplaceAll(r.ViewPath, "{{name}}", name)
}
