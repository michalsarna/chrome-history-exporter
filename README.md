# chrome history exporter

It's a small GO program that can be used to get internet history from Chorome web browser.

## build
    go get github.com/mattn/go-sqlite3
    go build chhiex.go

## run
    Usage of ./chhiex:
        -export-to-file
    	   false - don't export; true - export to file
        -in-file string
    	   History file to read from (SQLite format) (default "./History")
        -out-file string
    	   file to export data to (CSV format) (default "./export.csv")

It's still work in progress and my first GO program. 
