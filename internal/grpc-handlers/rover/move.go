package move

import (
	pb "github.com/DanilaNik/BAUMAN-HACK-IU5/github.com/DanilaNik/BAUMAN-HACK-IU5"
	"github.com/DanilaNik/BAUMAN-HACK-IU5/internal/ds"
)

type Request struct {
	Uuid string `json:"uuid"`
	Move string `json:"move"`
}

type Response struct {
	Uuid    string `json:"uuid"`
	Name    string `json:"name"`
	X       uint64 `json:"x"`
	Y       uint64 `json:"y"`
	Z       uint64 `json:"z"`
	Charge  uint64 `json:"charge"`
	Warning string `json:"warning"`
}

// func RoverLife(rover *ds.Rover) {
// 	ticker := time.NewTicker(time.Millisecond * 500)

// 	for {

// 	}
// }

type roverMover interface {
	MoveRover(uuid string, req *pb.Request) (*ds.Rover, string, error)
}

func MoveRover(req *pb.Request, moverRover roverMover) *pb.Response {
	// var req Request
	// err := json.Unmarshal([]byte(reqStr), &req)
	// if err != nil {
	// 	log.Fatalf("failed to unmarshal JSON: %v", err)
	// }

	// log.Printf("Unmarshalled Rover: %+v", req)

	rover, warning, err := moverRover.MoveRover(req.Uuid, req)
	if err != nil {
		return nil
	}

	// var resp Response = Response{
	// 	Uuid:    rover.Uuid,
	// 	Name:    rover.Name,
	// 	X:       rover.X,
	// 	Y:       rover.Y,
	// 	Z:       rover.Z,
	// 	Charge:  rover.Charge,
	// 	Warning: warning,
	// }

	// rover.X += uint64(req.X)
	// rover.Y += uint64(req.Y)
	// rover.Z += uint64(req.Z)
	resp := &pb.Response{
		Uuid:        rover.Uuid,
		X:           int64(rover.X),
		Y:           int64(rover.Y),
		Z:           int64(rover.Z),
		Charge:      float32(rover.Charge),
		Temperature: 0,
		Warning:     warning,
		Alert:       "",
	}

	// res, _ := json.Marshal(resp)
	// return string(res)
	return resp
}
