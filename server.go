package main

import (
	"github.com/Sosial-Media-App/sosialta/config"
	cDelivery "github.com/Sosial-Media-App/sosialta/features/contents/delivery"
	cRepo "github.com/Sosial-Media-App/sosialta/features/contents/repository"
	cServices "github.com/Sosial-Media-App/sosialta/features/contents/services"
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
	cRepo := cRepo.New(db)
	database.MigrateDB(db)
	uServices := uServices.New(uRepo)
	cServices := cServices.New(cRepo)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	uDelivery.New(e, uServices)
	cDelivery.New(e, cServices)

	e.Logger.Fatal(e.Start(":3000"))
}
