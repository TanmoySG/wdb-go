package schema

import (
	"encoding/json"
	"fmt"

	"github.com/TanmoySG/wunderDB/pkg/fs"
)

type CollectionSchema map[string]interface{}

func isJson(content []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(content, &js) == nil
}

func LoadSchemaFromFile(filepath string) (CollectionSchema, error) {
	if !fs.CheckFileExists(filepath) {
		return nil, fmt.Errorf("file not found: %s", filepath)
	}

	fileContentBytes, err := fs.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %s", filepath, err)
	}

	if !isJson(fileContentBytes) {
		return nil, fmt.Errorf("file %s, content is not JSON", filepath)
	}

	var cs CollectionSchema
	err = json.Unmarshal(fileContentBytes, &cs)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling file %s: %s", filepath, err)
	}

	return cs, nil
}
