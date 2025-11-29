package response

import (
	apperror "ms-practice/event-service/pkg/utils/app_error"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func Error(c *gin.Context, err apperror.AppError) {
	status := err.StatusCode
	if status == 0 {
		status = http.StatusInternalServerError
	}
	c.JSON(status, gin.H{"error": err.Message})
}
