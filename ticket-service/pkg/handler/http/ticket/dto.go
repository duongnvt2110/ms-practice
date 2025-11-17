package ticket

import "ms-practice/ticket-service/pkg/models"

type createTicketRequest struct {
	UserID       int    `json:"user_id" binding:"required"`
	BookingID    int    `json:"booking_id" binding:"required"`
	PaymentID    int    `json:"payment_id" binding:"required"`
	TicketTypeID int    `json:"ticket_type_id" binding:"required"`
	Code         string `json:"code"`
	QRURL        string `json:"qr_url"`
	Status       string `json:"status"`
}

func (r createTicketRequest) toModel() *models.Ticket {
	return &models.Ticket{
		UserId:       r.UserID,
		BookingId:    r.BookingID,
		PaymentId:    r.PaymentID,
		TicketTypeId: r.TicketTypeID,
		Code:         r.Code,
		QRURL:        r.QRURL,
		Status:       r.Status,
	}
}
