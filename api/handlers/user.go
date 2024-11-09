package handlers

import (
	"github.com/gin-gonic/gin"
	pb "ryg-api-gateway/gen_proto/user_service"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body pb.CreateUserRequest true "User information"
// @Success 200 {object} pb.User
// @Router /user/register [post]
func (h *RpcClientManager) Register(ctx *gin.Context) {
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

	ctx.JSON(200, res)
}
