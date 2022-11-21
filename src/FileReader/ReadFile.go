package filereader

type ReadedFile map[int8]map[int32]string

func ReadFile(Files_Path []string) (ReadedFile, []string) {

	var Readed ReadedFile = make(ReadedFile)

	for read := 0; read < len(Files_Path); read++ {

		//Reading file by file
		Readed[int8(read)] = ReadLines(Files_Path[read])

	}

	return Readed, Files_Path
}
