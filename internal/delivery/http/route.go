package http

import (
	"vinbigdata/config"
	"vinbigdata/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *Handler) InitRoute() {

	docs.SwaggerInfo.Title = "API Documentations"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.Domain
	docs.SwaggerInfo.Schemes = []string{"http"}

	// documentation public
	h.ginApp.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// telecom api
	mobile := h.ginApp.Group("/mobile")
	mobile.POST("/:user_name/call", h.telecomController.SaveCall)
	mobile.GET("/:user_name/billing", h.telecomController.GetBill)
}
