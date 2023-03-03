package requestModels

import "github.com/TanmoySG/wdb-go/schema"

type CreateCollection struct {
	Name   string                  `json:"name" xml:"name"`
	Schema schema.CollectionSchema `json:"schema" xml:"schema"`
}
