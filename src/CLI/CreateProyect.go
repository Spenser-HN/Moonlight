package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Create_proyect = &cobra.Command{
	Use:     "create-proyect",
	Aliases: []string{"create", "new"},
	Short:   "create a new moonlight proyect",
	//Long:
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Args:", args)
	},
}
