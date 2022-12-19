package entity

import "time"

type Email struct {
	ID          int64     `json:"id"`
	SentDate    time.Time `json:"sent_date"`
	Address     string    `json:"address"`
	TaskID      int64     `json:"task_id"`
	EmailTypeID int64     `json:"email_type_id"`
}
