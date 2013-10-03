package dangerlib

// We just need a type to bind to
type RPCFunc struct{}

// The request being made on the server
type Request struct {
	Command string
	Args    []string
}
