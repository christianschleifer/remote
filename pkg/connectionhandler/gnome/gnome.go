package gnome

import (
	"fmt"
	"github.com/ChristianSchleifer/mremoteng/pkg/controller/api"
	"os/exec"
)

type gnomeSshHandler struct{}

// NewHandler returns a fully initialized gnome-dependent implementation of an api.ConnectionHandler.
func NewHandler() api.ConnectionHandler {
	return &gnomeSshHandler{}
}

// Handle opens up a new terminal window and creates a ssh connection using the passed api.Connection data.
func (handler *gnomeSshHandler) Handle(connection api.Connection) {
	gnomeArguments := []string{"-x"}

	sshCommand := "ssh"
	authMethod := "PreferredAuthentications=password"
	changeToHomeDir := "cd " + connection.HomeDir
	target := fmt.Sprintf("%s@%s", connection.Username, connection.Hostname)

	sshConfiguration := []string{
		sshCommand,
		"-o",
		authMethod,
		"-p",
		connection.Port,
		target,
		"-t",
		changeToHomeDir,
		";",
		"bash",
		"--login",
	}

	if connection.Password != "" {
		sshConfiguration = append([]string{
			"sshpass",
			"-p",
			connection.Password,
		}, sshConfiguration...)
	}

	arguments := append(gnomeArguments, sshConfiguration...)

	cmd := exec.Command("gnome-terminal", arguments...)

	if err := cmd.Start(); err != nil {
		panic(err)
	}
}
