package order_pg

import (
	"ASSIGMENT2/entity"
	"ASSIGMENT2/repository/order_repo"
	"database/sql"
)

type orderPG struct {
	db *sql.DB
}

// untuk ngerturn instance yang telah mengimplementasikan suatu interface

func NewRepository(db *sql.DB) order_repo.Repository{
	return &orderPG{
		db: db,
	}
}
func (o *orderPG) CreateOrderWithItems(order entity.Order, items []entity.Item) error{
	// o.db.Exec("INSERT INTO orders")

	tx, err := o.db.Begin()

	if err != nil{

		return err
	}
	_, err = tx.Exec(create_new_order, order.OrderAt, order.CustomerName)

	if err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range items{
		_, err = tx.Exec(
		create_new_item, 
		item.ItemCode,
		item.Description,
		item.Description,
		item.Quantity,
		)
	}

	return nil
}