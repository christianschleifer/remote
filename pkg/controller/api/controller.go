package api

// Controller is the interface that wraps the basic ConnectionSelectedHandler and GetCollection methods.
// Implement this interface to use different business logic when interacting with Collection datastructures.
type Controller interface {
	ConnectionSelectedHandler(connectionId uint)
	GetCollection() Collection
}
