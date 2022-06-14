package main

import "golangs.org/snippetbox/pkg/models"


// storage for all dynamic data whitch should pass to HTML patterns
type templateData struct {
	Snippet *models.Snippet
}