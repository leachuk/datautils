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
	Data map[string]string
}

func ReadCSV(path string) []CsvData {
	reader := csvReader(path)

	var data []CsvData
	var heading []string
	var datamap map[string]string
	
	for i := 0; ; i++ {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		//fmt.Printf("[%d]%s\n", i, line)
		if i == 0 {
			for j, header := range line {
				fmt.Printf("heading[%d]%s\n",j, header)
				heading = append(heading, header)
			}
		} else {
			datamap = make(map[string]string)
			for k, row := range line {
				fmt.Printf("(%v)[heading:%s]data[%s]\n",k,heading[k],row)
				datamap[heading[k]] = row
			}

			data = append(data, CsvData{heading,datamap})
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
