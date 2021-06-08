package ssh

import (
	"github.com/ChristianSchleifer/mremoteng/pkg/controller/api"
	"os"
	"os/exec"
	"syscall"
)

type sshHandler struct{}


// NewHandler returns a fully initialized bash-dependent implementation of an api.ConnectionHandler.
func NewHandler() api.ConnectionHandler {
	return &sshHandler{}
}

// Returns if the  UI has to be release before the handler is invoked
func (handler *sshHandler) TransferControlForUI() bool {
	return true
}

// Handle opens up a new terminal window and creates a ssh connection using the passed api.Connection data.
func (handler *sshHandler) Handle(connection api.Connection) {

	var binary string
	var lookErr error
	var sshArgumentsPrefix []string
	var sshArguments []string

	if connection.Password != "" {
		binary, lookErr = exec.LookPath("sshpass")
		if lookErr != nil {
			panic(lookErr)
		}
		
		sshArgumentsPrefix = []string{
			binary,
			"-p",
			connection.Password,
			"ssh",
		}


	} else {	
		binary, lookErr = exec.LookPath("ssh")
		if lookErr != nil {
			panic(lookErr)
		}
	
		sshArgumentsPrefix = []string{binary}

	}	
	
	authMethod := "PreferredAuthentications=password"
	keepAlive := "ServerAliveInterval=60"
	strictHostKeyChecking := "StrictHostKeyChecking=no"
	changeToHomeDir := "cd " + connection.HomeDir
	sshCommand := changeToHomeDir + "; bash --login"

	sshArgumentsPostfix := []string{
		"-o",
		authMethod,
		"-o",
		keepAlive,
		"-o",
		strictHostKeyChecking,
		"-p",
		connection.Port,
		"-l",
		connection.Username,
		connection.Hostname,
		"-t",
		sshCommand,
	}

	sshArguments = append(sshArgumentsPrefix, sshArgumentsPostfix...)

	err := syscall.Exec(binary, sshArguments, os.Environ())
	panic(err)
}
