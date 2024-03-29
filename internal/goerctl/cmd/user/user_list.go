package user

import (
	"context"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"goer-startup/internal/apiserver/biz"
	"goer-startup/internal/apiserver/store"
	"goer-startup/internal/goerctl/util/templates"
	"goer-startup/pkg/cli/genericclioptions"

	cmdutil "goer-startup/internal/goerctl/cmd/util"
)

const (
	listUsageStr = "list"
	defaultLimit = 100
)

// ListOptions is an option struct to support 'list' sub command.
type ListOptions struct {
	// options
	Offset int
	Limit  int

	b biz.IBiz
	genericclioptions.IOStreams
}

var (
	listExample = templates.Examples(`
		# List all users
		user list

		# List users with limit and offset
		user list --offset=0 --limit=10`)
)

// NewListOptions returns an initialized ListOptions instance.
func NewListOptions(ioStreams genericclioptions.IOStreams) *ListOptions {
	return &ListOptions{
		IOStreams: ioStreams,
		Offset:    0,
		Limit:     defaultLimit,
	}
}

// NewCmdList returns new initialized instance of 'list' sub command.
func NewCmdList(ioStreams genericclioptions.IOStreams) *cobra.Command {
	o := NewListOptions(ioStreams)

	cmd := &cobra.Command{
		Use:                   listUsageStr,
		DisableFlagsInUseLine: true,
		Aliases:               []string{},
		Short:                 "Display all users",
		TraverseChildren:      true,
		Long:                  "Display all users",
		Example:               listExample,
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(o.Complete(cmd, args))
			cmdutil.CheckErr(o.Validate(cmd, args))
			cmdutil.CheckErr(o.Run(args))
		},
		SuggestFor: []string{},
	}

	// mark flag as deprecated
	cmd.Flags().IntVarP(&o.Offset, "offset", "o", o.Offset, "Specify the offset of the first row to be returned.")
	cmd.Flags().IntVarP(&o.Limit, "limit", "l", o.Limit, "Specify the amount records to be returned.")

	return cmd
}

// Complete completes all the required options.
func (o *ListOptions) Complete(cmd *cobra.Command, args []string) error {
	o.b = biz.NewBiz(store.S)

	return nil
}

// Validate makes sure there is no discrepancy in command options.
func (o *ListOptions) Validate(cmd *cobra.Command, args []string) error {
	return nil
}

// Run executes a list sub command using the specified options.
func (o *ListOptions) Run(args []string) error {
	resp, err := o.b.Users().List(context.Background(), o.Offset, o.Limit)
	if err != nil {
		return err
	}

	data := make([][]string, 0, 1)
	for _, user := range resp.Data {
		data = append(data, []string{
			user.Username,
			user.Nickname,
			user.Email,
			user.Phone,
			user.CreatedAt.Format("2006-01-02 15:04:05"),
			user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	table := tablewriter.NewWriter(o.Out)
	table = setHeader(table)
	table = cmdutil.TableWriterDefaultConfig(table)
	table.AppendBulk(data)
	table.Render()

	return nil
}
