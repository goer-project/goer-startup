package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"goer-startup/internal/goerctl/cmd/new"
	"goer-startup/internal/goerctl/cmd/version"
	genericapiserver "goer-startup/internal/pkg/server"
	"goer-startup/pkg/cli/genericclioptions"
)

func NewDefaultGoerCtlCommand() *cobra.Command {
	return NewGoerCtlCommand(os.Stdin, os.Stdout, os.Stderr)
}

func NewGoerCtlCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "goerctl",
		Short: "goerctl is the goer startup client",
		Long:  `goerctl is the client side tool for Goer startup.`,
		Run:   runHelp,
	}

	// Load config
	cobra.OnInitialize(initConfig)

	ioStreams := genericclioptions.IOStreams{In: in, Out: out, ErrOut: err}

	// Add commands
	cmd.AddCommand(new.NewCmdNew(ioStreams))
	cmd.AddCommand(version.NewCmdVersion(ioStreams))

	return cmd
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	genericapiserver.LoadConfig(viper.GetString(genericclioptions.FlagGoerConfig), "goerctl")
}

func runHelp(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}
