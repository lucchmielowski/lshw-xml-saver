package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Node struct {
	Id          string `xml:"id,attr"`
	Class       string `xml:"class,attr"`
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
	return fmt.Sprintf("%s.%s { %s, %s, %s, %s, %s, %s, %s, %s }\n",
		n.Id,
		n.Class,
		n.Version,
		n.Clock,
		n.Size,
		n.Description,
		n.Product,
		n.Vendor,
		n.Serial,
		n.BusInfo,
	)
}

func FindNodesByClass(n []Node, c string, r *[]Node) {
	for i := 0; i < len(n); i++ {
		elmt := n[i]
		if elmt.Class == c {
			*r = append(*r, elmt)
		}
		FindNodesByClass(elmt.ChildNodes, c, r)
	}
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

func CreateServerFromXML(serverName string) {
	xmlFile, err := os.Open(fmt.Sprintf("files/%s/%s-ALL-XML.xml", serverName, serverName))
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)

	var q Query
	xml.Unmarshal(b, &q)
	var r []Node
	FindNodesByClass(q.Nodes, "memory", &r)
	fmt.Println(r)
}
