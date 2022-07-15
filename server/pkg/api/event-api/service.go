package eventAPI

import (
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	pb "hexes/proto/go/hexes/v1"
)

type Service struct {
	*pb.UnimplementedEventAPIServer
}

func New() *Service {
	return &Service{}
}

func (s *Service) Connect(stream pb.EventAPI_ConnectServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)

	var recvErr error
	var sendErr error

	go func() {
		defer wg.Done()
		for {
			var event *pb.Event
			event, recvErr = stream.Recv()
			if recvErr == io.EOF {
				log.Printf("received an EOF from client")
				recvErr = nil
				break
			}
			if recvErr != nil {
				log.Printf("got an error receiving from stream: %s", recvErr)
				break
			}
			s.processClientEvent(event)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			event := &pb.Event{
				Payload: &pb.Event_Chat{
					Chat: &pb.PayloadChat{
						Text: fmt.Sprintf("[server] %d", i),
					},
				},
			}
			if sendErr = stream.Send(event); sendErr != nil {
				log.Printf("got an error sending to stream: %s", sendErr)
				break
			}
			time.Sleep(700 * time.Millisecond)
		}
	}()

	wg.Wait()
	log.Printf(
		"EventAPI/Connect finished: recvErr = %s | sendErr = %s",
		recvErr, sendErr,
	)

	if recvErr != nil {
		return fmt.Errorf("receiving from client: %w", recvErr)
	}
	if sendErr != nil {
		return fmt.Errorf("sending to client: %w", sendErr)
	}
	return nil
}

func (s *Service) processClientEvent(e *pb.Event) {
	log.Printf("received client event: %s", e.GetChat().Text)
}
