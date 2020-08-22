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
	_courseHttpDeliveryMiddleware "github.com/meroedu/meroedu/app/course/delivery/http/middleware"
	_courseRepo "github.com/meroedu/meroedu/app/course/repository/mysql"
	_courseUcase "github.com/meroedu/meroedu/app/course/usecase"
	"github.com/meroedu/meroedu/app/infrastructure/datastore"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
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
	config.ReadConfig()
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

	courseRepositry := _courseRepo.InitMysqlRepository(db)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := _courseUcase.NewCourseUseCase(courseRepositry, timeoutContext)
	_courseHttpDelivery.NewCourseHandler(e, au)

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
