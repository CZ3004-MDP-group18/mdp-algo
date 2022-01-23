package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteCSV(name string, records [][]string) {
	file, err := os.Create(name)
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	w.WriteAll(records)
}
