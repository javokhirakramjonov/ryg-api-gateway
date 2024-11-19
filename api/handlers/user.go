package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	token "ryg-api-gateway/api/token"
	pb "ryg-api-gateway/gen_proto/user_service"
	"ryg-api-gateway/model"
	"strconv"
)

// RegisterUser godoc
// @Summary RegisterUser a new user
// @Description RegisterUser a new user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body pb.CreateUserRequest true "User information"
// @Success 201 {object} pb.User
// @Router /register [post]
func (cm *RpcClientManager) RegisterUser(ctx *gin.Context) {
	req := pb.CreateUserRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := cm.User.CreateUser(ctx, &req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, res)
}

// GetProfile godoc
// @Summary Get user profile
// @Description Get user profile
// @Tags User
// @Produce json
// @Success 200 {object} pb.User
// @Router /users [get]
// @Security BearerAuth
func (cm *RpcClientManager) GetProfile(ctx *gin.Context) {
	req := pb.GetUserRequest{}

	req.Id = ctx.GetInt64("user_id")

	// Use ctx.Request.Context() to pass a gRPC-compatible context
	res, err := cm.User.GetUserById(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve user profile"})
		return
	}

	ctx.JSON(200, res)
}

// UpdateUser godoc
// @Summary UpdateUser user profile
// @Description UpdateUser user profile
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.UpdateUserRequest true "User information"
// @Success 200 {object} pb.User
// @Router /users [put]
// @Security BearerAuth
func (cm *RpcClientManager) UpdateUser(ctx *gin.Context) {
	req := &model.UpdateUserRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := cm.User.UpdateUser(ctx, &pb.UpdateUserRequest{
		Id:       ctx.GetInt64("user_id"),
		FullName: req.FullName,
		Email:    req.Email,
	})

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// DeleteUser godoc
// @Summary DeleteUser user profile
// @Description DeleteUser user profile
// @Tags User
// @Success 204
// @Router /users [delete]
// @Security BearerAuth
func (cm *RpcClientManager) DeleteUser(ctx *gin.Context) {
	req := &pb.DeleteUserRequest{}

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	req.Id = int64(id)

	res, err := cm.User.DeleteUser(ctx, req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(204, res)
}

// Login godoc
// @Summary Login user
// @Description Login user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body model.LoginRequest true "User information"
// @Success 200 {object} model.LoginResponse
// @Router /login [post]
func (cm *RpcClientManager) Login(ctx *gin.Context) {
	req := &model.LoginRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := cm.User.GetUserForLogin(ctx, &pb.GetUserForLoginRequest{Email: req.Email})

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		ctx.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	tkn, err := token.GenerateJWT(user.Id, user.Role)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"token": tkn})
}
