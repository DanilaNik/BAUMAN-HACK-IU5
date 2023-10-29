package move

import (
	"encoding/json"
	"log"

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

type roverMover interface {
	MoveRover(uuid string, move string) (*ds.Rover, string, error)
}

func MoveRover(reqStr string, moverRover roverMover) string {
	var req Request
	err := json.Unmarshal([]byte(reqStr), &req)
	if err != nil {
		log.Fatalf("failed to unmarshal JSON: %v", err)
	}

	log.Printf("Unmarshalled Rover: %+v", req)

	rover, warning, err := moverRover.MoveRover(req.Uuid, req.Move)
	if err != nil {
		return ""
	}

	var resp Response = Response{
		Uuid:    rover.Uuid,
		Name:    rover.Name,
		X:       rover.X,
		Y:       rover.Y,
		Z:       rover.Z,
		Charge:  rover.Charge,
		Warning: warning,
	}

	res, _ := json.Marshal(resp)
	return string(res)
}
