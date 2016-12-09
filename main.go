package main

import (
	"flag"
	"io/ioutil"
)

func main() {

	directoryPtr := flag.String("dir", "./files", "Path to the directory containing all servers sub-directories")
	dbPtr := flag.String("db", "lshw-xml", "Database name")
	flag.Parse()

	files, _ := ioutil.ReadDir(*directoryPtr)

	for _, f := range files {
		SaveServerFromXML(*directoryPtr, f.Name(), *dbPtr)
	}

}
