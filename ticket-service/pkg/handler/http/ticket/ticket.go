package ticket

import (
	"errors"
	apperror "ms-practice/ticket-service/pkg/utils/app_error"
	"ms-practice/ticket-service/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h Handler) ListTickets(c *gin.Context) {
	ctx := c.Request.Context()
	var userID *int
	if param := c.Query("user_id"); param != "" {
		if id, err := strconv.Atoi(param); err == nil {
			userID = &id
		} else {
			response.Error(c, apperror.ErrBadRequest.Wrap(err))
			return
		}
	}

	tickets, err := h.ticketUC.ListTickets(ctx, userID)
	if err != nil {
		response.Error(c, apperror.ErrInternalServer.Wrap(err))
		return
	}
	response.Success(c, tickets)
}

func (h Handler) GetTicket(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, apperror.ErrBadRequest.Wrap(err))
		return
	}
	ticket, err := h.ticketUC.GetTicket(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, apperror.ErrNotFound.Wrap(err))
			return
		}
		response.Error(c, apperror.ErrInternalServer.Wrap(err))
		return
	}
	response.Success(c, ticket)
}

func (h Handler) CreateTicket(c *gin.Context) {
	ctx := c.Request.Context()
	var req createTicketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, apperror.ErrBadRequest.Wrap(err))
		return
	}

	model := req.toModel()
	if err := h.ticketUC.CreateTicket(ctx, model); err != nil {
		response.Error(c, apperror.ErrInternalServer.Wrap(err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": model.Id})
}
