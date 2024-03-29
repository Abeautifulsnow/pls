package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "pls version: [v0.1.8]"
)

func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Prints the version of pls",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println(version)
		},
	}
	return cmd
}
