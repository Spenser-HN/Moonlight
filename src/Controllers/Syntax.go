package syntax

import (
	filereader "com.moonlight/app/src/FileReader"
)

type Tokens map[string]string
type Tokenized map[int64]Converted
type FileTokenized map[int8]Tokens

func GetTokens(ReadedFiles filereader.ReadedFile) Tokenized {

	var _Tokens Tokenized = Tokenized{}

	for file := range ReadedFiles {

		for line := range ReadedFiles[int8(file)] {

			_Tokens[int64(line)] = ConvertTotoken(ReadedFiles[int8(file)][int32(line)], line)

			delete(ReadedFiles[int8(file)], int32(line))

		}

		delete(ReadedFiles, int8(file))

	}

	return _Tokens

}
