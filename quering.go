// quering
package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func dbstart() *gocql.Session {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.ProtoVersion = 4
	cluster.Keyspace = "sampledata"
	cluster.Consistency = gocql.One
	cass, err := cluster.CreateSession()
	if err != nil {
		panic(fmt.Sprintf("Error creating session:%v", err))
	}
	return cass
}
func dbend(cass *gocql.Session) {
	cass.Close()
}
func insertintodb(cass *gocql.Session, sub Clientdata) {
	if err := cass.Query("INSERT INTO sampleclient (id,email,first,last) VALUES (?,?,?,?)", sub.ID, sub.Email, sub.First, sub.Last).Exec(); err != nil {
		log.Fatal(err)
	}
}
func readfromdb(cass *gocql.Session, id int) {
	sub := Clientdata{}
	iter := cass.Query("SELECT id,first,last,email FROM sampleclient WHERE id=?", id).Iter()
	for iter.Scan(&sub.ID, &sub.First, &sub.Last, &sub.Email) {
		log.Println(sub.First)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}
