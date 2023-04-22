package main

import (
	"github.com/gin-gonic/gin"
	"log"
	_ "main/docs"
	"main/router"
	"main/service"
)

//	@title			Melting API
//	@description	Backend system of Muxi_Melting
//	@description.markdown
//	@version		1.7
//	@contact.name	@a48zhang & @Cg1028
//	@contact.email	3557695455@qq.com 2194028175@qq.com
//	@schemes		https
//	@BasePath		/api/v1

func main() {
	service.Logger()
	service.InitService()
	errs := make(chan error, 1024)
	go func(ch chan error) {
		err := router.Register(gin.Default()).RunTLS(service.ServerAddr, service.TLSCert, service.TLSKey)
		if err != nil {
			errs <- err
			return
		}
	}(errs)

	go func(ch chan error) {
		err := router.WSHandlerRegister(gin.Default()).Run(":24769")
		if err != nil {
			errs <- err
			return
		}
	}(errs)

	if err, ok := <-errs; ok {
		log.Fatal(err)
	}
}
