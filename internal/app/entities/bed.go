package entities

//Bed: Tracks bed assignments and availability within rooms.

type Bed struct {
	Model
	Name        string
	Description string
	RoomId      string
	// room left here
}
