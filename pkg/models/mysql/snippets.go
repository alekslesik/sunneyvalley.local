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
func (m *SnippetModel) Instert(title, content, expires string) (int, error) {

	// SQL request
	stmt := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`


	// use Exec for execute request
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// use for taking last ID created in snippets
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// method for returning snippet data by ID
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// return the most uses 10 snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
