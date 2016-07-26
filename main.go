// quering

package main

//"fmt"
//"log"
//"github.com/gocql/gocql"

type Clientdata struct {
	ID    int
	First string
	Last  string
	Email string
}

func main() {
	sub := Clientdata{7, "kamal@gmail.com", "kamal", "giri"}
	cass := dbstart()
	insertintodb(cass, sub)
	cronjob(cass)
	dbend(cass)
}
