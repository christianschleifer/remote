package api

// ConnectionHandler is the interface that wraps the basic Handle method.
// Implement this interface to provide different ways of handling a Connection.
type ConnectionHandler interface {
	TransferControlForUI() bool
	Handle(connection Connection)
}
