package http_handler

import (
	"ms-practice/event-service/pkg/container"
	"ms-practice/event-service/pkg/handler/http/event"

	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine, c *container.Container) {
	eventHandler := event.NewHandler(c.Cfg, c.Usecases)
	events := r.Group("/v1/events")
	{
		events.GET("", eventHandler.ListEvents)
		events.GET("/:id", eventHandler.GetEvent)
		events.POST("", eventHandler.CreateEvent)
	}
}
