package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version string = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "moonlight",
	Version: version,
	Short:   "moonlight - a simple CLI to compile typescript (.ts) code into native code",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`
		
------ MOONLIGHT -----

Description:
%s

version(%s)

		`,
			"A simple CLI to compile typescript (.ts) code into native code",
			version)
	},
}

func Init() {

	//Adding Commands
	rootCmd.AddCommand(Create_proyect) //Beta
	rootCmd.AddCommand(Generate_keys)  //Done

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
