package api

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	db "github.com/AnkitNayan83/backend-boilerplate-go/db/sqlc"
	"github.com/AnkitNayan83/backend-boilerplate-go/util"
	"github.com/gin-gonic/gin"
)

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}

	_, err := server.store.GetUserByEmail(ctx, req.Email)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
			return
		}
	} else {
		ctx.JSON(http.StatusConflict, NewErrorResponse("user already exists"))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
		return
	}

	arg := db.CreateUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
		return
	}

	res := createUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	ctx.JSON(http.StatusCreated, NewSuccessResponse(res))
}

func (server *Server) getUserById(c *gin.Context) {
	userId := c.Param("id")

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid user id"))
		return
	}

	user, err := server.store.GetUserByID(c, int32(userIdInt))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, NewErrorResponse("user not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
		return
	}

	res := createUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, NewSuccessResponse(res))
}
