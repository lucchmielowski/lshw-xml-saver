package main

import (
	"flag"
	"io/ioutil"
)

func main() {

	directoryPtr := flag.String("dir", "./files", "Path du dossier contenant tout les sous-dossiers de serveurs")
	flag.Parse()

	files, _ := ioutil.ReadDir(*directoryPtr)

	for _, f := range files {
		SaveServerFromXML(*directoryPtr, f.Name())
	}

}
