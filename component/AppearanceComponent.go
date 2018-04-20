package component

// MyTurnComponent .
type AppearanceComponent struct {
	Character string
}

func (pc AppearanceComponent) GetType() string {
	return "AppearanceComponent"
}
