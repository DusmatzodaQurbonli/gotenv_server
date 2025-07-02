package main

import (
	"Gotenv/configs"
	"Gotenv/internal/routes"
	"Gotenv/internal/security"
	"Gotenv/internal/server"
	"Gotenv/pkg/db"
	"Gotenv/pkg/logger"
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Gotenv API
// @version 1.0.0

// @description API Server for Gotenv Application
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// цветной вывод
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	// загрузка конфигов
	if err := godotenv.Load(".env"); err != nil {
		if err = godotenv.Load("example.env"); err != nil {
			log.Fatalf("error loading .env file: %v", err)
		}
	}

	// конфиги
	var err error
	security.AppSettings, err = configs.ReadSettings()
	if err != nil {
		log.Fatalf("failed to read settings: %v", err)
	}
	security.SetConnDB(security.AppSettings)

	// логгер
	if err = logger.Init(); err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}

	// подключение к БД (2 попытки)
	if err = db.ConnectToDB(); err != nil {
		time.Sleep(10 * time.Second)
		if err = db.ConnectToDB(); err != nil {
			log.Fatalf("failed to connect to DB: %v", err)
		}
	}
	if err = db.Migrate(); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	// роутинг
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://localhost:3001",
			"http://localhost:3002",
			"http://localhost:3003",
			"http://localhost:3004",
			"http://localhost:3005",
			"http://localhost:3006",
			"http://localhost:3007",
			"http://localhost:3008",
			"http://localhost:3009",
			"http://localhost:3010",
			"https://dontenv.intelligent.tj"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	mainServer := new(server.Server)

	go func() {
		if err = mainServer.Run(security.AppSettings.AppParams.PortRun, routes.InitRoutes(router)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error while starting HTTPS Service: %s", err)
		}
	}()

	// ловим сигнал завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println(yellow("\nStart of service termination"))

	// Останавливаем всё с timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = mainServer.Shutdown(ctx); err != nil {
		log.Printf("HTTP shutdown error: %s", err)
	} else {
		fmt.Println(green("HTTP-service terminated successfully"))
	}

	if err = db.CloseDBConn(); err != nil {
		log.Printf("DB close error: %s", err)
	}

	fmt.Println(red("End of program completion"))
}
