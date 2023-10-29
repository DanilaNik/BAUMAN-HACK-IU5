package main

import (
	"log"
	"net"

	pb "github.com/DanilaNik/BAUMAN-HACK-IU5/github.com/DanilaNik/BAUMAN-HACK-IU5"
	move "github.com/DanilaNik/BAUMAN-HACK-IU5/internal/grpc-handlers/rover"
	"github.com/DanilaNik/BAUMAN-HACK-IU5/internal/strorage/sqlite"

	"google.golang.org/grpc"
)

const storagePath = "./wardenDB.db"

type server struct {
	pb.UnimplementedMyServiceServer
	storage *sqlite.Storage
}

func (s *server) BidirectionalStreaming(stream pb.MyService_BidirectionalStreamingServer) error {
	log.Print("Start server")
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		response := move.MoveRover(req.Data, s.storage)

		// Handle incoming request data
		log.Printf("Received request: %s", req.GetData())

		// Create and send the response
		err = stream.Send(&pb.Response{
			Result: response,
		})
		if err != nil {
			return err
		}
	}
}

func main() {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		log.Fatalf("failed to init storage %v", err)
	}
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{storage: storage})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
