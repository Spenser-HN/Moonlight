package filereader

import (
	"bufio"
	"log"
	"os"
)

func ReadLines(FilePath string) map[int32]string {

	// open file
	f, err := os.Open(FilePath)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//creating an empty array

	var Line map[int32]string = make(map[int32]string)
	ActualLine := 0

	for scanner.Scan() {

		ActualLine++
		Line[int32(ActualLine)] = string(scanner.Text())

	}

	return Line

}
