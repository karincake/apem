package apem

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

func keyOrYamlTag(key, yamlTag string) string {
	if yamlTag == "" {
		return strings.ToLower(key)
	}
	tagByte := []byte(yamlTag)
	pos := len(tagByte)
	for i, v := range tagByte {
		if v == 44 {
			pos = i
		}
	}
	return string(tagByte[:pos])
}

func joinInterfaceSlice(slice []interface{}) string {
	// Create a slice to hold the string representations
	strSlice := make([]string, len(slice))

	// Convert each element to a string
	for i, v := range slice {
		strSlice[i] = fmt.Sprintf("%v", v)
	}

	// Join the slice of strings with commas
	return strings.Join(strSlice, ",")
}

func reflectValueFiller(fv reflect.Value, vk reflect.Kind, ftName, rvs string) {
	switch {
	case vk == reflect.String:
		fv.SetString(rvs)
	case vk == reflect.Bool:
		if rvs == "true" || rvs == "yes" || rvs == "1" {
			fv.SetBool(true)
		} else if rvs == "false" || rvs == "no" || rvs == "0" {
			fv.SetBool(false)
		}
	case vk >= reflect.Uint && vk <= reflect.Uint64:
		if rvs != "" {
			rvsVal, err := strconv.ParseUint(rvs, 10, 64)
			if err != nil {
				panic("can not convert \"" + ftName + "\" (value: " + rvs + ") into number")
			}
			if fv.OverflowUint(uint64(rvsVal)) {
				panic("value overflow for \"" + ftName + "\" (value: " + rvs + ")")
			} else {
				fv.SetUint(uint64(rvsVal))
			}
		}
	case vk >= reflect.Int && vk <= reflect.Int64:
		if rvs != "" {
			rvsVal, err := strconv.Atoi(rvs)
			if err != nil {
				panic("can not convert \"" + ftName + "\" (value: " + rvs + ") into number")
			}
			if fv.OverflowInt(int64(rvsVal)) {
				panic("value overflow for \"" + ftName + "\" (value: " + rvs + ")")
			} else {
				fv.SetInt(int64(rvsVal))
			}
		}
	case vk >= reflect.Float32 && vk <= reflect.Float64:
		if rvs != "" {
			floatType := 32
			if ftName == "float64" {
				floatType = 64
			}
			rvsVal, err := strconv.ParseFloat(rvs, floatType)
			if err != nil {
				panic("can not convert \"" + ftName + "\" (value: " + rvs + ") into number")
			}
			if fv.OverflowFloat(rvsVal) {
				panic("value overflow for \"" + ftName + "\" (value: " + rvs + ")")
			} else {
				fv.SetFloat(rvsVal)
			}
		}
	case vk == reflect.Slice || vk == reflect.Array:
		if rvs != "" {
			// Split the input string into an array of strings
			rvsArr := strings.Split(rvs, ",")

			// Create a new slice with the appropriate length and capacity
			slice := reflect.MakeSlice(fv.Type(), len(rvsArr), len(rvsArr))

			// Set the slice elements
			for i := 0; i < len(rvsArr); i++ {
				elem := slice.Index(i)
				reflectValueFiller(elem, elem.Kind(), ftName, rvsArr[i])
			}

			// Set the field with the created slice
			fv.Set(slice)
		}
	}
}

func reflectPointerValueFiller(fv reflect.Value, vk reflect.Kind, ftName, rvs string) {
	switch {
	case vk == reflect.String:
		fv.Set(reflect.ValueOf(&rvs))
	case vk == reflect.Bool:
		if rvs == "true" || rvs == "yes" || rvs == "1" {
			rvsx := true
			fv.Set(reflect.ValueOf(&rvsx))
		} else if rvs == "false" || rvs == "no" || rvs == "0" {
			rvsx := false
			fv.Set(reflect.ValueOf(&rvsx))
		}
	case vk >= reflect.Uint && vk <= reflect.Uint64:
		if rvs != "" {
			rvsVal, err := strconv.ParseUint(rvs, 10, 64)
			if err != nil {
				panic("can not convert \"" + ftName + "\" (value: " + rvs + ") into number")
			}
			if fv.OverflowUint(uint64(rvsVal)) {
				panic("value overflow for \"" + ftName + "\" (value: " + rvs + ")")
			} else {
				rvsx := uint64(rvsVal)
				fv.Set(reflect.ValueOf(&rvsx))
			}
		}
	case vk >= reflect.Int && vk <= reflect.Int64:
		if rvs != "" {
			rvsVal, err := strconv.Atoi(rvs)
			if err != nil {
				panic("can not convert \"" + ftName + "\" (value: " + rvs + ") into number")
			}
			if fv.OverflowInt(int64(rvsVal)) {
				panic("value overflow for \"" + ftName + "\" (value: " + rvs + ")")
			} else {
				rvsx := int64(rvsVal)
				fv.Set(reflect.ValueOf(&rvsx))
			}
		}
	case vk >= reflect.Float32 && vk <= reflect.Float64:
		if rvs != "" {
			floatType := 32
			if ftName == "float64" {
				floatType = 64
			}
			rvsVal, err := strconv.ParseFloat(rvs, floatType)
			if err != nil {
				panic("can not convert \"" + ftName + "\" (value: " + rvs + ") into number")
			}
			if fv.OverflowFloat(rvsVal) {
				panic("value overflow for \"" + ftName + "\" (value: " + rvs + ")")
			} else {
				fv.Set(reflect.ValueOf(&rvsVal))
			}
		}
	}
}

func firstLetterToLower(s string) string {

	if len(s) == 0 {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToLower(r[0])

	return string(r)
}
