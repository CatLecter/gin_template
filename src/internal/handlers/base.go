package handlers

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "src/docs"
	"src/internal/services"
	"time"
)

type Handler struct{ service *services.Service }

func NewHandler(service *services.Service) *Handler { return &Handler{service: service} }

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     []string{"*"},
				AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
				AllowHeaders:     []string{"Origin", "Content-type", "Authorization"},
				AllowCredentials: true,
			},
		),
	)

	router.Use(
		gin.LoggerWithFormatter(
			func(param gin.LogFormatterParams) string {
				log := fmt.Sprintf(
					"{\"datetime\": \"%s\", \"client_ip\": \"%s\", \"method\": \"%s\", "+
						"\"path\": \"%s\", \"status_code\": \"%v\", \"latency\": \"%v ms\"}\n",
					param.TimeStamp.Format(time.DateTime),
					param.ClientIP,
					param.Method,
					param.Path,
					param.StatusCode,
					param.Latency.Microseconds(),
				)
				return log
			},
		),
	)

	docs.SwaggerInfo.Title = "Swagger API Docs - OpenAPI 2.0"
	docs.SwaggerInfo.Description = "API documentation for the Gin template"

	router.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)

	v1 := router.Group("/api/v1")
	user := v1.Group("/user")
	{
		user.POST("/", h.CreateUser)
		user.GET("/:uuid", h.GetUserByUUID)
	}
	return router
}
