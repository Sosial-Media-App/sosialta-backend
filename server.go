package main

import (
	"github.com/Sosial-Media-App/sosialta/config"
	uDelivery "github.com/Sosial-Media-App/sosialta/features/users/delivery"
	uRepo "github.com/Sosial-Media-App/sosialta/features/users/repository"
	uServices "github.com/Sosial-Media-App/sosialta/features/users/services"
	"github.com/Sosial-Media-App/sosialta/utils/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := database.InitDB(cfg)
	uRepo := uRepo.New(db)
	database.MigrateDB(db)
	uServices := uServices.New(uRepo)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	uDelivery.New(e, uServices)

	e.Logger.Fatal(e.Start(":3000"))
}
