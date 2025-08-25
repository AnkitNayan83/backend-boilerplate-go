package api

import (
	"net/http"

	db "github.com/AnkitNayan83/backend-boilerplate-go/db/sqlc"
	"github.com/AnkitNayan83/backend-boilerplate-go/util"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Name     string `json:"name" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type createUserResponse struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := createUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	ctx.JSON(http.StatusCreated, res)
}
