package v1

type taskRequestDTO struct {
	Author   string `json:"author"`
	StatusID int64  `json:"statusID"`
}
