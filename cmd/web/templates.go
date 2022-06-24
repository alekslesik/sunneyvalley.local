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

func newTemplateCache(dir string) (map[string]*template.Template, error) {

	// initialise new map whitch will storing a cache
	cache := map[string]*template.Template{}

	// use filepath.Glob() to get slice of all filepaths with extension "*page.html
	pages, err := filepath.Glob(filepath.Join(dir, "*page.html"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// extract complited file name from full filepath
		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// use ParseGlob() for adding all frame patterns
		// in this case it is only file base.layout.html
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.html"))
		if err != nil {
			return nil, err
		}

		// use ParseGlob() for adding all additional patterns
		// in this case it is file footer.partial.html
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.html"))
		if err != nil {
			return nil, err
		}

		// add received patterns set in cache
		cache[name] = ts
	}

	return cache, nil
}