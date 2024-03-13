package dto

import "time"

type CreateOrderRequestDto struct {
	OrderAt 		time.Time 				`json:"orderAt"`
	CustomerName	string					`json:"customerName"`
	Items 			[]CreateItemRequestDto	`json:"items"`

}