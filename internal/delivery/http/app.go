package http

import (
	"vinbigdata/internal/delivery/http/controller"
	"vinbigdata/internal/repository/postgres"
	"vinbigdata/internal/service"

	"gorm.io/gorm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitHttp(database *gorm.DB) *gin.Engine {

	r := gin.New()
	cf := cors.DefaultConfig()
	cf.AllowAllOrigins = true
	cf.AllowCredentials = true
	cf.AddAllowHeaders("X-Api-Version")
	cf.AddAllowHeaders("X-App-Version")
	cf.AddAllowHeaders("X-App-Platform")
	cf.AddAllowHeaders("X-App-Device-Id")
	cf.AddAllowHeaders("Authorization")
	r.Use(gin.Recovery())
	r.Use(cors.New(cf))

	// init repositories
	mobileRepo := postgres.NewMobileRepository(database)

	// init services
	mobileService := service.NewMobileService(mobileRepo)

	// init controller
	handler := NewHandler(
		WithGinEngine(r),
		WithTelecomController(controller.NewTelecomController(mobileService)),
	)

	// create route
	handler.InitRoute()

	return r
}
