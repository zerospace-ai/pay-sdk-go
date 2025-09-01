package rsa_utils

import (
	"bytes"
	"encoding/json"
	"sort"
	"strconv"
	"strings"
)

func ToStringMap(jStr []byte) map[string]string {
	resMap := make(map[string]string)
	var jsonObject map[string]json.RawMessage
	if err := json.Unmarshal(jStr, &jsonObject); err == nil {
		for key, value := range jsonObject {
			resMap[key] = valueAsString(value)
		}
	}
	return resMap
}

func valueAsString(element json.RawMessage) string {
	var asString string

	if err := json.Unmarshal(element, &asString); err == nil {
		return asString
	}

	var asObject map[string]json.RawMessage
	if err := json.Unmarshal(element, &asObject); err == nil {
		sortedKeys := make([]string, 0, len(asObject))
		for key := range asObject {
			sortedKeys = append(sortedKeys, key)
		}
		sort.Slice(sortedKeys, func(i, j int) bool {
			return strings.ToLower(sortedKeys[i]) < strings.ToLower(sortedKeys[j])
		})

		var buffer bytes.Buffer
		for _, key := range sortedKeys {
			buffer.WriteString(valueAsString(asObject[key]))
		}
		return buffer.String()
	}

	var asArray []json.RawMessage
	if err := json.Unmarshal(element, &asArray); err == nil {
		sortedKeys := make(map[string]string)
		for _, item := range asArray {
			var arrayObject map[string]json.RawMessage
			if err := json.Unmarshal(item, &arrayObject); err == nil {
				for key, value := range arrayObject {
					sortedKeys[key] = valueAsString(value)
				}
			}
		}

		var buffer bytes.Buffer
		keys := make([]string, 0, len(sortedKeys))
		for key := range sortedKeys {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, key := range keys {
			buffer.WriteString(sortedKeys[key])
		}
		return buffer.String()
	}

	var asInt int64
	if err := json.Unmarshal(element, &asInt); err == nil {
		return strconv.FormatInt(asInt, 10)
	}
	return ""
}
