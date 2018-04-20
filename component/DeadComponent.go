package component

// DeadComponent .
type DeadComponent struct {
}

func (pc DeadComponent) GetType() string {
	return "DeadComponent"
}
