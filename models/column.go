package models

type Column struct {
	ID    string `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
}
