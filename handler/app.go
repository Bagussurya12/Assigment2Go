package handler

import (
	db "ASSIGMENT2/infra/database"
	"ASSIGMENT2/repository/order_repo/order_pg"
	"ASSIGMENT2/service/order_service"

	"github.com/gin-gonic/gin"
) 

func StartApp() {

	db.InitialaizeDatabase()

	pgSql := db.GetDatabaseInstance()

	_ = pgSql

	orderRepo := order_pg.NewRepository(pgSql)

	_= orderRepo

	orderService := order_service.NewService(orderRepo)
	
	oh := NewOrderHandler(orderService)

	_ = orderService

	// route := gin.Default
	route := gin.Default()
	route.POST("/orders", oh.CreateOrderWithItems)
	route.Run(":8080")
	
}