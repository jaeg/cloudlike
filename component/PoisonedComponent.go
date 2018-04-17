package component

// MyTurnComponent .
type PoisonedComponent struct {
	Duration int
}

func (pc PoisonedComponent) GetType() string {
	return "PoisonedComponent"
}

func (pc *PoisonedComponent) Decay() bool {
	pc.Duration--

	return pc.Duration <= 0
}
