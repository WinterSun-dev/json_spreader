package writecsv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func WriteCsv(data [][]string, destination string) (string, error) {

	csvFile, err := os.Create(destination)

	if err != nil {
		return "", err
	}
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)

	for _, record := range data {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)

	}
	return fmt.Sprintf("%v atributes, %v data rows", len(data[0]), len(data)-1), nil
}
