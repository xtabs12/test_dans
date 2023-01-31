package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/xtabs12/test_dans/internal/controller"
	"github.com/xtabs12/test_dans/internal/repo"
	"github.com/xtabs12/test_dans/internal/ucase"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	errChan := make(chan error)

	echoServer := echo.New()
	echoServer.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowMethods: []string{"GET", "POST", "PUT", "OPTIONS", "DELETE"},
	}))
	echoServer.Use(middleware.RequestID())
	echoServer.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339_nano}][${id}],"remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))

	log.SetHeader(`[${time_rfc3339_nano}][${prefix}][${level}][${short_file}:${line}]`)
	serverGroup := echoServer.Group("/api/v1")
	//db
	db, dbConnErr := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if dbConnErr != nil {
		panic(dbConnErr)
	}

	//repo
	userRepo := repo.NewUser(db)

	// usecase
	authLogic := ucase.NewAuth(userRepo)
	jobLogic := ucase.NewJob()
	controller.NewJob(jobLogic, serverGroup)
	controller.NewAuth(authLogic, serverGroup)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errChan <- echoServer.Start(":8088")
	}()

	echoServer.Logger.Print("Starting ", "Golang Developer Test")
	err := <-errChan
	log.Error(err.Error())
}
