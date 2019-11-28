package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"{{projectName}}/app"
	"{{projectName}}/app/logging"
	"{{projectName}}/router"
)

func init() {
	app.InitConfig()
	app.InitDB()
	logging.InitLog()
}

func main() {
	gin.SetMode(app.Config.Server.RunMode)
	routersInit := router.InitRouter()
	readTimeout := app.Config.Server.ReadTimeout
	writeTimeout := app.Config.Server.WriteTimeout
	endPoint := fmt.Sprintf(":%d", app.Config.Server.HttpPort)
	maxHeaderBytes := 1 << 20
	server := &http.Server{Addr: endPoint, Handler: routersInit, ReadTimeout: readTimeout, WriteTimeout: writeTimeout, MaxHeaderBytes: maxHeaderBytes}
	log.Printf("[info] start http server listening %s", endPoint)
	log.Fatalln(server.ListenAndServe())
}
