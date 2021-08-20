package web

type CreateToDoRequest struct {
	Task       string `validate:"required,min=1,max=100" json:"task,omitempty"`
	StartingAt string `validate:"required" json:"starting_at,omitempty"`
	EndsAt     string `validate:"required" json:"ends_at,omitempty"`
}

type UpdateToDoRequest struct {
	Id         string `validate:"required,min=1,max=100" json:"id,omitempty"`
	Task       string `validate:"required,min=1,max=100" json:"task,omitempty"`
	StartingAt string `validate:"required" json:"starting_at,omitempty"`
	EndsAt     string `validate:"required" json:"ends_at,omitempty"`
}

type ToDoResponse struct {
	Id         string  `json:"id,omitempty"`
	Task       string  `json:"task,omitempty"`
	StartingAt string  `json:"starting_at,omitempty"`
	EndsAt     string  `json:"ends_at,omitempty"`
	Duration   float64 `json:"duration"`
	Expired    int     `json:"expired"`
	Done       int     `json:"done"`
}
