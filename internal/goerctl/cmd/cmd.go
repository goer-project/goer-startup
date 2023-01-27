package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"

	"goer-startup/internal/apiserver"
	"goer-startup/internal/goerctl/cmd/new"
	"goer-startup/internal/goerctl/cmd/user"
	"goer-startup/internal/goerctl/cmd/version"
	"goer-startup/internal/goerctl/util/templates"
	"goer-startup/internal/pkg/log"
	"goer-startup/pkg/cli/genericclioptions"
)

func NewDefaultGoerCtlCommand() *cobra.Command {
	return NewGoerCtlCommand(os.Stdin, os.Stdout, os.Stderr)
}

func NewGoerCtlCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	var cmds = &cobra.Command{
		Use:   "goerctl",
		Short: "goerctl is the goer startup client",
		Long:  `goerctl is the client side tool for Goer startup.`,
		Run:   runHelp,
	}

	// Load config
	cobra.OnInitialize(initConfig)

	ioStreams := genericclioptions.IOStreams{In: in, Out: out, ErrOut: err}

	groups := templates.CommandGroups{
		{
			Message: "Basic Commands:",
			Commands: []*cobra.Command{
				new.NewCmdNew(ioStreams),
			},
		},
		{
			Message: "Advanced Commands:",
			Commands: []*cobra.Command{
				user.NewCmdUser(ioStreams),
			},
		},
	}
	groups.Add(cmds)

	filters := []string{""}
	templates.ActsAsRootCommand(cmds, filters, groups...)

	// Config file
	cmds.PersistentFlags().StringVarP(&apiserver.CfgFile, "config", "c", "", "The path to the configuration file. Empty string for no configuration file.")

	// Add commands
	cmds.AddCommand(version.NewCmdVersion(ioStreams))

	return cmds
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	apiserver.InitConfig()

	// Init store
	if err := apiserver.InitStore(); err != nil {
		log.Fatalw(err.Error())
	}
}

func runHelp(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}
