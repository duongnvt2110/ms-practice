package event

import (
	"errors"
	"ms-practice/catalog-service/pkg/models"
	"time"
)

type createEventRequest struct {
	Name        string              `json:"name" binding:"required"`
	Title       string              `json:"title" binding:"required"`
	Banner      string              `json:"banner"`
	Location    string              `json:"location"`
	Status      string              `json:"status"`
	StartAt     string              `json:"start_at"`
	EndAt       string              `json:"end_at"`
	TicketTypes []ticketTypeRequest `json:"ticket_types"`
}

type ticketTypeRequest struct {
	Position    int     `json:"position"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	ImageURL    string  `json:"image_url"`
	Status      string  `json:"status"`
	Qty         int     `json:"qty" binding:"required"`
	Price       int     `json:"price" binding:"required"`
	Location    string  `json:"location"`
	SaleAt      *string `json:"sale_at"`
	SaleEnd     *string `json:"sale_end"`
}

func (r createEventRequest) toModel() (*models.Event, error) {
	if len(r.TicketTypes) == 0 {
		return nil, errors.New("ticket_types required")
	}

	startAt, err := parseTimeOrDefault(r.StartAt)
	if err != nil {
		return nil, err
	}
	endAt, err := parseTimeOrDefault(r.EndAt)
	if err != nil {
		return nil, err
	}

	event := &models.Event{
		Name:     r.Name,
		Title:    r.Title,
		Banner:   r.Banner,
		Location: r.Location,
		Status:   r.Status,
		StartAt:  startAt,
		EndAt:    endAt,
	}

	for _, item := range r.TicketTypes {
		saleAt, err := parseOptionalTime(item.SaleAt)
		if err != nil {
			return nil, err
		}
		saleEnd, err := parseOptionalTime(item.SaleEnd)
		if err != nil {
			return nil, err
		}
		event.TicketTypes = append(event.TicketTypes, models.TicketType{
			Position:    item.Position,
			Name:        item.Name,
			Description: item.Description,
			ImageURL:    item.ImageURL,
			Status:      item.Status,
			Qty:         item.Qty,
			Price:       item.Price,
			Location:    item.Location,
			SaleAt:      saleAt,
			SaleEnd:     saleEnd,
		})
	}

	return event, nil
}

func parseTimeOrDefault(value string) (time.Time, error) {
	if value == "" {
		return time.Now(), nil
	}
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func parseOptionalTime(value *string) (*time.Time, error) {
	if value == nil || *value == "" {
		return nil, nil
	}
	t, err := time.Parse(time.RFC3339, *value)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
