package main

import (
	"os"
	routes "whatapp/routes"

	"github.com/gin-gonic/gin"
)

//
func setupRouter() *gin.Engine {
	r := gin.Default()
	routes.Init(r)

	return r
}

//
func setupConfigs() {
	os.Setenv("AuthDBNURL", "postgres://cogjgedlgavael:cf43a86f559ebdd296331ca10991a0bfc87dfcf1fb7c83d3407698719348a669@ec2-18-204-74-74.compute-1.amazonaws.com:5432/d7jnruc4m8g23q")
	os.Setenv("WEBSERVER_PORT", "8080")
}

//
func scopeSetupConfigs(router *gin.Engine) {

}

//
func main() {
	//call configuerations
	setupConfigs()

	r := setupRouter()

	scopeSetupConfigs(r)
	//set up Web host
	r.Run(":" + os.Getenv("WEBSERVER_PORT"))
}
