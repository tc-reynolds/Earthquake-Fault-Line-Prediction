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
	fmt.Println("Loading database...")
	csvFile, err := os.Open("./data/database.csv")
	fmt.Println("Loaded...")
	if err != nil {
		log.Fatal("Couldn't open the csv file", err)
	}
	// Parse the file
	r := csv.NewReader(csvFile)
	fmt.Println("New Reader init...")

	records, err := r.ReadAll()
	fmt.Println("Array created...")
	fmt.Println("records[1][2]")
	fmt.Println(records[0][2])
	fmt.Println(records[0][3])

	return records
}

