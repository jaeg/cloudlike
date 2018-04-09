package component

// HelloWorldComponent test
type HelloWorldComponent struct {
	ID int
}

//GetType Get the type of component
func (HelloWorldComponent) GetType() string {
	return "Hello World"
}
