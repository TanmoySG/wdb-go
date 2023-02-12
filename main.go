package main

import (
	"fmt"

	"github.com/TanmoySG/wdb-go/internal/routes"
	// f "github.com/TanmoySG/wunderDB/pkg/fs"
)

func main() {
	// f.CreateFile("./test.txt")
	r := routes.LoginUser.Format("test", nil)
	// r2 := routes.CreateUser.Format("test")
	fmt.Println(r)
	// fmt.Println(r2)

}
