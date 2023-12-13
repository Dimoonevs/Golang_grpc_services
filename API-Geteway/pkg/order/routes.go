package order

import (
	auth "my_project/pkg/auth"
	conf "my_project/pkg/config"
	routes "my_project/pkg/order/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *conf.Config, authSrvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSrvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/order")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateOrder)

}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
