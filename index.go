package main

import (
	"fmt"

	syntax "com.moonlight/app/src/Controllers"
	filereader "com.moonlight/app/src/FileReader"
)

func main() {
	result := filereader.ReadFile([]string{"./src/Test/index.ts"})
	fmt.Println(syntax.GetTokens(result))
}
