package util

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	//"io"
	"log"
	"os"
)

func LoadDatabase() [][]string{
	// Open the file
	csvFile, err := os.Open("./data/database.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	fmt.Printf("records[1][2]")
	fmt.Printf(records[1][2])
	return records
}

