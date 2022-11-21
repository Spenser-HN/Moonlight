package cli

import (
	"encoding/json"
	"io"
	"log"
	"os"

	syntax "com.moonlight/app/src/Controllers"
	filereader "com.moonlight/app/src/FileReader"
	"github.com/spf13/cobra"
)

var Generate_keys = &cobra.Command{
	Use:     "get-keys",
	Aliases: []string{"generate-keys", "keys"},
	Short:   "get keys from typescript file(s)",
	//Long:
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			log.Fatal("Not enough arguments")
		}

		Readed, filesPaths := filereader.ReadFile(args)

		Tokens := syntax.GetTokens(Readed, filesPaths)

		JSON, err := json.Marshal(Tokens)

		if err != nil {
			log.Fatal(err)
		}

		if _, err := io.WriteString(os.Stdout, string(JSON)); err != nil {
			log.Fatal(err)
		}

	},
}
