package main

import (
	"github.com/gin-gonic/gin"
	"log"
	_ "main/docs"
	"main/router"
	"main/service"
)

// @title			Melting API
// @description	Backend system of Muxi_Melting
// @version		1.5
// @contact.name	@a48zhang & @Cg1028
// @contact.email	3557695455@qq.com 2194028175@qq.com
// @schemes		http
// @BasePath		/api/v1
func main() {
	service.Logger()
	service.InitService()
	err := router.Register(gin.Default()).RunTLS(service.ServerAddr, service.TLSCert, service.TLSKey)
	if err != nil {
		log.Fatal(err)
	}
}
