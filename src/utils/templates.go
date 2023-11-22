package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadTemplates insere os templates na variavel templates
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// ExecuteTemplate renderiza uma página html na tela
func ExecuteTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
