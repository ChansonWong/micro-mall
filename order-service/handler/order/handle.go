package order

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"
	"order-service/proto"
	"order-service/service/order"
)

type OrderHandle struct{}

func (o OrderHandle) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest, resp *proto.CommonResult) error {
	log.Log("Received Dict.ListRoot request")
	orderId := req.OrderId
	orderName := req.OrderName
	// 创建订单
	order.CreateOrder(orderId, orderName)
	resp.Code = "1"
	resp.Info = "添加成功"
	return nil
}



