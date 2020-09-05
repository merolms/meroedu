package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/meroedu/meroedu/api_docs"
	_categoryHttpDelivery "github.com/meroedu/meroedu/internal/category/delivery/http"
	_categoryRepo "github.com/meroedu/meroedu/internal/category/repository/mysql"
	_categoryUcase "github.com/meroedu/meroedu/internal/category/usecase"
	"github.com/meroedu/meroedu/internal/config"
	_courseHttpDelivery "github.com/meroedu/meroedu/internal/course/delivery/http"
	_courseHttpDeliveryMiddleware "github.com/meroedu/meroedu/internal/course/delivery/http/middleware"
	_courseRepo "github.com/meroedu/meroedu/internal/course/repository/mysql"
	_courseUcase "github.com/meroedu/meroedu/internal/course/usecase"
	_healthHttpDelivery "github.com/meroedu/meroedu/internal/health/delivery/http"
	datastore "github.com/meroedu/meroedu/pkg/database"
	"github.com/meroedu/meroedu/pkg/log"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	configPath = kingpin.Flag("config", "Location of config.yml").Default("./config.yml").String()
)

// @title Mero Edu API
// @version 0.1
// @description Mero Edu is a software application for the administration, documentation, tracking, reporting, automation and delivery of educational courses, training programs, or learning and development programs for school.

// @contact.name Mero Edu
// @contact.url https://meroedu.com

// @license.name MIT License
// @license.url https://github.com/meroedu/meroedu/blob/master/LICENSE
// @BasePath /
func main() {

	// Parse the CLI flags and load the config
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	// Load the config
	config.ReadConfig(*configPath)
	db, err := datastore.NewDB()
	if err != nil {
		log.Error(err)
	}
	defer func() {
		log.Info("Closing database connection")
		if err := db.Close(); err != nil {
			log.Error(err)
		}
	}()

	e := echo.New()

	// Init Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	middL := _courseHttpDeliveryMiddleware.InitMiddleware()
	e.Use(middL.CORS)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	// healthcheck
	_healthHttpDelivery.NewHealthHandler(e)
	// Courses
	courseRepositry := _courseRepo.Init(db)
	_courseHttpDelivery.NewCourseHandler(e, _courseUcase.NewCourseUseCase(courseRepositry, timeoutContext))

	// Categories
	categoryRepository := _categoryRepo.Init(db)
	_categoryHttpDelivery.NewCategroyHandler(e, _categoryUcase.NewCategoryUseCase(categoryRepository, timeoutContext))

	// Start HTTP Server
	go func() {
		if err := e.Start(viper.GetString("server.address")); err != nil {
			log.Info("Shutting down the server")
		}
	}()

	//Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
