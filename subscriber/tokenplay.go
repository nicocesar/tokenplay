package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	tokenplay "github.com/nicocesar/tokenplay/proto/tokenplay"
)

type Tokenplay struct{}

func (e *Tokenplay) Handle(ctx context.Context, msg *tokenplay.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *tokenplay.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
