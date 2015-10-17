package ccserialize

import (
	"reflect"
	"strconv"
	"strings"
)

func Serialize(input interface{}) string {
	switch reflect.ValueOf(input).Kind() {
	case reflect.Int:
		return strconv.Itoa(input.(int))
	case reflect.Int64:
		return strconv.FormatInt(input.(int64), 10)
	case reflect.Float64:
		return strconv.FormatFloat(input.(float64), 'f', 6, 64)
	case reflect.String:
		return "\"" + strings.Replace(input.(string), "\"", "\\\"", -1) + "\""
	case reflect.Map:
		t := reflect.TypeOf(input)
		v := reflect.ValueOf(input)
		it := reflect.TypeOf((*interface{})(nil)).Elem()
		m := reflect.MakeMap(reflect.MapOf(t.Key(), it))

		for _, mk := range v.MapKeys() {
			m.SetMapIndex(mk, v.MapIndex(mk))
		}

		return mapSerialize(m.Interface().(map[string]interface{}))
	case reflect.Slice:
		s := reflect.ValueOf(input)
		ret := make([]interface{}, s.Len())

		for i := 0; i < s.Len(); i++ {
			ret[i] = s.Index(i).Interface()
		}

		return arraySerialize(ret)
	case reflect.Struct:
		return structSerialize(input)
	default:
		panic("Attempt to serialize unsupported type.")
	}
}

func structSerialize(input interface{}) string {
	serializedString := "{"

	intStr := reflect.ValueOf(input)
	str := reflect.TypeOf(input)
	for i := 0; i < str.NumField(); i++ {
		f := intStr.Field(i)
		keyName := str.Field(i).Name
		if str.Field(i).Tag.Get("ccserialize") != "" {
			keyName = str.Field(i).Tag.Get("ccserialize")
		}

		serializedString = serializedString + "[\"" + strings.ToLower(keyName) + "\"]="
		serializedString = serializedString + Serialize(f.Interface()) + ","
	}

	return serializedString + "}"
}

func arraySerialize(input []interface{}) string {
	serializedString := "{"
	for k, v := range input {
		serializedString = serializedString + "[" + strconv.Itoa(k+1) + "]=" + Serialize(v) + ","
	}
	return serializedString + "}"
}

func mapSerialize(input map[string]interface{}) string {
	serializedString := "{"
	for k, v := range input {
		serializedString = serializedString + "[\"" + k + "\"]=" + Serialize(v) + ","
	}
	return serializedString + "}"
}
