package main

import (
	"log"
	"local/web-service-gin/config"
	"local/web-service-gin/internal/api"
	"local/web-service-gin/internal/infra/mssql"
	"local/web-service-gin/internal/infra/redis"
	"local/web-service-gin/internal/repository"
	"local/web-service-gin/internal/service"
  "fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

  //remove following these are added for testing
  fmt.Printf("cfg: %+v\n", cfg.App.DB.ConnString())
//	fmt.Println("Connection String:", cfg.App.DB.ConnString())

	db, err := mssql.NewConnection(cfg.App.DB.ConnString())
	if err != nil {
		log.Fatalf("MSSQL error: %v", err)
	}

	rdb := redis.NewClient(cfg.Redis)

  _ = rdb
	// Wire up dependencies
	userRepo := repository.NewUserRepository(db)
	downloadSvc := service.NewDownloadService(userRepo)

	// Init Gin
	r := gin.Default()

	api.RegisterRoutes(r, downloadSvc)

	log.Printf("Server running on port %s...", cfg.Server.Port)
	r.Run(":" + cfg.Server.Port)
}
