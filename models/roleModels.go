package models

import "github.com/TanmoySG/wunderDB/model"

type GrantRoles struct {
	Username   string            `json:"username" xml:"username"`
	Permission model.Permissions `json:"permissions" xml:"permissions"`
}