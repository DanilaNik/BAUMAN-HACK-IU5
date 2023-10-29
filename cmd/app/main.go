package main

import (
	"log"
	"math"
	"net"
	"sync"
	"time"

	pb "github.com/DanilaNik/BAUMAN-HACK-IU5/github.com/DanilaNik/BAUMAN-HACK-IU5"
	"github.com/DanilaNik/BAUMAN-HACK-IU5/internal/ds"
	"github.com/DanilaNik/BAUMAN-HACK-IU5/internal/strorage/sqlite"

	"google.golang.org/grpc"
)

const storagePath = "./wardenDB.db"

type server struct {
	pb.UnimplementedSimulationServer
	storage *sqlite.Storage
}

// func (s *server) AddRover(rover *ds.Rover) {
// 	s.rover = rover
// }

var rover *ds.Rover = &ds.Rover{
	ID:     123,
	Uuid:   "123",
	Name:   "rover 1",
	X:      1,
	Y:      2,
	Z:      3,
	Angle:  0,
	Charge: 100,
}

var Warning string = "none"
var Alert string = "none"

var mx *sync.Mutex

func (s *server) BidirectionalStreaming(stream pb.Simulation_BidirectionalStreamingServer) error {
	go chargeDrain()
	log.Print("Start server")
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		if req.X != 0 || req.Y != 0 || req.Z != 0 {
			go asyncMove(req)
		}
		// Handle incoming request data
		log.Printf("Received request: %s", req)

		// Create and send the response
		err = stream.Send(&pb.Response{
			Uuid:        s.storage.Rover.Uuid,
			X:           int64(rover.X),
			Y:           int64(rover.Y),
			Z:           int64(rover.Z),
			Charge:      int64(rover.Charge),
			Temperature: 10,
			Warning:     Warning,
			Alert:       Alert,
		})
		if err != nil {
			return err
		}
	}
}

func asyncMove(change *pb.Request) {
	wg := &sync.WaitGroup{}
	Warning = "moving"
	wg.Add(3)
	go func() {
		step := 1
		if change.X < 0 {
			step = -1
		}
		for i := 0; i < int(math.Abs(float64(change.X))); i++ {
			rover.X += int64(step)
			mx.Lock()
			rover.Charge -= 0.01
			mx.Unlock()
			time.Sleep(time.Millisecond * 200)
		}
		wg.Done()
	}()
	go func() {
		step := 1
		if change.Y < 0 {
			step = -1
		}
		for i := 0; i < int(math.Abs(float64(change.Y))); i++ {
			rover.Y += int64(step)
			mx.Lock()
			rover.Charge -= 0.01
			mx.Unlock()
			time.Sleep(time.Millisecond * 200)
		}
		wg.Done()
	}()
	go func() {
		step := 1
		if change.Z < 0 {
			step = -1
		}
		for i := 0; i < int(math.Abs(float64(change.Z))); i++ {
			rover.Z += int64(step)
			mx.Lock()
			rover.Charge -= 0.01
			mx.Unlock()
			time.Sleep(time.Millisecond * 500)
		}
		wg.Done()
	}()
	wg.Wait()
	Warning = "none"
}

func chargeDrain() {
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			mx.Lock()
			if rover.Charge == 0 {
				Alert = "no charge"
			} else {
				rover.Charge -= 1
			}
			mx.Unlock()
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
	log.Print("test")
	s := grpc.NewServer()

	pb.RegisterSimulationServer(s, &server{
		storage: storage,
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
