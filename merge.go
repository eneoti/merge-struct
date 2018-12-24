package merp

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"reflect"
)

// MergeOverwrite overwrite map
func MergeOverwrite(to, from, dst interface{}) error {
	var toMap map[string]interface{}
	var result map[string]interface{}
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
		mapBase(k, v, &fromMap)
		toMap[k] = fromMap[k]
	}
	_, ok := toMap["Base"]
	if ok {
		delete(toMap, "Base")
	}

	conv := ConventionalMarshaller{Value: toMap}
	b, err := conv.MarshalJSON()
	if err != nil {
		return err
	}
	json.Unmarshal(b, &result)
	// fmt.Println(result)
	// fmt.Println(toMap)
	if err := mapstructure.Decode(toMap, dst); err != nil {
		return errors.Wrap(err, "failed to decode")
	}
	return nil
}
func mapBase(key string, value interface{}, org *(map[string]interface{})) {
	result := *org
	if key == "Base" {
		temp := value.(map[string]interface{})
		for k, v := range temp {
			result[k] = v
		}
		delete(result, "Base")
		return
	}
	valueRf := reflect.ValueOf(value)
	kind := valueRf.Kind()
	switch kind {
	case reflect.Map:
		temp := value.(map[string]interface{})
		for k, v := range temp {
			mapBase(k, v, &temp)
		}
	case reflect.Slice:
		temp := value.([]interface{})
		for _, v := range temp {
			temp2 := v.(map[string]interface{})
			for k1, v1 := range temp2 {
				mapBase(k1, v1, &temp2)
			}
		}
	case reflect.Ptr:
		if !valueRf.IsNil() {
			temp := valueRf.Elem()
			// fmt.Println(reflect.TypeOf(value))
			fmt.Println(temp)
			// temp = temp.(map[string]interface{})
			// for k, v := range temp {
			// 	mapBase(k, v, &temp)
			// }
			// value = &temp
		}
	case reflect.String:
		result[key] = value
	}
	org = &result
}
