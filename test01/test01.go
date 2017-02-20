package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// https://github.com/GoesToEleven/golang-web-dev/tree/master/012_hands-on/09_hands-on

func main() {
	tblF, err := os.Open("table.csv")
	if err != nil {
		log.Println("Error opening file", err)
	}
	defer tblF.Close()

	tbl := csv.NewReader(tblF)
	tbl.Comma = ','
	lineCount := 0

	for {
		// read just one record, but we could ReadAll() as well
		record, err := tbl.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// record is an array of string so is directly printable
		fmt.Println("Record", lineCount, "is", record, "and has", len(record), "fields")
		// and we can iterate on top of that
		for i := 0; i < len(record); i++ {
			fmt.Println(" ", record[i])
		}
		fmt.Println()
		lineCount += 1
	}
}
