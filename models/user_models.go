package models

type CreateUser struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

type CreateRole struct {
	Role    string   `json:"role" xml:"role" form:"role"`
	Allowed []string `json:"allowed" xml:"allowed" form:"allowed"`
	Denied  []string `json:"denied" xml:"denied" form:"denied"`
}
