package main

import (
	"encoding/xml"
	"fmt"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

type UnitElmt struct {
	Value string `xml:",chardata"`
	Unit  string `xml:"units,attr"`
}

type Node struct {
	Id          string     `xml:"id,attr"`
	Class       string     `xml:"class,attr"`
	Version     string     `xml:"version"`
	Clock       []UnitElmt `xml:"clock"`
	Size        []UnitElmt `xml:"size"`
	Width       []UnitElmt `xml:"width"`
	Disabled    bool       `xml:"disabled,attr"`
	Description string     `xml:"description"`
	Product     string     `xml:"product"`
	Vendor      string     `xml:"vendor"`
	Serial      string     `xml:"serial"`
	BusInfo     string     `xml:"businfo"`
	ChildNodes  []Node     `xml:"node"`
}

type Query struct {
	Nodes []Node `xml:"node>node>node"`
}

func (u UnitElmt) String() string {
	return fmt.Sprintf("%s|%s", u.Value, u.Unit)
}

func (u UnitElmt) toUnitValue() UnitValue {
	value, err := strconv.Atoi(u.Value)
	if err != nil {
		panic(err)
	}
	return UnitValue{Unit: u.Unit, Value: value}
}

func (n Node) String() string {
	return fmt.Sprintf("\n%s.%s { Version: %s,\n Clock: %s,\n Size: %s,\n Width: %s,\n Description: %s,\n Product: %s,\n Vendor: %s,\n Serial: %s,\n Businfo: %s,\n Disabled: %t }\n",
		n.Id,
		n.Class,
		n.Version,
		n.Clock,
		n.Size,
		n.Width,
		n.Description,
		n.Product,
		n.Vendor,
		n.Serial,
		n.BusInfo,
		n.Disabled,
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

func FilterMatchingNodes(n []Node, p string) []Node {
	nc := make([]Node, 0)
	r, _ := regexp.Compile(p)
	for _, node := range n {
		if r.MatchString(node.Id) {
			nc = append(nc, node)
		}
	}
	return nc
}

func FilterDisabledNodes(n []Node) []Node {
	var nc []Node
	for _, node := range n {
		if !node.Disabled {
			nc = append(nc, node)
		}
	}
	return nc
}

func SaveServerFromXML(path, serverName, dbName string) {
	db, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	xmlFile, err := os.Open(fmt.Sprintf("%s/%s/%s-ALL-XML.xml", path, serverName, serverName))
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)

	var q Query
	xml.Unmarshal(b, &q)
	s := GenerateServerFromNodes(serverName, q.Nodes)
	saveServer(s, db, dbName)
}
