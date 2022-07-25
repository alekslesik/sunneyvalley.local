package main

import (
	"html/template"
	"path/filepath"

	"github.com/alekslesik/snippetbox/pkg/models"
	"github.com/alekslesik/snippetbox/pkg/models/page"
)

// storage for all dynamic data whitch should pass to HTML patterns
type templateData struct {
	PageData *page.PageData
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {

	templateDir := "/root/go/src/github.com/alekslesik/snippetbox/template"
	templateIncludesDir := "/root/go/src/github.com/alekslesik/snippetbox/template/includes"
	componentsDir := "/root/go/src/github.com/alekslesik/snippetbox/pkg/components"

	// init new map keeping cache
	cache := map[string]*template.Template{}

	// use func Glob to get all filepathes slice with '.page.html' ext

	// get dir before index.html
	pages, err := filepath.Glob(filepath.Join(dir, "/*/index.html"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// get dir before index.html
		name := filepath.Dir(page)
		name = filepath.Base(name)

		// handle page
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// use ParseGlob to add all frame patterns (base.layout.html)
		ts, err = ts.ParseGlob(filepath.Join(templateDir, "*.layout.html"))
		if err != nil {
			return nil, err
		}

		// footer, header
		ts, err = ts.ParseGlob(filepath.Join(templateDir, "*.partial.html"))
		if err != nil {
			return nil, err
		}

		// all from /includes
		ts, err = ts.ParseGlob(filepath.Join(templateIncludesDir, "*.html"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(componentsDir, "*/*.html"))
		if err != nil {
			return nil, err
		}

		// add received patterns set to cache, using page name
		cache[name] = ts
	}

	return cache, nil
}
