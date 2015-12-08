package store

import (
	"encoding/json"
	"fmt"
	"path"
	"strings"
)

func ensurePrefix(prefix, key string) string {
	if strings.HasPrefix(prefix, key) {
		return key
	}

	return path.Join(prefix, key)
}

func flattenJSONForWriting(body []byte) (map[string][]byte, error) {
	result := map[string][]byte{}

	values := map[string]interface{}{}
	err := json.Unmarshal(body, &values)
	if err != nil {
		return nil, err
	}

	type Item struct {
		path  string
		value interface{}
	}
	queue := []Item{}

	for k, v := range values {
		queue = append(queue, Item{k, v})
	}

	for len(queue) != 0 {
		item := queue[0]
		queue = queue[1:]

		props, ok := item.value.(map[string]interface{})
		if ok {
			for k, v := range props {
				queue = append(queue, Item{item.path + "/" + k, v})
			}
		} else {
			result[item.path] = []byte(fmt.Sprintf("%v", item.value))
		}
	}

	return result, nil
}
