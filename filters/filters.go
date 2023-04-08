package filters

import (
	"fmt"
)

type Filter struct {
	Key   string
	Value interface{}
}

func GetFilter(key string, value interface{}) (*Filter, error) {
	if key == "" || value == nil {
		return nil, fmt.Errorf("key/value missing for filter")
	}

	return &Filter{
		Key:   key,
		Value: value,
	}, nil
}

func (f Filter) IsValid() bool {
	return !(f.Key == "" || f.Value == nil)
}
