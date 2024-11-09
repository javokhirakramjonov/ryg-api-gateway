package handlers

import (
	"github.com/gin-gonic/gin"
	pb "ryg-api-gateway/gen_proto/user_service"
	"strconv"
)

// RegisterUser godoc
// @Summary RegisterUser a new user
// @Description RegisterUser a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body pb.CreateUserRequest true "User information"
// @Success 201 {object} pb.User
// @Router /users/register [post]
func (h *RpcClientManager) RegisterUser(ctx *gin.Context) {
	req := pb.CreateUserRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.User.CreateUser(ctx, &req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, res)
}

// GetProfile godoc
// @Summary Get user profile
// @Description Get user profile
// @Tags user
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} pb.User
// @Router /users/{id} [get]
func (h *RpcClientManager) GetProfile(ctx *gin.Context) {
	req := pb.GetUserRequest{}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	req.Id = int32(id)

	// Use ctx.Request.Context() to pass a gRPC-compatible context
	res, err := h.User.GetUserById(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve user profile"})
		return
	}

	ctx.JSON(200, res)
}

// UpdateUser godoc
// @Summary UpdateUser user profile
// @Description UpdateUser user profile
// @Tags user
// @Accept json
// @Produce json
// @Param user body pb.UpdateUserRequest true "User information"
// @Success 200 {object} pb.User
// @Router /users [put]
func (h *RpcClientManager) UpdateUser(ctx *gin.Context) {
	req := &pb.UpdateUserRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.User.UpdateUser(ctx, req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// DeleteUser godoc
// @Summary DeleteUser user profile
// @Description DeleteUser user profile
// @Tags user
// @Param id path int true "User ID"
// @Success 204
// @Router /users/{id} [delete]
func (h *RpcClientManager) DeleteUser(ctx *gin.Context) {
	req := &pb.DeleteUserRequest{}

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	req.Id = int32(id)

	res, err := h.User.DeleteUser(ctx, req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(204, res)
}
