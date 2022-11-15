package main

import (
	syntax "com.moonlight/app/src/Controllers"
	filereader "com.moonlight/app/src/FileReader"
)

func main() {
	result := filereader.ReadFile([]string{"./index.go", "./src/Controllers/ConvertTokens.go"})
	syntax.GetTokens(result)
}
