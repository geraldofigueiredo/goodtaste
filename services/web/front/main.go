package main

import (
	"context"
	"github.com/geraldofigueiredo/goodtaste/pkg/logger"
	"github.com/geraldofigueiredo/goodtaste/services/web/front/handlers"
	"github.com/geraldofigueiredo/goodtaste/services/web/front/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	lg = logger.NewLogger()
)

func main() {
	e := echo.New()

	registerMiddlewares(e)

	orderSvc := service.NewOrderService()

	handlers.RegisterHandlers(e, orderSvc, lg)

	startServer(e)
}

func registerMiddlewares(e *echo.Echo) {
	configureLogging(e)
	e.Use(middleware.CORS())
}

func configureLogging(e *echo.Echo) {
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			lg.Info("request: ",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))
}

func startServer(e *echo.Echo) {
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
