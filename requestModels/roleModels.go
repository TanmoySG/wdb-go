package requestModels

import wdbModels "github.com/TanmoySG/wunderDB/model"

type GrantRoles struct {
	Username   string                `json:"username" xml:"username"`
	Permission wdbModels.Permissions `json:"permissions" xml:"permissions"`
}
