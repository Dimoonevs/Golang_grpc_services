package product

import (
	"log"
	auth "my_project/pkg/auth"
	conf "my_project/pkg/config"
	routes "my_project/pkg/product/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, config *conf.Config, authServ *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authServ)

	srvc := &ServiceClient{
		Client: InitServiceClient(config),
	}
	routes := r.Group("/product")
	routes.Use(a.AuthRequired)
	routes.POST("/", srvc.CreateProduct)
	routes.GET("/:id", srvc.FindOne)

}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}
func (svc *ServiceClient) DecreaseStock(ctx *gin.Context) {
	log.Println("func DecreaseStock")
}
