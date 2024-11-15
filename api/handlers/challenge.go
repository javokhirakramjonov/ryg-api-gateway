package handlers

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "ryg-api-gateway/gen_proto/task_service"
	"ryg-api-gateway/model"
	"strconv"
	"time"
)

// CreateChallenge godoc
// @Security BearerAuth
// @Summary Create a new challenge
// @Description Create a new challenge
// @Tags challenge
// @Accept json
// @Produce json
// @Param challenge body model.CreateChallengeRequest true "Challenge information"
// @Success 201 {object} pb.Challenge
// @Router /challenges [post]
func (cm *RpcClientManager) CreateChallenge(ctx *gin.Context) {
	req := model.CreateChallengeRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := cm.Challenge.CreateChallenge(ctx, &pb.CreateChallengeRequest{
		Title:       req.Title,
		Description: req.Description,
		StartDate:   timestamppb.New(time.Unix(req.StartDate, 0)),
		EndDate:     timestamppb.New(time.Unix(req.EndDate, 0)),
	})

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, res)
}

// GetChallenge godoc
// @Security BearerAuth
// @Summary Get a challenge by ID
// @Description Get a challenge by ID
// @Tags challenge
// @Produce json
// @Param challenge_id path int true "Challenge ID"
// @Success 200 {object} pb.Challenge
// @Router /challenges/{challenge_id} [get]
func (cm *RpcClientManager) GetChallenge(ctx *gin.Context) {
	req := pb.GetChallengeRequest{}

	id, err := strconv.Atoi(ctx.Param("challenge_id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid challenge ID"})
		return
	}

	req.Id = int64(id)

	res, err := cm.Challenge.GetChallengeById(ctx, &req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// GetChallenges godoc
// @Security BearerAuth
// @Summary Get all challenges
// @Description Get all challenges
// @Tags challenge
// @Produce json
// @Success 200 {object} pb.ChallengeList
// @Router /challenges [get]
func (cm *RpcClientManager) GetChallenges(ctx *gin.Context) {
	req := pb.GetChallengesRequest{}

	req.UserId = int64(ctx.GetInt("user_id"))

	res, err := cm.Challenge.GetChallengesByUserId(ctx, &req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// UpdateChallenge godoc
// @Security BearerAuth
// @Summary Update a challenge by ID
// @Description Update a challenge by ID
// @Tags challenge
// @Accept json
// @Produce json
// @Param challenge_id path int true "Challenge ID"
// @Param challenge body model.CreateChallengeRequest true "Challenge information"
// @Success 200 {object} pb.Challenge
// @Router /challenges/{challenge_id} [put]
func (cm *RpcClientManager) UpdateChallenge(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("challenge_id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid challenge ID"})
		return
	}

	req := model.CreateChallengeRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := cm.Challenge.UpdateChallenge(ctx, &pb.UpdateChallengeRequest{
		Id:          int64(id),
		Title:       req.Title,
		Description: req.Description,
		StartDate:   timestamppb.New(time.Unix(req.StartDate, 0)),
		EndDate:     timestamppb.New(time.Unix(req.EndDate, 0)),
		UserId:      int64(ctx.GetInt("user_id")),
	})

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// DeleteChallenge godoc
// @Security BearerAuth
// @Summary Delete a challenge by ID
// @Description Delete a challenge by ID
// @Tags challenge
// @Param challenge_id path int true "Challenge ID"
// @Success 204
// @Router /challenges/{challenge_id} [delete]
func (cm *RpcClientManager) DeleteChallenge(ctx *gin.Context) {
	req := pb.DeleteChallengeRequest{}

	id, err := strconv.Atoi(ctx.Param("challenge_id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid challenge ID"})
		return
	}

	req.Id = int64(id)

	res, err := cm.Challenge.DeleteChallenge(ctx, &req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(204, res)
}
