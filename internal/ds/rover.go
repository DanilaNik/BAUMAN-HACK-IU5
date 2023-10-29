package ds

type Rover struct {
	ID     uint64  `json:"id"`
	Uuid   string  `json:"uuid"`
	Name   string  `json:"name"`
	X      int64   `json:"x"`
	Y      int64   `json:"y"`
	Z      int64   `json:"z"`
	Angle  uint64  `json:"angle"`
	Charge float64 `json:"charge"`
}
