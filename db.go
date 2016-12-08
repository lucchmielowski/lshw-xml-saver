package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

const DB_NAME = "xml-test"

func getAllServers(s *mgo.Session) {

}

func findServerById(id string, s *mgo.Session) {
	session := s.Copy()
	defer session.Close()
}

func saveServer(serv Server, s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB(DB_NAME).C("servers")

	err := c.Insert(serv)
	if err != nil {
		if mgo.IsDup(err) {
			fmt.Println("This server already exists")
		}

		fmt.Println("Failed saving the server: DB Error")
	}
}
