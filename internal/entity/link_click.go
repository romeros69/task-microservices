package entity

import "time"

type LinkClick struct {
	ID           int64     `json:"id"`
	ActionDate   time.Time `json:"action_date"`
	TaskID       int64     `json:"task_id"`
	ActionAuthor string    `json:"action_author"`
	ActionResult bool      `json:"action_result"`
}
