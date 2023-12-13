package auth

import (
	"context"
	pb "my_project/pkg/auth/pb"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(scv *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{scv}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authoriztion := ctx.Request.Header.Get("authorization")

	if authoriztion == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token := strings.Split(authoriztion, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})
	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.UserId)

	ctx.Next()
}
