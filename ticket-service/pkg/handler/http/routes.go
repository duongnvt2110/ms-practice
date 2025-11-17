package http_handler

import (
	"ms-practice/ticket-service/pkg/container"
	"ms-practice/ticket-service/pkg/handler/http/ticket"

	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine, c *container.Container) {
	ticketHandler := ticket.NewHandler(c.Cfg, c.Usecases)
	tickets := r.Group("/tickets")
	{
		tickets.GET("", ticketHandler.ListTickets)
		tickets.GET("/:id", ticketHandler.GetTicket)
		tickets.POST("", ticketHandler.CreateTicket)
	}
}
