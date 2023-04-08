package service

type OrderService interface {
	NewOrder() error
	GetOrder() error
}

type orderService struct{}

func NewOrderService() OrderService {
	return &orderService{}
}

func (o orderService) NewOrder() error {
	return nil
}

func (o orderService) GetOrder() error {
	return nil
}
