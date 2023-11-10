package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("starting payments microservice")

	defer log.Println("closing payments microservice")

	ctx := cmd.Context()

	paymentsInterface := createPaymentMicroservice()

	if err := paymentsInterface.Run(ctx); err != nil {
		panic(err)
	}
}

func createPaymentMicroservice() amqp.paymentsInterface {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	paymentsService := payments_app.NewPaymentService(
		payments_infra_orders.NewHTTPClient(os.Getenv("SHOP_ORDERS_SERVICE_ADDR")),
	)

	paymentsInterface, err := amqp.NewPaymentsInterface(
		fmt.Sprintf("amqp://%s/", os.Getenv("SHOP_RABBITMQ_ADDR")),
		os.Getenv("SHOP_RABBITMQ_TO_PAY_QUEUE"),
		paymentsService,
	)

	if err != nil {
		panic(err)
	}

	return paymentsInterface
}
