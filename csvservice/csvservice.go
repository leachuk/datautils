package csvservice

import (
	"fmt"
	"encoding/csv"
	"io"
	"log"
	"os"
	"bufio"
	"path/filepath"
)

type CsvData struct {
	Heading	[]string
	Row		[]string
}

func ReadCSV(path string) CsvData {
	reader := csvReader(path)

	var data CsvData
	for i := 0; ; i++ {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		fmt.Printf("[%d]%s\n", i, line)
		if i == 0 {
			for j, header := range line {
				fmt.Printf("		heading[%d]%s\n",j, header)
				data.Heading = append(data.Heading, header)
			}
		} else {
			for _, row := range line {
				data.Row = append(data.Row, row)
			}
		}
	}

	return data
}

func csvReader(path string) (s *csv.Reader) {
	rootPath, _ := filepath.Abs(".")
	fmt.Println(rootPath)

	csvFile, _ := os.Open(rootPath + path)

	return csv.NewReader(bufio.NewReader(csvFile))
}
