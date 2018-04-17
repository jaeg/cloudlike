package component

// InitiativeComponent .
type InitiativeComponent struct {
	DefaultValue  int
	OverrideValue int
	Ticks         int
}

func (pc InitiativeComponent) GetType() string {
	return "InitiativeComponent"
}
