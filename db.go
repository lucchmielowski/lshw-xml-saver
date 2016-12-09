package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

func getAllServers(s *mgo.Session) {

}

func saveServer(serv Server, s *mgo.Session, dbName string) {
	session := s.Copy()
	defer session.Close()

	c := session.DB(dbName).C("servers")

	err := c.Insert(serv)
	if err != nil {
		if mgo.IsDup(err) {
			fmt.Println("This server already exists")
		}

		fmt.Println("Failed saving the server: DB Error")
	}
}
