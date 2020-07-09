package order

import "github.com/micro/go-micro/v2/util/log"
func CreateOrder(orderId, orderName string){
	log.Infof("成功创建订单，订单号=%s，订单名称=%s", orderId, orderName)

}
