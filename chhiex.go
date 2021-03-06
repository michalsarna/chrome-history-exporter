package main

import (
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func writeToCsvFile(data [][]string, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Cannot create file ", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	//change delimiter for csv fields from ","
	writer.Comma = '|'
	writer.WriteAll(data)
	if err := writer.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

func convertTime(sourceTime int) string {
	var convertedTime int64
	convertedTime = int64(((sourceTime / 1000000) - 11644473600))
	tm := time.Unix(convertedTime, 0)
	return tm.String()
}

func getHistory(dbPtr string, exportToFile bool, outputFile string) {
	db, err := sql.Open("sqlite3", dbPtr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var sqlQuerry = "select urls.id, urls.title, urls.url, urls.last_visit_time, urls.visit_count from urls order by urls.id;"

	rows, err := db.Query(sqlQuerry)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var ansTable [][]string
	var iterator int

	ansTable = append(ansTable, []string{"ID", "Site Title", "Site URL", "Last Visited", "Visits Count"})

	for rows.Next() {
		var id int
		var title string
		var url string
		var lastVisitTime int
		var visitCount int
		rows.Scan(&id, &title, &url, &lastVisitTime, &visitCount)
		rowToAdd := []string{strconv.Itoa(id), title, url, convertTime(lastVisitTime), strconv.Itoa(visitCount)}
		ansTable = append(ansTable, rowToAdd)
		iterator++
	}
	if exportToFile {
		writeToCsvFile(ansTable, outputFile)
	} else {
		for i := 0; i < iterator; i++ {
			fmt.Println(ansTable[i][0], ansTable[i][1], ansTable[i][2], ansTable[i][3], ansTable[i][4])
			//printing with some additional formating
		}
	}

}

func main() {
	dbFile := flag.String("in-file", "./History", "History file to read from (SQLite format)")
	exportToFile := flag.Bool("export-to-file", false, "false - don't export; true - export to file")
	outFilePtr := flag.String("out-file", "./export.csv", "file to export data to (CSV format)")
	flag.Parse()
	fmt.Println("Input DB File Name: ", *dbFile)
	fmt.Println("Export to file ? ", *exportToFile)
	fmt.Println("Output File Name: ", *outFilePtr)
	//add options to choose what data from history file to export
	getHistory(*dbFile, *exportToFile, *outFilePtr)
}
