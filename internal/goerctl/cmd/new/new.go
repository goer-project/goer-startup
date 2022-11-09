// Package new used to generate demo command code.
// nolint: predeclared
package new

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/marmotedu/component-base/pkg/util/fileutil"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	cmdutil "goer-startup/internal/goerctl/cmd/util"
	"goer-startup/internal/goerctl/util/templates"
	"goer-startup/pkg/cli/genericclioptions"
)

const (
	newUsageStr = "new CMD_NAME | CMD_NAME CMD_DESCRIPTION"
)

var (
	//go:embed stubs
	stubsFS embed.FS

	newLong = templates.LongDesc(`Used to generate demo command source code file.

Can use this command generate a command template file, and do some modify based on your needs.
This can improve your R&D efficiency.`)

	newExample = templates.Examples(`
		# Create a default 'test' command file without a description
		new test

		# Create a default 'test' command file in /tmp/
		new test -d /tmp/

		# Create a default 'test' command file with a description
		new test "This is a test command"

		# Create command 'test' with two subcommands
		new -g test "This is a test command with two subcommands"`)

	newUsageErrStr = fmt.Sprintf(
		"expected '%s'.\nat least CMD_NAME is a required argument for the new command",
		newUsageStr,
	)

	// Read stub.
	cmdTemplate     string
	mainCmdTemplate string
	subCmd1Template string
	subCmd2Template string
)

// NewOptions is an option struct to support 'new' sub command.
type NewOptions struct {
	Group     bool
	Directory string

	// command template options, will render to command template
	CommandName         string
	CommandDescription  string
	CommandFunctionName string
	Dot                 string

	genericclioptions.IOStreams
}

// NewNewOptions returns an initialized NewOptions instance.
func NewNewOptions(ioStreams genericclioptions.IOStreams) *NewOptions {
	return &NewOptions{
		Group:              false,
		Directory:          ".",
		CommandDescription: "A brief description of your command",
		Dot:                "`",
		IOStreams:          ioStreams,
	}
}

// NewCmdNew returns new initialized instance of 'new' sub command.
func NewCmdNew(ioStreams genericclioptions.IOStreams) *cobra.Command {
	o := NewNewOptions(ioStreams)

	cmd := &cobra.Command{
		Use:                   newUsageStr,
		DisableFlagsInUseLine: true,
		Short:                 "Generate demo command code",
		Long:                  newLong,
		Example:               newExample,
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(o.Complete(cmd, args))
			cmdutil.CheckErr(o.Validate(cmd))
			cmdutil.CheckErr(o.Run(args))
		},
		Aliases:    []string{},
		SuggestFor: []string{},
	}

	cmd.Flags().BoolVarP(&o.Group, "group", "g", o.Group, "Generate two subcommands.")
	cmd.Flags().StringVarP(&o.Directory, "directory", "d", o.Directory, "Where to create demo command files.")

	return cmd
}

// Complete completes all the required options.
func (o *NewOptions) Complete(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return cmdutil.UsageErrorf(cmd, newUsageErrStr)
	}

	o.CommandName = strings.ToLower(args[0])
	if len(args) > 1 {
		o.CommandDescription = args[1]
	}

	o.CommandFunctionName = cases.Title(language.English).String(o.CommandName)

	// Read stub
	cmdTemplateBytes, _ := stubsFS.ReadFile("stubs/cmd.stub")
	mainCmdTemplateBytes, _ := stubsFS.ReadFile("stubs/cmd_main.stub")
	subCmd1TemplateBytes, _ := stubsFS.ReadFile("stubs/cmd_sub1.stub")
	subCmd2TemplateBytes, _ := stubsFS.ReadFile("stubs/cmd_sub2.stub")

	cmdTemplate = string(cmdTemplateBytes)
	mainCmdTemplate = string(mainCmdTemplateBytes)
	subCmd1Template = string(subCmd1TemplateBytes)
	subCmd2Template = string(subCmd2TemplateBytes)

	return nil
}

// Validate makes sure there is no discrepancy in command options.
func (o *NewOptions) Validate(cmd *cobra.Command) error {
	return nil
}

// Run executes a new sub command using the specified options.
func (o *NewOptions) Run(args []string) error {
	if o.Group {
		return o.CreateCommandWithSubCommands()
	}

	return o.CreateCommand()
}

// CreateCommand create the command with options.
func (o *NewOptions) CreateCommand() error {
	return o.GenerateGoCode(o.CommandName+".go", cmdTemplate)
}

// CreateCommandWithSubCommands create sub commands with options.
func (o *NewOptions) CreateCommandWithSubCommands() error {
	if err := o.GenerateGoCode(o.CommandName+".go", mainCmdTemplate); err != nil {
		return err
	}

	if err := o.GenerateGoCode(o.CommandName+"_sub1.go", subCmd1Template); err != nil {
		return err
	}

	if err := o.GenerateGoCode(o.CommandName+"_sub2.go", subCmd2Template); err != nil {
		return err
	}

	return nil
}

// GenerateGoCode generate go source file.
func (o *NewOptions) GenerateGoCode(name, codeTemplate string) error {
	tmpl, err := template.New("cmd").Parse(codeTemplate)
	if err != nil {
		return err
	}

	err = fileutil.EnsureDirAll(o.Directory)
	if err != nil {
		return err
	}

	filename := filepath.Join(o.Directory, name)
	fd, err := os.Create(filename)
	if err != nil {
		return err
	}
	// defer fd.Close()

	err = tmpl.Execute(fd, o)
	if err != nil {
		return err
	}

	fmt.Printf("Command file generated: %s\n", filename)

	return nil
}
