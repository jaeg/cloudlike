package component

// Component base component interface
type DecayingComponent interface {
	Decay() bool
	GetType() string
}
