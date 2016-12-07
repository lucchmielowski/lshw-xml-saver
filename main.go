package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Node struct {
	Id          string `xml:"id,attr"`
	Version     string `xml:"version"`
	Clock       string `xml:"clock"`
	Size        string `xml:"size"`
	Description string `xml:"description"`
	Product     string `xml:"product"`
	Vendor      string `xml:"vendor"`
	Serial      string `xml:"serial"`
	BusInfo     string `xml:"businfo"`
	ChildNodes  []Node `xml:"node"`
}

type Query struct {
	Nodes []Node `xml:"node>node>node"`
}

func (n Node) String() string {
	return fmt.Sprintf("%s { %s, %s, %s, %s, %s, %s, %s, %s }\n",
		n.Id,
		n.Version,
		n.Clock,
		n.Size,
		n.Description,
		n.Product,
		n.Vendor,
		n.Serial,
		n.BusInfo)
}

func FindNodeById(n []Node, id string) Node {
	for i := 0; i < len(n); i++ {
		elmt := n[i]
		if elmt.Id == id {
			return elmt
		}
		value := FindNodeById(elmt.ChildNodes, id)
		if value.Id != "" {
			return value
		}
	}
	return Node{}
}

func main() {
	xmlFile, err := os.Open("files/Serveur-666/Serveur-666-ALL-XML.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)

	var q Query
	xml.Unmarshal(b, &q)
	fmt.Println(FindNodeById(q.Nodes, "volume:1"))
}
