package handler

import (
	"github.com/micro/go-micro/v2/server"
	"order-service/handler/order"
	"order-service/proto"
)

func Init(svr server.Server)  {
	proto.RegisterOrderHandler(svr, new(order.OrderHandle))
}