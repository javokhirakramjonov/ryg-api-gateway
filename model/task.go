package model

type CreateChallengeRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Days        int32  `json:"days"`
}

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	WeekDays    int32  `json:"week_days"`
}

type UpdateTaskStatusRequest struct {
	Date   string `json:"date"`
	Status string `json:"status"`
}

type AddUserToChallengeRequest struct {
	UserId int64 `json:"user_id"`
}
