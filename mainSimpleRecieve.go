package main


func main() {
	rabbitmq := NewRabbitMQSimple("" +
		"imoocSimple")
	rabbitmq.ConsumeSimple()
}
