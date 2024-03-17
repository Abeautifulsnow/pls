package cmd

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func ListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all commands that are available",
		Run: func(_ *cobra.Command, _ []string) {
			listCommand()
		},
		Aliases: []string{"l", "lis", "ls"},
	}

	return cmd
}

func listCommand() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"command", "description"})
	table.SetRowLine(true)

	for k, v := range unmarshalIndex() {
		table.Append([]string{k, v.Description})
	}

	table.Render()
}
