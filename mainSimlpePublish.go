package main

import (
	"fmt"
)

func main() {
	rabbitmq :=NewRabbitMQSimple("" +
		"imoocSimple")
	rabbitmq.PublishSimple("Hello imooc!")
	fmt.Println("发送成功！")
}
