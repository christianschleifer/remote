package api

// ConfigSource is the interface that wraps the basic GetConfig method.
// Implement this interface to provide different kinds of data sources.
type ConfigSource interface {
	GetConfig() (Collection, []*Connection)
}
