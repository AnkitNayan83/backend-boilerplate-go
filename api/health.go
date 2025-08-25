package api

import "github.com/gin-gonic/gin"

func (server *Server) healthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "OK",
	})
}
