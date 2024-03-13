package entity

import "time"

type Order struct {
	OrderId      uint
	CustomerName string
	OrderAt      time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}