package main

import (
	"fmt"
	"log"

	wdbgo "github.com/TanmoySG/wdb-go"
	"github.com/TanmoySG/wdb-go/privileges"
)

// "log"

// "github.com/TanmoySG/wdb-go/internal/methods"
// "github.com/TanmoySG/wdb-go/internal/queries"
// "github.com/TanmoySG/wdb-go/internal/routes"
// f "github.com/TanmoySG/wunderDB/pkg/fs"

func main() {

	// rt := routes.ApiPing.Format("http://localhost:8089")

	// qc := queries.NewQueryClient("admin", "admin", "someone")
	// re, err := qc.Query(rt, string(methods.ApiPing), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// r, err := re.ApiResponse()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(r.Action, *r.Data)

	u, p := "admin", "admin"

	appnme := "hyttt"

	g := wdbgo.NewWdbClient(u, p, "http://localhost:8086", &appnme)

	// fmt.Print(g)

	// r, err := g.LoginUser(u, p)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// r, err := g.CreateUser(u, p)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	allowed := []privileges.Privilege{
		privileges.AddData,
		privileges.CreateCollection,
	}

	denied := []privileges.Privilege{
		privileges.DeleteCollection,
	}

	r, err := g.CreateRole("xyz", allowed, denied)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(r)

}
