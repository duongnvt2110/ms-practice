package event

import (
	"errors"
	apperror "ms-practice/event-service/pkg/utils/app_error"
	"ms-practice/event-service/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h Handler) ListEvents(c *gin.Context) {
	ctx := c.Request.Context()
	status := c.Query("status")
	events, err := h.eventUC.ListEvents(ctx, status)
	if err != nil {
		response.Error(c, apperror.ErrInternalServer.Wrap(err))
		return
	}
	response.Success(c, events)
}

func (h Handler) GetEvent(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, apperror.ErrBadRequest.Wrap(err))
		return
	}

	event, err := h.eventUC.GetEvent(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, apperror.ErrNotFound.Wrap(err))
			return
		}
		response.Error(c, apperror.ErrInternalServer.Wrap(err))
		return
	}
	response.Success(c, event)
}

func (h Handler) CreateEvent(c *gin.Context) {
	ctx := c.Request.Context()
	var req createEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, apperror.ErrBadRequest.Wrap(err))
		return
	}

	model, err := req.toModel()
	if err != nil {
		response.Error(c, apperror.ErrBadRequest.Wrap(err))
		return
	}

	if err := h.eventUC.CreateEvent(ctx, model); err != nil {
		response.Error(c, apperror.ErrInternalServer.Wrap(err))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": model.Id})
}
