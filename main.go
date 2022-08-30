package main

import (
	"projeto/Americanas/controller"
	"projeto/Americanas/service"

	"github.com/gin-gonic/gin"
)

type Application struct {
	Route *gin.Engine
}

func NewApp() *Application {
	return &Application{Route: gin.Default()}
}

func (app *Application) DefineRoutes() *Application {
	controllers := app.Route.Group("/api")
	{
		controllers.GET("/list", controller.ListPlanets)
		controllers.POST("/create", controller.CreatePlanet)
		controllers.DELETE("/remove", controller.RemovePlanet)
	}
	return app
}

func (app *Application) Start(port string) *Application {
	service.SyncSwapi()
	app.Route.Run(":" + port)
	return app
}

func main() {
	//utils.GetConnectMongoDB()
	route := gin.Default()
	controllers := route.Group("/api")
	{
		controllers.GET("/list", controller.ListPlanets)
		controllers.POST("/create", controller.CreatePlanet)
		controllers.DELETE("/remove", controller.RemovePlanet)
	}
	route.Run(":8080")
	//NewApp().DefineRoutes().Start("8080")
}
