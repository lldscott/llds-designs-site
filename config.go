package config

import "html/template"

// sets configuration with formatting for go app does not import anything just uses present data

// AppConfig holds the application config sitewise
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
