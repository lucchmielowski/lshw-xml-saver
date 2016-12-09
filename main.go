package main

import (
	"io/ioutil"
)

func main() {
	files, _ := ioutil.ReadDir("files")

	for _, f := range files {
		SaveServerFromXML(f.Name())
	}

}
