package privileges

type Privilege struct {
	privilege   string
	description string
}

// TODO: add description

var (
	CreateUser = createPrivilege("createUser", "")
	GrantRole  = createPrivilege("grantRole", "")
	LoginUser  = createPrivilege("loginUser", "")

	CreateRole = createPrivilege("createRole", "")
	UpdateRole = createPrivilege("updateRole", "")
	ListRole   = createPrivilege("listRole", "")

	CreateDatabase = createPrivilege("createDatabase", "")
	ReadDatabase   = createPrivilege("readDatabase", "")
	UpdateDatabase = createPrivilege("updateDatabase", "")
	DeleteDatabase = createPrivilege("deleteDatabase", "")

	CreateCollection = createPrivilege("createCollection", "")
	ReadCollection   = createPrivilege("readCollection", "")
	UpdateCollection = createPrivilege("updateCollection", "")
	DeleteCollection = createPrivilege("deleteCollection", "")

	AddData    = createPrivilege("addData", "")
	ReadData   = createPrivilege("readData", "")
	UpdateData = createPrivilege("updateData", "")
	DeleteData = createPrivilege("deleteData", "")
)

func createPrivilege(privilege, description string) Privilege {
	return Privilege{
		privilege:   privilege,
		description: description,
	}
}

func (p Privilege) Name() string {
	return p.privilege
}

