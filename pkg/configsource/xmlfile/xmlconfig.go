package xmlfile

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"github.com/ChristianSchleifer/mremoteng/pkg/controller/api"
	"golang.org/x/crypto/pbkdf2"
	"io/ioutil"
)

const secret = "mR3m"

type xmlConfigSource struct {
	filepath  string
	data      *node
	idCounter uint

	collection  api.Collection
	connections []*api.Connection
}

// NewConfigSource returns a fully initialized xml file based implementation of an api.ConfigSource.
func NewConfigSource(filepath string) api.ConfigSource {
	xmlSource := &xmlConfigSource{
		filepath: filepath,
		data:     &node{},
	}
	err := xmlSource.parseConfig()
	if err != nil {
		panic(err)
	}
	return xmlSource
}

func (xmlSource *xmlConfigSource) parseConfig() error {
	data, err := ioutil.ReadFile(xmlSource.filepath)
	if err != nil {
		return errors.New("couldn't find file")
	}

	err = xml.Unmarshal(data, xmlSource.data)
	if err != nil {
		return errors.New("couldn't unmarshal data")
	}

	return nil
}

// GetConfig returns the parsed configuration data.
func (xmlSource *xmlConfigSource) GetConfig() (api.Collection, []*api.Connection) {
	xmlSource.collection = api.Collection{
		Name: "MRemoteNGConfig",
	}

	xmlSource.nodeToCollection(xmlSource.data, &xmlSource.collection)

	return xmlSource.collection, xmlSource.connections
}

func (xmlSource *xmlConfigSource) nodeToCollection(node *node, root api.Node) {

	switch node.Type {
	case "", "Container":
		collection, ok := root.(*api.Collection)
		if !ok {
			panic("a connection cannot be a parent element")
		}
		newCollection := &api.Collection{
			Name:     node.Name,
			Children: nil,
		}
		collection.Children = append(collection.Children, newCollection)

		for _, child := range node.Nodes {
			xmlSource.nodeToCollection(&child, newCollection)
		}
	case "Connection":
		collection, ok := root.(*api.Collection)
		if !ok {
			panic("a connection cannot be a parent element")
		}

		password, _ := decryptPassword(node.Password)

		newConnection := &api.Connection{
			Name:     node.Name,
			Username: node.Username,
			Password: password,
			Hostname: node.Hostname,
			HomeDir:  node.HomeDir,
			Port:     node.Port,
			Id:       xmlSource.idCounter,
		}
		collection.Children = append(collection.Children, newConnection)
		xmlSource.connections = append(xmlSource.connections, newConnection)
		xmlSource.idCounter++
	}
}

func decryptPassword(base64EncodedEncryptedPassword string) (string, error) {
	if len(base64EncodedEncryptedPassword) == 0 {
		return "", errors.New("password is empty")
	}
	encryptedPassword, err := base64.StdEncoding.DecodeString(base64EncodedEncryptedPassword)
	if err != nil {
		return "", err
	}

	salt := encryptedPassword[:16]
	additionalData := encryptedPassword[:16]
	ciphertext := encryptedPassword[32:]

	key := pbkdf2.Key([]byte(secret), salt, 1000, 32, sha1.New)

	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherMode, err := cipher.NewGCMWithNonceSize(cipherBlock, 16)
	if err != nil {
		return "", err
	}

	decryptedPassword, err := cipherMode.Open(nil, encryptedPassword[16:32], ciphertext, additionalData)
	if err != nil {
		return "", err
	}

	return string(decryptedPassword), nil
}
