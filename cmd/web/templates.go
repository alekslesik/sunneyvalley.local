package main

import (
	"golangs.org/snippetbox/pkg/models"
	"golangs.org/snippetbox/pkg/models/page"
)

// storage for all dynamic data whitch should pass to HTML patterns
type templateData struct {
	PageData *page.PageData
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
