package entity

import "time"

type Task struct {
	ID           int64     `json:"id"`
	CreationDate time.Time `json:"creation_date"`
	Author       string    `json:"author"`
	StatusID     int64     `json:"status_id"`
}
