// Package app configures and runs application.
package app

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"go-docker-devcontainer-starter/config"
	v1 "go-docker-devcontainer-starter/internal/controller/http/v1"
	"go-docker-devcontainer-starter/internal/usecase"
	"go-docker-devcontainer-starter/internal/usecase/repo"
	"go-docker-devcontainer-starter/internal/usecase/webapi"
	"go-docker-devcontainer-starter/pkg/httpserver"
	"go-docker-devcontainer-starter/pkg/logger"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)
	val := url.Values{}
	if cfg.DB.ParseTime != "" {
		val.Add("parseTime", cfg.DB.ParseTime)
	}

	if cfg.DB.Loc != "" {
		val.Add("loc", cfg.DB.Loc)
	}

	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	db, err := sql.Open(`mysql`, dsn)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - sql.Open: %w", err))
	}
	defer db.Close()

	gorm, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - gorm.Open: %w", err))
	}

	// Use case
	translationUseCase := usecase.New(
		repo.New(gorm),
		webapi.New(),
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, translationUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
