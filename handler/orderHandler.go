package handler

import (
	"ASSIGMENT2/dto"
	"ASSIGMENT2/service/order_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderService order_service.Service
}

func NewOrderHandler(orderService order_service.Service) *orderHandler {
	return &orderHandler{
		orderService: orderService,
	}
}

func (o *orderHandler) CreateOrderWithItems(ctx *gin.Context) {
	// Implementasi logika untuk menangani permintaan HTTP di sini
	var payload dto.CreateOrderRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": "invalid Request Json",
		})
		return
	}
	if err := o.orderService.CreateOrderWithItems(payload); err != nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,map[string]string{
			"message" : "invalid Order Service",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Order created successfully"})
}
