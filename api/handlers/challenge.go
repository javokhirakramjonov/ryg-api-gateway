package handlers

import (
	"github.com/gin-gonic/gin"
	pb "ryg-api-gateway/gen_proto/task_service"
	"ryg-api-gateway/model"
	"strconv"
)

// CreateChallenge godoc
// @Security BearerAuth
// @Summary Create a new challenge
// @Description Create a new challenge
// @Tags Challenge
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
		Days:        req.Days,
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
// @Tags Challenge
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
// @Tags Challenge
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
// @Tags Challenge
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
		Days:        req.Days,
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
// @Tags Challenge
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

// StartChallenge godoc
// @Security BearerAuth
// @Summary Start a challenge by ID
// @Description Start a challenge by ID
// @Tags Challenge
// @Accept json
// @Produce json
// @Param challenge_id path int true "Challenge ID"
// @Success 200 {object} pb.Challenge
// @Router /challenges/{challenge_id}/start [post]
func (cm *RpcClientManager) StartChallenge(ctx *gin.Context) {
	req := pb.StartChallengeRequest{}

	challengeId, err := strconv.Atoi(ctx.Param("challenge_id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid challenge ID"})
		return
	}

	req.ChallengeId = int64(challengeId)
	req.UserId = int64(ctx.GetInt("user_id"))

	res, err := cm.Challenge.StartChallenge(ctx, &req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// FinishChallenge godoc
// @Security BearerAuth
// @Summary Finish a challenge by ID
// @Description Finish a challenge by ID
// @Tags Challenge
// @Accept json
// @Produce json
// @Param challenge_id path int true "Challenge ID"
// @Success 200 {object} pb.Challenge
// @Router /challenges/{challenge_id}/finish [post]
func (cm *RpcClientManager) FinishChallenge(ctx *gin.Context) {
	req := pb.FinishChallengeRequest{}

	challengeId, err := strconv.Atoi(ctx.Param("challenge_id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid challenge ID"})
		return
	}

	req.ChallengeId = int64(challengeId)
	req.UserId = int64(ctx.GetInt("user_id"))

	res, err := cm.Challenge.FinishChallenge(ctx, &req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// InviteUser godoc
// @Security BearerAuth
// @Summary Invite a user to a challenge
// @Description Invite a user to a challenge
// @Tags Challenge
// @Accept json
// @Produce json
// @Param challenge_id path int true "Challenge ID"
// @Param add_user_to_challenge_request body model.AddUserToChallengeRequest true "User ID"
// @Success 200
// @Router /challenges/{challenge_id}/invite [post]
func (cm *RpcClientManager) InviteUser(ctx *gin.Context) {
	req := model.AddUserToChallengeRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	challengeId, err := strconv.Atoi(ctx.Param("challenge_id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid challenge ID"})
		return
	}

	res, err := cm.Challenge.AddUserToChallenge(ctx, &pb.AddUserToChallengeRequest{
		ChallengeId: int64(challengeId),
		UserId:      ctx.GetInt64("user_id"),
		UserToAddId: req.UserId,
		Email:       ctx.GetString("email"),
	})

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// AcceptChallengeInvitation godoc
// @Security BearerAuth
// @Summary Accept a challenge invitation
// @Description Accept a challenge invitation
// @Tags Challenge
// @Accept json
// @Produce json
// @Success 200
// @Router /challenges/accept [post]
func (cm *RpcClientManager) AcceptChallengeInvitation(ctx *gin.Context) {
	invitationToken := ctx.Query("token")

	req := pb.SubscribeToChallengeRequest{}

	req.Token = invitationToken
	req.UserId = int64(ctx.GetInt("user_id"))

	res, err := cm.Challenge.SubscribeToChallenge(ctx, &req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// UnsubscribeFromChallenge godoc
// @Security BearerAuth
// @Summary Unsubscribe from a challenge
// @Description Unsubscribe from a challenge
// @Tags Challenge
// @Accept json
// @Produce json
// @Param challenge_id path int true "Challenge ID"
// @Success 204
// @Router /challenges/{challenge_id}/unsubscribe [delete]
func (cm *RpcClientManager) UnsubscribeFromChallenge(ctx *gin.Context) {
	req := pb.UnsubscribeFromChallengeRequest{}

	challengeId, err := strconv.Atoi(ctx.Param("challenge_id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid challenge ID"})
		return
	}

	req.ChallengeId = int64(challengeId)
	req.UserId = int64(ctx.GetInt("user_id"))

	res, err := cm.Challenge.UnsubscribeFromChallenge(ctx, &req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(204, res)
}
