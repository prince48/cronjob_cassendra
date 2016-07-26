package main

import (
	"fmt"
	"gocron"
	"time"

	"github.com/gocql/gocql"
)

func cronjob(cass *gocql.Session) {
	ch := gocron.Start()
	go test(ch)
	for i := 0; i < 8; i++ {
		gocron.Every(5).Seconds().Do(readfromdb, cass, i)
	}
	<-ch
}
func test(stop chan bool) {
	time.Sleep(24 * time.Second)
	gocron.Clear()
	fmt.Println("Job Done.")
	close(stop)
}
