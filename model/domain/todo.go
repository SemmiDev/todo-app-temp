package domain

type ToDo struct {
	Id         string
	Task       string
	StartingAt string  // RFC850
	EndsAt     string  // RFC850
	Duration   float64 // in minutes
	Done       int     // 1 == true || 0 == false
	IsExpired  int     // 1 == true || 0 == false
}
