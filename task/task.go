package task

type Task struct {
	ID          uint64 `json:"id"`
	Description string `json:"description"`
	Author      uint64 `json:"author"`
	Deadline    string `json:"deadline"`
	IsDone      bool   `json:"isDone"`
}
