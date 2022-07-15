package main

import (
	"github.com/alekslesik/snippetbox/pkg/models"
	"github.com/alekslesik/snippetbox/pkg/models/page"
)

// storage for all dynamic data whitch should pass to HTML patterns
type templateData struct {
	PageData *page.PageData
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
