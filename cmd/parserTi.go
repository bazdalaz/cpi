package cmd

import (
	"fmt"
	"github.com/bazdalaz/cpi/pkg/parser/ti"
	"github.com/spf13/cobra"
)

// parseTiCmd represents the command to parse TDC files
var parseTiCmd = &cobra.Command{
	Use:   "ti",
	Short: "Parses TI controller layout files",
	Long:  `Parses TI (FSS) layout files.`,
	Run: func(cmd *cobra.Command, args []string) {

		parser := ti.New()
		err := parser.ParseDir()
		if err != nil {
			fmt.Printf("Error parsing FSS file: %v\n", err)
			return
		}
	},
}

func init() {
	// If it's a subcommand of parse, add it like this:
	parseCmd.AddCommand(parseTiCmd)

	// If it's a standalone command, you would add it to the root command:
	// rootCmd.AddCommand(parseTdcCmd)
}
