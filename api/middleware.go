package api

import (
	"net/http"
	"strings"

	"github.com/AnkitNayan83/backend-boilerplate-go/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "auth_payload_key"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, NewErrorResponse("authorization header is not provided"))
			return
		}

		fields := strings.Fields(authorizationHeader)

		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, NewErrorResponse("invalid authorization header format"))
			return
		}

		if strings.ToLower(fields[0]) != authorizationTypeBearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, NewErrorResponse("unsupported authorization type"))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, NewErrorResponse("invalid access token"))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
