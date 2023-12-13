package main

import (
	"log"
	"my_project/pkg/auth"
	"my_project/pkg/config"
	"my_project/pkg/order"
	"my_project/pkg/product"

	"github.com/gin-gonic/gin"
)

func main() {
	conf, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSrvc := *auth.RegisterRoutes(r, &conf)
	product.RegisterRoutes(r, &conf, &authSrvc)
	order.RegisterRoutes(r, &conf, &authSrvc)

	r.Run(conf.Port)

}
