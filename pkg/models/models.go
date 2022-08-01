package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: matching entry not found")

type Equipment struct {
	ID      int
	Title   string
	Content string
	Image string
	Url string
	Created time.Time
}
