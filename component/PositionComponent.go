package component

// PositionComponent .
type PositionComponent struct {
	X, Y  int
	Level int
}

func (pc PositionComponent) GetType() string {
	return "PositionComponent"
}
