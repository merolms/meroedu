package main

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/meroedu/meroedu/app/config"
	_courseHttpDelivery "github.com/meroedu/meroedu/app/course/delivery/http"
	_courseHttpDeliveryMiddleware "github.com/meroedu/meroedu/app/course/delivery/http/middleware"
	_courseRepo "github.com/meroedu/meroedu/app/course/repository/mysql"
	_courseUcase "github.com/meroedu/meroedu/app/course/usecase"
	"github.com/meroedu/meroedu/app/infrastructure/datastore"
	"github.com/spf13/viper"
)

func main() {
	config.ReadConfig()
	db := datastore.NewDB()
	// db.AutoMigrate(domain.Course{}, domain.Category{})
	e := echo.New()

	// Swagger endpoint

	middL := _courseHttpDeliveryMiddleware.InitMiddleware()
	e.Use(middL.CORS)
	courseRepositry := _courseRepo.InitMysqlRepository(db)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := _courseUcase.NewCourseUseCase(courseRepositry, timeoutContext)
	_courseHttpDelivery.NewCourseHandler(e, au)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
