package v1

type taskRequestDTO struct {
	CreationDate string `json:"creationDate"`
	Author       string `json:"author"`
	StatusID     int64  `json:"statusID"`
}
