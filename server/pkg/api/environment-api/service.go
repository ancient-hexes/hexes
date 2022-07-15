package environmentAPI

import (
	"log"
	"strconv"
	"time"

	pb "hexes/proto/go/hexes/v1"
)

type Service struct {
	*pb.UnimplementedEnvironmentAPIServer
}

func New() *Service {
	return &Service{}
}

func (s *Service) List(request *pb.Environment_ListRequest, stream pb.EnvironmentAPI_ListServer) error {
	log.Printf("EnvironmentAPI/List %s\n", request)

	for i := 0; i < 5; i++ {
		response := &pb.Environment_ListResponse{
			Items: []*pb.Environment{
				{
					Id: &pb.Environment_ID{
						Key: strconv.Itoa(i),
					},
				},
			},
		}

		if err := stream.Send(response); err != nil {
			return err
		}

		time.Sleep(50 * time.Millisecond)
	}
	return nil
}
