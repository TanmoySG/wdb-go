package main

import (
	"os"

	wdbgo "github.com/TanmoySG/wdb-go"
	log "github.com/sirupsen/logrus"
)

func main() {
	uname, pword := os.Getenv("ADMIN_ID"), os.Getenv("ADMIN_PASSWORD")
	appName := "example-wdb-go-app"
	wdbAddress := os.Getenv("WDB_URL")

	// skip connection check by adding wdbgo.SkipConnectionCheck as argument
	wdb, err := wdbgo.NewClient(uname, pword, wdbAddress, &appName)
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]interface{}{
		"age": 19,
	}

	err = wdb.AddData(data, "test-database", "collection-1")
	if err != nil {
		log.Error(err)
	} else {
		log.Infof("data inserted")
	}

}
