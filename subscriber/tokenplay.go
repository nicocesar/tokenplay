package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	tokenplay "github.com/nicocesar/tokenplay/proto/tokenplay"
)

type Tokenplay struct{}

func (e *Tokenplay) Handle(ctx context.Context, msg *tokenplay.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *tokenplay.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
