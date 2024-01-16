package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var version = "0.0.1"
var RootCmd = &cobra.Command{
	Use:     "ms",
	Version: version,
	Short:   "marketsync - otto marketplace upload",
	Long: `otto marketplace upload 
   
Sync your local inventory with otto's marketplace'`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
