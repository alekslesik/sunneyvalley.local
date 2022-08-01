package mysql

import (
	"database/sql"
	"errors"

	"github.com/alekslesik/snippetbox/pkg/models"
)

// define type whitch wrapping connecting pull sql.DB
type EquipmentModel struct {
	DB *sql.DB
}

// method for create new snippet in data base
func (m *EquipmentModel) Instert(title, content, expires string) (int, error) {

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
func (m *EquipmentModel) Get(id int) (*models.Equipment, error) {

	// sql request for getting data of one snippet
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	// initialise pointer to new struct Snippet
	s := &models.Equipment{}

	// copy values from row to corresponding field s Snippet struct
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created,)
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
func (m *EquipmentModel) Dirs() ([]*models.Equipment, error) {
	smtp := `SELECT id, title, content, image, url, created FROM dirs`

	// get sql.Rows, included result of our request
	rows, err := m.DB.Query(smtp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// slice for storing models.Snippets objects
	var equipments []*models.Equipment

	// use for enumeration result "rows"
	for rows.Next() {
		// create a pointer to new struct
		s := &models.Equipment{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Image, &s.Url, &s.Created)
		if err != nil {
			return nil, err
		}

		// add struct to slice
		equipments = append(equipments, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return equipments, nil
}
