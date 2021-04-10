package api

// Node is a marker interface for nodes in the config tree.
type Node interface {
	childNode()
}

// Collection is a named container for an list of nodes.
type Collection struct {
	Name     string
	Children []Node
}

func (collection *Collection) childNode() {
}

func (collection *Collection) NodeName() string {
	return collection.Name
}

// Connection wraps data that defines an ssh connection.
type Connection struct {
	Name     string
	Username string
	Password string
	Hostname string
	HomeDir  string
	Port     string
	Id       uint
}

func (connection *Connection) childNode() {
}
