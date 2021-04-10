package xmlfile

type node struct {
	Name     string `xml:"Name,attr"`
	Type     string `xml:"Type,attr"`
	Username string `xml:"Username,attr"`
	Hostname string `xml:"Hostname,attr"`
	HomeDir  string `xml:"UserField,attr"`
	Password string `xml:"Password,attr"`
	Port     string `xml:"Port,attr"`
	Nodes    []node `xml:"Node"`
}
