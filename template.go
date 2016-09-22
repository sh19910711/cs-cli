package main

import (
	"io"
	"os"
	"text/template"
)

func render(w io.Writer, tmpl string, data interface{}) error {
	t := template.New("tmpl")
	template.Must(t.Parse(tmpl))
	if err := t.Execute(w, data); err != nil {
		return err
	}
	return nil
}

func renderErrorTemplate(tmpl string, data interface{}) error {
	return render(os.Stderr, tmpl, data)
}
