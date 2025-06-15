package http_handler

import (
	"booking-service/pkg/config"
	"booking-service/pkg/handler/http/user"
	"booking-service/pkg/kafka"

	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine, cfg *config.Config, kafka kafka.KafkaClient) {
	bookingHandler := user.NewBookingHandler(cfg, kafka)
	bookingGroup := r.Group("bookings")
	{
		bookingGroup.GET("", bookingHandler.GetBookings)
		bookingGroup.GET("/:id", bookingHandler.GetBooking)
	}

	// r.GET("/users/{id}", userHandler.GetUser)
	// r.GET("/users", userHandler.GetUser)
	// r.GET("/test", func(w http.ResponseWriter, r *http.Request) {
	// 	time.Sleep(10 * time.Second)
	// 	log.Println("testGracefulShutdown job completed")
	// 	result, _ := json.Marshal(map[string]interface{}{"status": "completed"})
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(200)
	// 	w.Write(result)
	// })
}
