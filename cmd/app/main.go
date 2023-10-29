package main

import (
	"encoding/json"
	"log"
	"math"
	"net"
	"os"
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

var storage *sqlite.Storage = sqlite.New(storagePath)

var rover *ds.Rover = storage.GetRoverByUUID("00112233-4455-6677-8899-aabbccddeeff")

//var rover *ds.Rover = ds.Rover{
// 	ID:     123,
// 	Uuid:   "123",
// 	Name:   "rover 1",
// 	X:      1,
// 	Y:      2,
// 	Z:      3,
// 	Angle:  0,
// 	Charge: 100,
// }

var Temperature float32 = 0

func temp(z int64) float32 {

	return float32(z) / 100 * 10
}

var Warning string = "none"
var Alert string = "none"

var mx sync.Mutex

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
			Name:        rover.Name,
			X:           int64(rover.X),
			Y:           int64(rover.Y),
			Z:           int64(rover.Z),
			Charge:      float32(rover.Charge),
			Temperature: temp(rover.Z),
			Warning:     Warning,
			Alert:       Alert,
		})

		log.Printf("Sent response: %v", rover)
		if err != nil {
			return err
		}
	}
}

func asyncMove(change *pb.Request) {
	wg := &sync.WaitGroup{}
	Warning = "moving"
	if rover.Charge <= 0 {
		return
	}
	wg.Add(3)

	go func() {
		step := 1
		if change.X < 0 {
			step = -1
		}
		for i := 0; i < int(math.Abs(float64(change.X))); i++ {
			tmp := rover.X + int64(step)
			if tmp <= 0 || tmp >= 100 {
				continue
			}
			// if int64(PointMap[uint64(rover.X)][uint64(rover.Y)]) <= rover.Z {
			// 	rover.Z -= 1
			// }
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
			// if rover.Y >= 1 && rover.Y <= 94 {
			tmp := rover.Y + int64(step)
			if tmp <= 0 || tmp >= 100 {
				continue
			}
			// if int64(PointMap[uint64(rover.X)][uint64(rover.Y)]) <= rover.Z {
			// 	rover.Z -= 1
			// }
			rover.Y += int64(step)
			// }
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
			tmp := rover.Z + int64(step)
			if tmp <= 0 || tmp >= 100 {
				continue
			}
			// if tmp >= int64(PointMap[uint64(rover.X)][uint64(rover.Y)]) {
			// 	continue
			// }
			rover.Z += int64(step)

			mx.Lock()
			if rover.Charge-0.01 >= 0 {
				rover.Charge -= 0.01
			}
			mx.Unlock()
			time.Sleep(time.Millisecond * 500)
		}
		wg.Done()
	}()
	wg.Wait()
	Warning = "none"
}

func chargeDrain() {
	ticker := time.NewTicker(time.Second * 20)
	for {
		select {
		case <-ticker.C:
			mx.Lock()
			if rover.Charge == 0 {
				Alert = "no charge"
			} else {
				if rover.Charge-0.1 >= 0 {
					rover.Charge -= 0.1
				}
			}
			mx.Unlock()
			//storage.UpdateRover(rover)
		}
	}
}

type Point struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
	Z int64 `json:"z"`
}

var PointArr = make([]Point, 0)
var PointMap = make(map[uint64]map[uint64]uint64)

func main() {
	//storage, err := sqlite.New(storagePath)
	// if err != nil {
	// 	log.Fatalf("failed to init storage %v", err)
	// }

	fileBytes, _ := os.ReadFile("result.json")
	_ = json.Unmarshal(fileBytes, PointArr)
	for _, v := range PointArr {
		PointMap[uint64(v.X)][uint64(v.Y)] = uint64(v.Z)
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
