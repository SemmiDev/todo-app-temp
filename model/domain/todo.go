package domain

type ToDo struct {
	Id         string
	Task       string
	StartingAt string
	EndsAt     string
	Duration   float64
	IsExpired  int
	Done       int
}
