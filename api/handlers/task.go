package handlers

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
	"ryg-api-gateway/model"
	"strconv"
	"time"
)
import pb "ryg-api-gateway/gen_proto/task_service"

const (
	dateMonthYearFormat = "02-01-2006"
)

// CreateTask godoc
// @Security BearerAuth
// @Summary Create a new task
// @Description Create a new task
// @Tags Task
// @Accept json
// @Produce json
// @Param challenge_id path int true "Challenge ID"
// @Param task body model.CreateTaskRequest true "Task information"
// @Success 201 {object} pb.Task
// @Router /challenges/{challenge_id}/tasks [post]
func (cm *RpcClientManager) CreateTask(ctx *gin.Context) {
	req := model.CreateTaskRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	challengeId, err := strconv.Atoi(ctx.Param("challenge_id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid challenge ID"})
		return
	}

	res, err := cm.Task.CreateTask(ctx, &pb.CreateTaskRequest{
		Title:       req.Title,
		Description: req.Description,
		ChallengeId: int64(challengeId),
		UserId:      int64(ctx.GetInt("user_id")),
	})

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, res)
}

// CreateTasks godoc
// @Security BearerAuth
// @Summary Create multiple tasks
// @Description Create multiple tasks
// @Tags Task
// @Accept json
// @Produce json
// @Param challenge_id path int true "Challenge ID"
// @Param tasks body []model.CreateTaskRequest true "Tasks information"
// @Success 201 {object} pb.TaskList
// @Router /challenges/{challenge_id}/tasks/bulk [post]
func (cm *RpcClientManager) CreateTasks(ctx *gin.Context) {
	req := make([]model.CreateTaskRequest, 0)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	challengeId, err := strconv.Atoi(ctx.Param("challenge_id"))
	userId := int64(ctx.GetInt("user_id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid challenge ID"})
		return
	}

	tasks := make([]*pb.CreateTaskRequest, 0)

	for _, r := range req {
		tasks = append(tasks, &pb.CreateTaskRequest{
			Title:       r.Title,
			Description: r.Description,
			ChallengeId: int64(challengeId),
			UserId:      userId,
		})
	}

	res, err := cm.Task.CreateTasks(ctx, &pb.CreateTasksRequest{
		TaskRequests: tasks,
	})

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, res)
}

// GetTask godoc
// @Security BearerAuth
// @Summary Get a task by ID
// @Description Get a task by ID
// @Tags Task
// @Produce json
// @Param challenge_id path int true "Challenge ID"
// @Param task_id path int true "Task ID"
// @Success 200 {object} pb.Task
// @Router /challenges/{challenge_id}/tasks/{task_id} [get]
func (cm *RpcClientManager) GetTask(ctx *gin.Context) {
	req := pb.GetTaskRequest{}

	id, err := strconv.Atoi(ctx.Param("task_id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}

	challengeId, err := strconv.Atoi(ctx.Param("challenge_id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid challenge ID"})
		return
	}

	req.Id = int64(id)
	req.ChallengeId = int64(challengeId)
	req.UserId = int64(ctx.GetInt("user_id"))

	res, err := cm.Task.GetTaskById(ctx, &req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve task"})
		return
	}

	ctx.JSON(200, res)
}

// GetTasks godoc
// @Security BearerAuth
// @Summary Get all tasks by challenge ID
// @Description Get all tasks by challenge ID
// @Tags Task
// @Produce json
// @Param challenge_id path int true "Challenge ID"
// @Param date query string false "Date"
// @Success 200 {object} pb.TaskList
// @Success 202 {object} pb.TaskWithStatusList
// @Router /challenges/{challenge_id}/tasks [get]
func (cm *RpcClientManager) GetTasks(ctx *gin.Context) {
	challengeId, err := strconv.Atoi(ctx.Param("challenge_id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid challenge ID"})
		return
	}

	date := ctx.Query("date")

	if date == "" {
		req := pb.GetTasksByChallengeIdRequest{}
		req.ChallengeId = int64(challengeId)
		req.UserId = int64(ctx.GetInt("user_id"))

		res, err := cm.Task.GetTasksByChallengeId(ctx, &req)

		if err != nil {
			ctx.JSON(500, gin.H{"error": "Failed to retrieve tasks"})
			return
		}

		ctx.JSON(200, res)
	} else {
		dateAsTime, err := time.Parse(dateMonthYearFormat, date)

		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid date format"})
			return
		}

		req := pb.GetTaskByChallengeIdAndDateRequest{}
		req.ChallengeId = int64(challengeId)
		req.UserId = int64(ctx.GetInt("user_id"))
		req.Date = timestamppb.New(dateAsTime)

		res, err := cm.Task.GetTasksByChallengeIdAndDate(ctx, &req)

		if err != nil {
			ctx.JSON(500, gin.H{"error": "Failed to retrieve tasks"})
			return
		}

		ctx.JSON(202, res)
	}
}

