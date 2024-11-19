package model

type CreateChallengeRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Days        int32  `json:"days"`
}

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskStatusRequest struct {
	Date   string `json:"date"`
	Status string `json:"status"`
}
