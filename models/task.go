package models

import "time"

type Task struct {
	ID                   string    `json:"id" db:"id"`
	Title                string    `json:"title" db:"title"`
	Description          string    `json:"description" db:"description"`
	Assignee             string    `json:"assignee" db:"assignee"`
	DueDate              time.Time `json:"dueDate" db:"due_date"`
	CreatedAt            time.Time `json:"createdAt" db:"created_at"`
	Priority             string    `json:"priority" db:"priority"`
	IsBlocked            bool      `json:"isBlocked" db:"is_blocked"`
	ColumnID             string    `json:"column" db:"column_id"`
	Project              string    `json:"project" db:"project"`
	CompletionPercentage int       `json:"completionPercentage" db:"completion_percentage"`
}
