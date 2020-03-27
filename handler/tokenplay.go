package handler

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"
	tokenplay "github.com/nicocesar/tokenplay/proto/tokenplay"
)

type Tokenplay struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Tokenplay) AddWord(ctx context.Context, req *tokenplay.Request, rsp *tokenplay.Response) error {
	log.Log("Received Tokenplay.AddWord request")

	// ToDo: We add the word to the database.

	rsp.Msg = "Added the word " + req.Word

	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Tokenplay) Stream(ctx context.Context, req *tokenplay.StreamingRequest, stream tokenplay.Tokenplay_StreamStream) error {
	log.Logf("Received Tokenplay.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
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
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&tokenplay.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
