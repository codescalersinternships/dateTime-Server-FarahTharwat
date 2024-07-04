package pkg

import "github.com/gin-gonic/gin"

type App struct {
	router *gin.Engine
}

func NewApp() *App {
	router := gin.New()
	router.GET("/dateTime", DateTimeEndpointGin(dateTimeNow))
	return &App{router: router}
}

func (app *App) Run(addr string) error {
	return app.router.Run(addr)
}

func (app *App) GetRouter() *gin.Engine {
	return app.router
}
