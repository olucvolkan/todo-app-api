package requests

type Task struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	IsComplete bool   `json:"isComplete"`
}
