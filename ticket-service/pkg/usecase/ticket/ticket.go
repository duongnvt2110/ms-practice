package ticket

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"ms-practice/ticket-service/pkg/models"
	ticketRepo "ms-practice/ticket-service/pkg/repositories/ticket"
)

const (
	defaultTicketStatus = "issued"
)

type Usecase interface {
	CreateTicket(ctx context.Context, ticket *models.Ticket) error
	GetTicket(ctx context.Context, id int) (*models.Ticket, error)
	ListTickets(ctx context.Context, userID *int) ([]models.Ticket, error)
}

type usecase struct {
	ticketRepo ticketRepo.Repository
}

func NewUsecase(ticketRepo ticketRepo.Repository) Usecase {
	return &usecase{ticketRepo: ticketRepo}
}

func (u *usecase) CreateTicket(ctx context.Context, ticket *models.Ticket) error {
	setTicketDefaults(ticket)
	return u.ticketRepo.Create(ctx, ticket)
}

func (u *usecase) GetTicket(ctx context.Context, id int) (*models.Ticket, error) {
	return u.ticketRepo.GetByID(ctx, id)
}

func (u *usecase) ListTickets(ctx context.Context, userID *int) ([]models.Ticket, error) {
	return u.ticketRepo.List(ctx, userID)
}

func setTicketDefaults(ticket *models.Ticket) {
	if ticket.Code == "" {
		ticket.Code = fmt.Sprintf("TCK-%s", randomHex(8))
	}
	if ticket.Status == "" {
		ticket.Status = defaultTicketStatus
	}
	if ticket.QRURL == "" {
		ticket.QRURL = fmt.Sprintf("https://tickets.local/%s", ticket.Code)
	}
}

func randomHex(length int) string {
	buf := make([]byte, length)
	if _, err := rand.Read(buf); err != nil {
		return "UNKNOWN"
	}
	return hex.EncodeToString(buf)
}
