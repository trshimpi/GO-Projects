package config

import (
	"html/template"
)

// Appconfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
