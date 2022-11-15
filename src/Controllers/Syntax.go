package syntax

import (
	filereader "com.moonlight/app/src/FileReader"
)

type Tokens map[string]string
type FileTokenized map[int8]Tokens

func GetTokens(ReadedFiles filereader.ReadedFile) {

	for file := range ReadedFiles {

		for line := range ReadedFiles[int8(file)] {

			ConvertTotoken(ReadedFiles[int8(file)][int32(line)], line)

			delete(ReadedFiles[int8(file)], int32(line))

		}

		delete(ReadedFiles, int8(file))

	}

}
