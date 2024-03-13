package order_repo

import "ASSIGMENT2/entity"

type Repository interface {
	CreateOrderWithItems(order entity.Order, items []entity.Item) error
	
}