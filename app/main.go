package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/meroedu/meroedu/api_docs"
	"github.com/meroedu/meroedu/app/config"
	_courseHttpDelivery "github.com/meroedu/meroedu/app/course/delivery/http"
	"gopkg.in/alecthomas/kingpin.v2"

	_categoryHttpDelivery "github.com/meroedu/meroedu/app/category/delivery/http"
	_categoryRepo "github.com/meroedu/meroedu/app/category/repository/mysql"
	_categoryUcase "github.com/meroedu/meroedu/app/category/usecase"
	_courseHttpDeliveryMiddleware "github.com/meroedu/meroedu/app/course/delivery/http/middleware"
	_courseRepo "github.com/meroedu/meroedu/app/course/repository/mysql"
	_courseUcase "github.com/meroedu/meroedu/app/course/usecase"
	"github.com/meroedu/meroedu/app/infrastructure/datastore"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
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
	db := datastore.NewDB()
	// db.AutoMigrate(domain.Course{}, domain.Category{})
	e := echo.New()

	// Init HealthCheck
	e.GET("/", HealthCheck)

	// Init Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	middL := _courseHttpDeliveryMiddleware.InitMiddleware()
	e.Use(middL.CORS)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	// Courses
	courseRepositry := _courseRepo.InitMysqlRepository(db)
	_courseHttpDelivery.NewCourseHandler(e, _courseUcase.NewCourseUseCase(courseRepositry, timeoutContext))

	// Categories
	categoryRepository := _categoryRepo.InitMysqlRepository(db)
	_categoryHttpDelivery.NewCategroyHandler(e, _categoryUcase.NewCategoryUseCase(categoryRepository, timeoutContext))

	// Start HTTP Server
	go func() {
		if err := e.Start(viper.GetString("server.address")); err != nil {
			e.Logger.Info("Shutting down the server")
		}
	}()

	//Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags health
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"health": "UP",
	})
}
