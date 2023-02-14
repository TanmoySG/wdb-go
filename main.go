package main

import (
	"fmt"
	"log"

	"github.com/TanmoySG/wdb-go/internal/queries"
	"github.com/TanmoySG/wdb-go/internal/queries/methods"
	"github.com/TanmoySG/wdb-go/internal/routes"
	// f "github.com/TanmoySG/wunderDB/pkg/fs"
)

func main() {

	rt := routes.ApiPing.Format("http://localhost:8089")

	qc := queries.NewQueryClient("admin", "admin", "someone")
	re, err := qc.Query(rt, string(methods.ApiPing), nil)
	if err != nil {
		log.Fatal(err)
	}

	r, err := re.ApiResponse()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r.Action, r.Data)

}
