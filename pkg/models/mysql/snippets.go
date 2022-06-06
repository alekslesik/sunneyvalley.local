package mysql

import (
	"database/sql"

	// import package models
	"golangs.org/snippetbox/pkg/models"
)

// define type whitch wrapping connecting pull sql.DB
type SnippetModel struct {
	DB *sql.DB
}

// method for create new snippet in data base
func (m *SnippetModel) Instert(title, content, expires string) (int, error){
	return 0, nil
}

// method for returning snippet data by ID
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// return the most uses 10 snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error)  {
	return nil, nil
}