package api

// Viewer is the interface that wraps the basic View method.
// Implement this interface to provide different kinds of visual representations of the Collection data.
type Viewer interface {
	View()
}
