package model

type CreateChallengeRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   int64  `json:"start_date"`
	EndDate     int64  `json:"end_date"`
}

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
