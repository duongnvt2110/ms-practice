package http_handler

import (
	"ms-practice/catalog-service/pkg/container"
	"ms-practice/catalog-service/pkg/handler/http/event"

	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine, c *container.Container) {
	eventHandler := event.NewHandler(c.Cfg, c.Usecases)
	events := r.Group("/events")
	{
		events.GET("", eventHandler.ListEvents)
		events.GET("/:id", eventHandler.GetEvent)
		events.POST("", eventHandler.CreateEvent)
	}
}
