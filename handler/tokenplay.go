package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	tokenplay "github.com/nicocesar/tokenplay/proto/tokenplay"
)

type Tokenplay struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Tokenplay) Call(ctx context.Context, req *tokenplay.Request, rsp *tokenplay.Response) error {
	log.Info("Received Tokenplay.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Tokenplay) Stream(ctx context.Context, req *tokenplay.StreamingRequest, stream tokenplay.Tokenplay_StreamStream) error {
	log.Infof("Received Tokenplay.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&tokenplay.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Tokenplay) PingPong(ctx context.Context, stream tokenplay.Tokenplay_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&tokenplay.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
