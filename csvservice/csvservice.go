package main

import (
	"fmt"
	"encoding/csv"
	"io"
	"log"
	"os"
	"bufio"
	"path/filepath"
)

func csvReader() (s *csv.Reader) {
	rootPath, _ := filepath.Abs(".")
	fmt.Println(rootPath)

	csvFile, _ := os.Open(rootPath + "/csvservice/datasample.csv")

	return csv.NewReader(bufio.NewReader(csvFile))
}

func main() {
	fmt.Println("csvservice main")

	reader := csvReader()
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		fmt.Println(line)
	}
}