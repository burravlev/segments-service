package app

import (
	"github.com/burravlev/avito-tech-test/config"
	"github.com/burravlev/avito-tech-test/internal/controllers"
	"github.com/burravlev/avito-tech-test/internal/repositories"
	v1 "github.com/burravlev/avito-tech-test/internal/routes/api/v1"
	"github.com/burravlev/avito-tech-test/internal/services"
	"github.com/burravlev/avito-tech-test/pkg/database"
	"github.com/burravlev/avito-tech-test/pkg/logger"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//	@title			Avito A/B testing service
//	@version		1.0
//	@desccription	Microservice for sement users

//	@host	localhost:8080
// BasePath /

func Run(configPath string) {
	// configuration
	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("config error: %s", err)
	}
	logger.Init(cfg.Log.Level)
	log.Info("Initializing database...")
	db, err := database.Connect(&cfg.DB)
	if err != nil {
		log.Fatalf("app - Run - database.Connect: %s", err)
	}
	r := gin.Default()

	segmentRepository := repositories.SegmentRespository(db)
	segmentService := services.SegmentService(segmentRepository)
	segmentController := controllers.SegmentsController(segmentService)
	v1.SegmentRoutes(r, segmentController)
	v1.FileRoutes(r)
	r.Run()
}
