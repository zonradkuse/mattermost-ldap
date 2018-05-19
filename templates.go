package main

import (
	"github.com/pkg/errors"

	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, id string) {
	renderTemplateWithData(w, id, "")
}

// renderTemplate is a convenience helper for rendering templates.
func renderTemplateWithData(w http.ResponseWriter, id string, d interface{}) bool {
	if t, err := template.New(id).ParseFiles(LOCAL_PATH_TEMPLATES + id); err != nil {
		http.Error(w, errors.Wrap(err, "Could not render template").Error(), http.StatusInternalServerError)
		return false
	} else if err := t.Execute(w, d); err != nil {
		http.Error(w, errors.Wrap(err, "Could not render template").Error(), http.StatusInternalServerError)
		return false
	}
	return true
}
