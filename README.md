# wdb-go, Go Client for wunderDB

wdb-go is a Go Client library for wunderDB.

## Install

To use wdb-go, install it using `go get`.

```sh
go get github.com/TanmoySG/wdb-go
```

## Usage

Then in your go code, import it as

```go
import wdbgo "github.com/TanmoySG/wdb-go"
```

Initialize a new wdb client using `NewClient()` method passing the URL of wdb instance, username and password of the authenticating user.
```go
wdb, err := wdbgo.NewClient(uname, pword, wdbAddress, nil)
if err != nil {
 log.Fatal(err)
}
```

You an also pass in a custom application name to be included as the user-agent. If you do not need app name you can pass `nil`, as above. `NewWdbClient` checks if the connection can be eshtablished, otherwise returns error. To skip the first time check you can pass `wdbgo.SkipConnectionCheck` as the last argument.

```go
wdb, err := wdbgo.NewWdbClient(uname, pword, wdbAddress, nil, wdbgo.SkipConnectionCheck)
```

### Create User

To crete a user, use the `CreateUser()` function.
```go
err := CreateUser(username, password)
```

It returns error if no user was created, else returns nil error.

### Create Role

To create a role, use `CreateRole()` method - passing the `name` of role to create, and the lists of Allowed and Denied Privileges. 

```go
err := CreateRole(roleName string, allowedPrivileges, deniedPrivileges []privileges.Privilege)
```

Use the privileges available in the `github.com/TanmoySG/wdb-go/privileges` sub-package as `privileges.PrivilegeName`.

```go
import privileges "github.com/TanmoySG/wdb-go/privileges"

allowedPrivileges := []privileges.Privilege{
    privileges.ReadDatabase,
}
```