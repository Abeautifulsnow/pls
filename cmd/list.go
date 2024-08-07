package cmd

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func ListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all commands that are available",
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
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
	)

	table.SetRowLine(true)

	requestData := unmarshalIndex()
	data := make([][]string, 0, len(requestData))

	for k, v := range requestData {
		data = append(data, []string{k, v.Description})
	}

	// Set color to column and last row.
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
	)
	table.SetFooter([]string{"Total", strconv.Itoa(len(requestData))})
	table.SetFooterColor(
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.FgHiRedColor},
	)
	table.AppendBulk(data)
	table.Render()
}
