package main

import (
	"strings"

	log "github.com/sirupsen/logrus"

	wdbgo "github.com/TanmoySG/wdb-go"
	"github.com/TanmoySG/wdb-go/privileges"
)

func main() {

	uname, pword := "admin", "admin"
	appnme := "example-wdb-go-app"
	wdbAddress := "http://localhost:8086"

	// create client
	wdb := wdbgo.NewWdbClient(uname, pword, wdbAddress, &appnme)

	// login users
	resp, err := wdb.LoginUser(uname, pword)
	if err != nil {
		log.Error(err)
	} else {
		log.Info(resp)
	}

	// create users
	err = wdb.CreateUser(uname, pword)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("created user")
	}

	// privileges
	allowed := []privileges.Privilege{
		privileges.AddData,
		privileges.CreateCollection,
	}

	denied := []privileges.Privilege{
		privileges.DeleteCollection,
	}

	err = wdb.CreateRole("xyz", allowed, denied)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("created role")
	}

	// err = wdb.GrantRoles(uname, "xyz", "databadse")
	// if err != nil {
	// 	log.Error(err)
	// } else {
	// 	log.Info("granted role")
	// }

	rolesList, err := wdb.ListRoles()
	if err != nil {
		log.Error(err)
	} else {
		res := []string{}
		for roleName := range rolesList {
			res = append(res, roleName)
		}
		log.Infof("[ %s ]", strings.Join(res, " , "))
	}

	err = wdb.CreateDatabase("test-database")
	if err != nil {
		log.Error(err)
	} else {
		log.Info("created db")
	}

	db, err := wdb.GetDatabase("test-database")
	if err != nil {
		log.Error(err)
	} else {
		log.Infof("Collections [ %s ]", db.Collections)
	}

	err = wdb.DeleteDatabase("databadse")
	if err != nil {
		log.Error(err)
	} else {
		log.Info("database deleted")
	}

}

//databadse
