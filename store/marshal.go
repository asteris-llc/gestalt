package store

import (
	"fmt"
	"strconv"
)

func marshal(value interface{}) []byte {
	return []byte(fmt.Sprintf("%v", value))
}

func unmarshal(value []byte, typ string) (interface{}, error) {
	stringed := string(value)

	switch typ {
	case "string":
		return stringed, nil

	case "integer":
		return strconv.Atoi(stringed)

	case "float":
		return strconv.ParseFloat(stringed, 64)

	case "boolean":
		return strconv.ParseBool(stringed)

	default:
		return nil, fmt.Errorf(`don't know how to unmarshal "%v"`, typ)
	}
}
