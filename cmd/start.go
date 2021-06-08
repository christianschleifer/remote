package cmd

import (
	"github.com/ChristianSchleifer/mremoteng/pkg/configsource/xmlfile"
	"github.com/ChristianSchleifer/mremoteng/pkg/connectionhandler/gnome"
	"github.com/ChristianSchleifer/mremoteng/pkg/connectionhandler/ssh"
	"github.com/ChristianSchleifer/mremoteng/pkg/controller"
	"github.com/ChristianSchleifer/mremoteng/pkg/controller/api"
	"github.com/ChristianSchleifer/mremoteng/pkg/viewer/terminal"
	"github.com/spf13/cobra"
	"path/filepath"
)

const GnomeHandler = "gnome"
const SshHandler = "ssh"
const TerminalViewer = "terminal"

var configFile string
var handlerConfig string
var viewerConfig string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Runs the graphical user interface",
	Run:   startCmdFn,
}

func init() {
	startCmd.Flags().StringVarP(
		&configFile,
		"file",
		"f",
		"",
		"Location of the mRemoteNG config file (required)",
	)
	_ = startCmd.MarkFlagRequired("file")

	startCmd.Flags().StringVarP(
		&handlerConfig,
		"handler",
		"",
		GnomeHandler,
		"The specific connection handler to be used to handle connections",
	)

	startCmd.Flags().StringVarP(
		&viewerConfig,
		"viewer",
		"",
		TerminalViewer,
		"The specific viewer to be used to interact with the mRemoteNG config",
	)
}

func startCmdFn(_ *cobra.Command, _ []string) {
	var handler api.ConnectionHandler
	var viewer api.Viewer

	absolutePath, err := filepath.Abs(configFile)
	if err != nil {
		panic("could not find file at location " + absolutePath)
	}
	source := xmlfile.NewConfigSource(absolutePath)

	switch handlerConfig {
	case GnomeHandler:
		handler = gnome.NewHandler()
	case SshHandler:
		handler = ssh.NewHandler()
	default:
		panic("Currently, only 'gnome' and 'ssh' connection handler exists.")
	}

	ctrl := controller.NewController(source, handler)

	if viewerConfig == TerminalViewer {
		viewer = terminal.NewViewer(ctrl)
	} else {
		panic("Currently, only a 'terminal' viewer exists.")
	}

	viewer.View()
}
