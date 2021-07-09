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
	os.Setenv("AuthDBNURL", "postgresql://postgres:password:@localhost:5432/whatsapp")
	os.Setenv("KAFKAUSERNAME", "user")

}

//
func scopeSetupConfigs(router *gin.Engine) {

}

//
func main() {
	//call configuerations
	setupConfigs()
	PORT := "3000"
	r := setupRouter()

	// dynamicaaly  scope router configuerations
	scopeSetupConfigs(r)
	//set up local host
	r.Run(":" + os.Getenv(PORT) + PORT)
}

// "postgresql://postgres:password:@localhost:5432/whatsapp"
