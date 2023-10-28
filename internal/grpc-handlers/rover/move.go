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
	Uuid   string `json:"uuid"`
	Name   string `json:"name"`
	X      uint64 `json:"x"`
	Y      uint64 `json:"y"`
	Angle  uint64 `json:"angle"`
	Charge uint64 `json:"charge"`
}

var rover ds.Rover = ds.Rover{
	ID:     1,
	Uuid:   "00112233-4455-6677-8899-saabbccddeeff",
	Name:   "Rover X",
	X:      50,
	Y:      50,
	Charge: 86,
}

func MoveRover(rover1 *ds.Rover, reqStr string) string {
	var req Request
	err := json.Unmarshal([]byte(reqStr), &req)
	if err != nil {
		log.Fatalf("failed to unmarshal JSON: %v", err)
	}

	log.Printf("Unmarshalled Rover: %+v", req)

	switch req.Move {
	case "up":
		rover.Y -= 1
	case "down":
		rover.Y += 1
	case "right":
		rover.X += 1
	case "left":
		rover.X -= 1
	}

	res, _ := json.Marshal(rover)
	return string(res)
}
