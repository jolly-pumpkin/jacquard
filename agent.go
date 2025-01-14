package jacquard

type Agent struct {
	Name string
}

// New creates a new Agent instance.
func New(name string) *Agent {
	return &Agent{Name: name}
}