// UpdateTask godoc
// @Security BearerAuth
// @Summary Update a task by ID
// @Description Update a task by ID
// @Tags Task
// @Accept json
// @Produce json
// @Param challenge_id path int true "Challenge ID"
// @Param task_id path int true "Task ID"
// @Param task body model.CreateTaskRequest true "Task information"
// @Success 200 {object} pb.Task
// @Router /challenges/{challenge_id}/tasks/{task_id} [put]
func (cm *RpcClientManager) UpdateTask(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("task_id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}

	challengeId, err := strconv.Atoi(ctx.Param("challenge_id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid challenge ID"})
		return
	}

	req := model.CreateTaskRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := cm.Task.UpdateTask(ctx, &pb.UpdateTaskRequest{
		Id:          int64(id),
		Title:       req.Title,
		Description: req.Description,
		UserId:      int64(ctx.GetInt("user_id")),
		ChallengeId: int64(challengeId),
	})

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// DeleteTask godoc
// @Security BearerAuth
// @Summary Delete a task by ID
// @Description Delete a task by ID
// @Tags Task
// @Param challenge_id path int true "Challenge ID"
// @Param task_id path int true "Task ID"
// @Success 204
// @Router /challenges/{challenge_id}/tasks/{task_id} [delete]
func (cm *RpcClientManager) DeleteTask(ctx *gin.Context) {
	req := pb.DeleteTaskRequest{}

	id, err := strconv.Atoi(ctx.Param("task_id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}

	req.Id = int64(id)
	req.UserId = int64(ctx.GetInt("user_id"))

	challengeId, err := strconv.Atoi(ctx.Param("challenge_id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid challenge ID"})
		return
	}

	req.ChallengeId = int64(challengeId)

	res, err := cm.Task.DeleteTask(ctx, &req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(204, res)
}

// UpdateTaskStatus godoc
// @Security BearerAuth
// @Summary Update a task status by ID
// @Description Update a task status by ID
// @Tags Task
// @Accept json
// @Produce json
// @Param challenge_id path int true "Challenge ID"
// @Param task_id path int true "Task ID"
// @Param status body model.UpdateTaskStatusRequest true "Task status information"
// @Success 200 {object} pb.TaskWithStatus
// @Router /challenges/{challenge_id}/tasks/{task_id}/status [put]
func (cm *RpcClientManager) UpdateTaskStatus(ctx *gin.Context) {
	taskId, err := strconv.Atoi(ctx.Param("task_id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}

	challengeId, err := strconv.Atoi(ctx.Param("challenge_id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid challenge ID"})
		return
	}

	req := model.UpdateTaskStatusRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	format := "02-01-2006"
	date, err := time.Parse(format, req.Date)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid date format"})
		return
	}

	res, err := cm.Task.UpdateTaskStatus(ctx, &pb.UpdateTaskStatusRequest{
		TaskId:      int64(taskId),
		Status:      req.Status,
		UserId:      int64(ctx.GetInt("user_id")),
		Date:        timestamppb.New(date),
		ChallengeId: int64(challengeId),
	})

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}
