package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/AnkitNayan83/backend-boilerplate-go/util"
	"github.com/gin-gonic/gin"
)

func (server *Server) credentialLogin(c *gin.Context) {
	var req loginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnauthorized, NewErrorResponse(err.Error()))
		return
	}

	user, err := server.store.GetUserByEmail(c, req.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusUnauthorized, NewErrorResponse("user not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
		return
	}

	err = util.CheckPasswordHash(req.Password, user.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, NewErrorResponse("invalid credentials"))
		return
	}

	token, _, err := server.tokenMaker.CreateToken(user.ID, server.config.TokenDuration)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
		return
	}

	res := &loginUserResponse{
		AccessToken: token,
	}

	c.JSON(http.StatusOK, res)

}

func (server *Server) googleOAuth(c *gin.Context) {

}
