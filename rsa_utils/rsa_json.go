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
	// Trying to parse as nil (JSON null)
	if string(element) == "null" {
		return ""
	}

	// Trying to parse as bool
	var asBool bool
	if err := json.Unmarshal(element, &asBool); err == nil {
		return strconv.FormatBool(asBool)
	}

	// Trying to parse as Number
	var asNumber json.Number
	if err := json.Unmarshal(element, &asNumber); err == nil {
		return asNumber.String()
	}

	// Trying to parse as string
	var asString string
	if err := json.Unmarshal(element, &asString); err == nil {
		return asString
	}

	// Trying to parse as an object
	var asObject map[string]json.RawMessage
	if err := json.Unmarshal(element, &asObject); err == nil {
		keys := make([]string, 0, len(asObject))
		for k := range asObject {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool {
			return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
		})

		var buffer bytes.Buffer
		for _, k := range keys {
			buffer.WriteString(valueAsString(asObject[k]))
		}
		return buffer.String()
	}

	// Trying to parse as an array
	var asArray []json.RawMessage
	if err := json.Unmarshal(element, &asArray); err == nil {
		var buffer bytes.Buffer
		for _, item := range asArray {
			buffer.WriteString(valueAsString(item))
		}
		return buffer.String()
	}

	return ""
}
