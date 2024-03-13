package order_service

import (
	"ASSIGMENT2/dto"
	"ASSIGMENT2/entity"
	"ASSIGMENT2/repository/order_repo"
)

type Service interface {
	CreateOrderWithItems(payload dto.CreateOrderRequestDto) error
}

type orderService struct {
	orderRepo order_repo.Repository
}

func NewService(orderRepo order_repo.Repository) Service {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (o *orderService) CreateOrderWithItems(payload dto.CreateOrderRequestDto) error {
	// Implementasi fungsi CreateOrderWithItems di sini menggunakan o.orderRepo

	order := entity.Order{
		OrderAt: payload.OrderAt,
		CustomerName: payload.CustomerName,
	}
	items := []entity.Item{}

	for _, value := range payload.Items{
		item := entity.Item{
			ItemCode: value.Description,
			Quantity: uint(value.Quantity),
			Description: value.Description,
		}
		items = append(items, item)
	}

	err := o.orderRepo.CreateOrderWithItems(order, items)

	if err != nil{
		return err
	}
	
	return nil
}
