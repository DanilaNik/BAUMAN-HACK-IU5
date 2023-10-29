package ds

type Rover struct {
	ID     uint64 `json:"id"`
	Uuid   string `json:"uuid"`
	Name   string `json:"name"`
	X      uint64 `json:"x"`
	Y      uint64 `json:"y"`
	Z      uint64 `json:"z"`
	Angle  uint64 `json:"angle"`
	Charge uint64 `json:"charge"`
}
