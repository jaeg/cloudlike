package component

// DecayingComponent Component base component interface
type DecayingComponent interface {
	Decay() bool
	GetType() string
}
