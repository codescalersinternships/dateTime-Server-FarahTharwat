package main

import (
	"log"

	"github.com/codescalersinternships/dateTime-Server-FarahTharwat/pkg"
)

func main() {
	app := pkg.NewApp()
	//router.GET("/dateTime", pkg.DateTimeEndpointGin(time.Now().String))
	if err := app.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
