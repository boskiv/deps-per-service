package app

import (
	"deps_per_service/internal/config"
	"deps_per_service/internal/core"
	"deps_per_service/internal/handler"
	"deps_per_service/internal/repository"
	"deps_per_service/internal/service"
	transport "deps_per_service/internal/transport/http"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type App struct {
	services *service.Services
	repos    *repository.Repositories
	handlers *handler.Handlers
}

func NewApp() *App {
	println(config.Get().DB.Host)
	db, err := gorm.Open(postgres.Open(config.Get().DB.DSN()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	err = db.AutoMigrate(&core.User{}, &core.Profile{}, &core.Order{}, &core.Bill{})
	if err != nil {
		log.Fatal(err)
	}

	appRepositories := repository.NewRepositories(db)
	appServices := service.NewServices(appRepositories)
	appHandlers := handler.NewHandlers(appServices)
	return &App{
		services: appServices,
		repos:    appRepositories,
		handlers: appHandlers,
	}
}

func (a *App) Run() {
	fiberApp := transport.NewFiberApp(a.handlers)
	fiberApp.Serve()
}
