package mysql

import (
	"database/sql"
	"errors"

	
	"github.com/alekslesik/snippetbox/pkg/models"
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

	// sql request for getting data of one snippet
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	// initialise pointer to new struct Snippet
	s := &models.Snippet{}

	// copy values from row to corresponding field s Snippet struct
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	// return Snippet object
	return s, nil

}

// return the most uses 10 snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	smtp := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	// get sql.Rows, included result of our request
	rows, err := m.DB.Query(smtp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// slice for storing models.Snippets objects
	var snippets []*models.Snippet

	// use for enumeration result "rows"
	for rows.Next() {
		// create a pointer to new struct
		s := &models.Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		// add struct to slice
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
