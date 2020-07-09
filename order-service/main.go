package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
	_ "order-service/database"
)

// 订单服务

func main(){
	// 链接RabbitMq
	broker := rabbitmq.NewBroker(
		broker.Addrs("amqp://guest:guest@localhost:5672"),
	)

	broker.Init()
	broker.Connect()

	// New Service
	service := micro.NewService(
		micro.Name("superatom.svc.order"),
		micro.Version("latest"),
		micro.Broker(broker),
	)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
