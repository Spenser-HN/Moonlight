package syntax

import (
	filereader "com.moonlight/app/src/FileReader"
)

type Tokens map[string]string
type Tokenized map[int64]Converted
type FileTokenized map[int8]Tokens
type TokenizedFiles map[string]Tokenized

func GetTokens(ReadedFiles filereader.ReadedFile, Files []string) TokenizedFiles {

	var _Tokens Tokenized = make(Tokenized)

	var _TokenizedFiles TokenizedFiles = make(TokenizedFiles)

	for file := range ReadedFiles {

		for line := range ReadedFiles[int8(file)] {

			_Tokens[int64(line)] = ConvertTotoken(ReadedFiles[int8(file)][int32(line)], line)

			delete(ReadedFiles[int8(file)], int32(line))

		}

		_TokenizedFiles[Files[file]] = _Tokens
		delete(ReadedFiles, int8(file))
		_Tokens = make(Tokenized)
	}

	return _TokenizedFiles

}
