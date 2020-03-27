package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2"
	"github.com/nicocesar/tokenplay/handler"
	"github.com/nicocesar/tokenplay/subscriber"

	tokenplay "github.com/nicocesar/tokenplay/proto/tokenplay"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("tokenplay"),
		micro.Version("latest"),
		micro.Broker(mqtt.NewBroker()),
	)

	// Initialise service
	service.Init()

	// Register Handler
	tokenplay.RegisterTokenplayHandler(service.Server(), new(handler.Tokenplay))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("tokenplay", service.Server(), new(subscriber.Tokenplay))

	// Register Function as Subscriber
	micro.RegisterSubscriber("tokenplay", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
