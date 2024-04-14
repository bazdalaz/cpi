package cmd

import (
	"fmt"
	"github.com/bazdalaz/cpi/pkg/parser/tdc"
	"github.com/spf13/cobra"
)

// parseTdcCmd represents the command to parse TDC files
var parseTdcCmd = &cobra.Command{
	Use:   "tdc",
	Short: "Parses TDC controller layout files",
	Long:  `Parses Theoretical Data Controller (TDC) layout files.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: No TDC file path provided")
			return
		}
		filePath := args[0]

		parser := tdc.New()
		err := parser.ParseFile(filePath)
		if err != nil {
			fmt.Printf("Error parsing TDC file: %v\n", err)
			return
		}
	},
}

func init() {
	// If it's a subcommand of parse, add it like this:
	parseCmd.AddCommand(parseTdcCmd)

	// If it's a standalone command, you would add it to the root command:
	// rootCmd.AddCommand(parseTdcCmd)
}
