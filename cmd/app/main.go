package main

import (
	"log"
	"net"

	pb "github.com/DanilaNik/BAUMAN-HACK-IU5/github.com/DanilaNik/BAUMAN-HACK-IU5"
	"github.com/DanilaNik/BAUMAN-HACK-IU5/internal/ds"
	move "github.com/DanilaNik/BAUMAN-HACK-IU5/internal/grpc-handlers/rover"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMyServiceServer
}

func (s *server) BidirectionalStreaming(stream pb.MyService_BidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		var rover *ds.Rover = &ds.Rover{
			ID:     1,
			Uuid:   "00112233-4455-6677-8899-saabbccddeeff",
			Name:   "Rover X",
			X:      0,
			Y:      0,
			Charge: 86,
		}

		response := move.MoveRover(rover, req.Data)

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
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// import (
// 	"log"
// 	"net"

// 	pb "github.com/DanilaNik/BAUMAN-HACK-IU5/github.com/DanilaNik/BAUMAN-HACK-IU5"
// 	"google.golang.org/grpc"
// 	//"github.com/sirupsen/logrus"
// )

// type server struct {
// 	pb.UnimplementedMyServiceServer
// }

// func (s *server) BidirectionalStreaming(stream pb.MyService_BidirectionalStreamingServer) error {
// 	for {
// 		req, err := stream.Recv()
// 		if err != nil {
// 			return err
// 		}

// 		// Handle incoming request data
// 		log.Printf("Received request: %s", req.GetData())

// 		// Create and send the response
// 		err = stream.Send(&pb.Response{
// 			Result: "Response from server",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 	}
// }

// func main() {
// 	lis, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterMyServiceServer(s, &server{})
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }

// func main() {
// 	err := runGRPCClient()
// 	if err != nil {
// 		log.Fatalf("failed to run gRPC client: %v", err)
// 	}

// 	// Start HTTP server
// 	err = runHTTPServer()
// 	if err != nil {
// 		log.Fatalf("failed to run HTTP server: %v", err)
// 	}
// }

// func runGRPCClient() error {
// 	conn, err := grpc.Dial("192.168.137.68:50051", grpc.WithInsecure())
// 	if err != nil {
// 		return err
// 	}
// 	defer conn.Close()

// 	client := stationpb.NewMyServiceClient(conn)

// 	stream, err := client.BidirectionalStreaming(context.Background())
// 	if err != nil {
// 		return err
// 	}

// 	go func() {
// 		// Send multiple requests
// 		for i := 0; i < 5; i++ {
// 			req := &stationpb.Request{
// 				Data: "Request from client",
// 			}
// 			err := stream.Send(req)
// 			if err != nil {
// 				log.Fatalf("error sending request: %v", err)
// 			}
// 			time.Sleep(time.Second)
// 		}
// 		stream.CloseSend()
// 	}()

// 	go func() {
// 		// Receive and print responses
// 		for {
// 			res, err := stream.Recv()
// 			if err != nil {
// 				log.Printf("finished receiving: %v", err)
// 				break
// 			}
// 			log.Printf("Received response: %s", res.GetResult())
// 		}
// 	}()

// 	return nil
// }

// func runHTTPServer() error {
// 	// Start HTTP server
// 	err := http.ListenAndServe("192.168.137.151:50051", nil)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
