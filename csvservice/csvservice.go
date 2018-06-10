package csvservice

import (
	"fmt"
	"encoding/csv"
	"io"
	"log"
	"os"
	"bufio"
	"path/filepath"
	"flag"
)

type CsvData struct {
	Heading	[]string
	Data map[string]string
}

func ReadCSV(path string) []CsvData {
	isDebug := flag.Lookup("debug").Value.(flag.Getter).Get().(bool)

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

		if i == 0 {
			for _, header := range line {
				heading = append(heading, header)
			}
		} else {
			datamap = make(map[string]string)
			for k, row := range line {
				if isDebug {
					fmt.Printf("(%v)[heading:%s]data[%s]\n", k, heading[k], row)
				}
				datamap[heading[k]] = row
			}

			data = append(data, CsvData{heading,datamap})
		}

	}

	return data
}

func csvReader(path string) (s *csv.Reader) {
	isDebug := flag.Lookup("debug").Value.(flag.Getter).Get().(bool)

	rootPath, _ := filepath.Abs(".")

	if isDebug {
		fmt.Println("csvReader file path:" + rootPath + "/" + path)
	}
	csvFile, _ := os.Open(rootPath + "/" + path)

	return csv.NewReader(bufio.NewReader(csvFile))
}

func WriteCSV(path string, data [][]string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}
}
