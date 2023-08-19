package main

import (
	embedregistry "ebs_server/embed-registry"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	videoHandler := NewVideoHandler(embedregistry.NewEmbedRegistry())
	e.GET("/video", videoHandler.GetVideo)
	e.POST("/video", videoHandler.UpdateVideo)
	e.Logger.Fatal(e.Start(":3333"))
}
