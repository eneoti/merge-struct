package merp

import (
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"reflect"
)

// MergeOverwrite overwrite map
func MergeOverwrite(to, from, dst interface{}) error {
	var toMap map[string]interface{}
	var fromMap map[string]interface{}
	kindTo := reflect.ValueOf(to)
	kindFrom := reflect.ValueOf(from)
	if kindTo.Kind() == reflect.Map {
		toMap = to.(map[string]interface{})
	} else {
		toMap = structs.Map(to)
	}
	if kindFrom.Kind() == reflect.Map {
		fromMap = from.(map[string]interface{})
	} else {
		fromMap = structs.Map(from)
	}
	for k, v := range fromMap {
		_, ok := toMap[k]
		if !ok {
			continue
		}
		toMap[k] = v
	}
	if err := mapstructure.Decode(toMap, dst); err != nil {
		return errors.Wrap(err, "failed to decode")
	}
	return nil
}
