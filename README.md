# wdb-go, Go Client for wunderDB

wdb-go is a Go Client library for wunderDB.

## Install

To use wdb-go, install it using `go get`

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
err := wdb.CreateUser(username, password)
```

It returns error if no user was created, else returns nil error.

### Create Role

To create a role, use `CreateRole()` method - passing the `name` of role to create, and the lists of Allowed and Denied Privileges.

```go
err := wdb.CreateRole(roleName, allowedPrivileges, deniedPrivileges)
```

Use the privileges available in the `github.com/TanmoySG/wdb-go/privileges` sub-package as `privileges.PrivilegeName`, refer to [this for more](./README.md#sub-packages)

### Grant Role

Once user and role are created use the `GrantRole()` method to grant the role to the user - passing the username, role-name and database to grant the role on. In addition to the database, a role can also be granted on a collection by passing the collection name as the last argument, which is an option argument.

```go
// role granted on collection
err := wdb.GrantRole(username, roleName, database, collection)

// collection is an optional argument, role granted only on database
err := wdb.GrantRole(username, roleName, database) 
```

The collection has to be a child of the database, if the role needs to be granted on the collection.

### List Roles

Use the `ListRoles()` method to list the roles in wunderdb. Returns map of roles and error.

```go
roles, err := wdb.ListRoles() 
```

### Create Database

Create Databases using the `CreateDatabase()` function passing the name of the database to create.

```go
err := wdb.CreateDatabase(databaseName)
```

### Get Database

Fetch database details using `GetDatabase()` method.

```go
databases, err := wdb.GetDatabase(databaseName)
```

Returns the database, of type [`*wdbModels.Database`](https://github.com/TanmoySG/wunderDB/blob/main/model/models.go#L29) or error, if any.

### Delete Database

Delete a Database using `DeleteDatabase()` method passing the database to delete.

```go
err := wdb.DeleteDatabase(databaseName)
```

Returns error, if any.

### Create Collection

Create Collections, in an existing Database, using the `CreateCollection()` function - passing the name of the database to create collection in, the name of the collection to create and the collection schema.

```go
err := wdb.CreateCollection(databaseName, collectionName, collectionSchema)
```

The collection schema passed is of type [`schema.CollectionSchema`](./schema/schema.go). Schemas can be loaded into your code using the methods available in the `schema` subpackage, read more [here](#wdb-goschema).

### Get Collection

Fetch Collection details, in a database, using `GetCollection()` method.

```go
collection, err := wdb.GetCollection(databaseName, collectionName)
```

Returns the collection, of type [`*wdbModels.Collection`](https://github.com/TanmoySG/wunderDB/blob/main/model/models.go#L35) or error, if any.

### Delete Collection

Delete a collection using `DeleteCollection()` method passing the collection to delete, and it's parent database.

```go
err := wdb.DeleteCollection(databaseName, collectionName)
```

Returns error, if any.

### Add/Insert Data

Insert Data in a Collection using the `AddData()` function passing the data to be inserted, followed by target database and collection name.

```go
err := wdb.AddData(data, databaseName, collectionName)
```

The `data` parameter is of type `interface` and can have any value. Data to insert must be schema validated. Returns error if any.

### Read Data

To read/fetch data from a Collection, use the `ReadData()` function passing the database and collection name.

```go
data, err := wdb.ReadData(databaseName, collectionName, filters)
```

The `filter` parameter is of type `dataFilters.Filter` and filters can be used to filter out data based on key-value. The methods in [`filters`](#wdb-gofilters) subpackage help in using filters in the host application.

Read more about wunderDB filters [here](https://github.com/TanmoySG/wunderDB/blob/main/documentation/README.md#filters)

### Update Data

To update specific data from a Collection, use the `UpdateData()` function passing the updates to the data, database and collection name along with filters to specify the data to update.

```go
data, err := wdb.UpdateData(updatedFields, databaseName, collectionName, filters)
```

The updated data can be part or entire data object. The updated data must be schema validated.

### Delete Data

To delete specific data from a Collection, use the `DeleteData()` function passing the database and collection name along with filters to specify the data to delete.

```go
err := wdb.DeleteData(databaseName, collectionName, filters)
```

Returns error, if any.

## Sub Packages

Subpackages in wdb-go are useful for development.

### wdb-go/privileges

Use the `privileges` object from the `wdb-go/privileges` sub-package to use the various wdb privileges available.

```go
import privileges "github.com/TanmoySG/wdb-go/privileges"

allowedPrivileges := []privileges.Privilege{
    privileges.ReadDatabase,
}
```

### wdb-go/filters

Use the `filters` object from the `wdb-go/filters` sub-package to use the filters functionalities of wunderDB.

```go
import filters "github.com/TanmoySG/wdb-go/filters"

filter, err := filters.GetFilter("field", "value")

isFilterValid := filter.IsValid()
```

### wdb-go/schema

Use the `schema` object from the `wdb-go/schema` sub-package to use the schema functionalities of wdg-go tool.

```go
import schema "github.com/TanmoySG/wdb-go/schema"

// load schema from json file, passing filepath to JSON Schema file
loadedCollectionSchema, err := schema.LoadSchemaFromFile(filepath) 
```
