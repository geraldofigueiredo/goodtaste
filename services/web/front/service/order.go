package service

import (
	"github.com/geraldofigueiredo/goodtaste/services/web/front/entity"
	"github.com/google/uuid"
	"time"
)

type OrderService interface {
	NewOrder() error
	GetOrder(id string) (*entity.Order, error)
}

type orderService struct{}

func NewOrderService() OrderService {
	return &orderService{}
}

func (o orderService) NewOrder() error {
	return nil
}

func (o orderService) GetOrder(id string) (*entity.Order, error) {
	return &entity.Order{
		ID:        uuid.New(),
		OwnerName: "owner name",
		Infos:     []string{},
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}, nil
}
