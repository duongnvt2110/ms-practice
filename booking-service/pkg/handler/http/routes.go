package http_handler

import (
	"ms-practice/booking-service/pkg/config"
	"ms-practice/booking-service/pkg/handler/http/booking"
	"ms-practice/booking-service/pkg/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config, usecases *usecase.Usecase) {
	bookingHandler := booking.NewBookingHandler(cfg, usecases.BookingUC)
	bookingGroup := r.Group("/v1/bookings")
	{
		bookingGroup.GET("", bookingHandler.GetBookings)
		bookingGroup.GET("/:id", bookingHandler.GetBooking)
		bookingGroup.POST("", bookingHandler.CreateBooking)
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
