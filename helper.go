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

func reflectValueFiller(fv reflect.Value, vk reflect.Kind, ftName, rvs any) {
	switch vk {
	case reflect.String:
		fv.SetString(fmt.Sprint(rvs))

	case reflect.Bool:
		switch val := rvs.(type) {
		case bool:
			fv.SetBool(val)
		case string:
			lower := strings.ToLower(val)
			if lower == "true" || lower == "yes" || lower == "1" {
				fv.SetBool(true)
			} else if lower == "false" || lower == "no" || lower == "0" {
				fv.SetBool(false)
			}
		case int, int64, uint, uint64:
			fv.SetBool(fmt.Sprint(val) != "0")
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var num uint64
		switch val := rvs.(type) {
		case uint64:
			num = val
		case uint, uint32, uint16, uint8:
			num = reflect.ValueOf(val).Convert(reflect.TypeOf(uint64(0))).Uint()
		case int, int64, int32, int16, int8:
			num = uint64(reflect.ValueOf(val).Int())
		case string:
			if val != "" {
				parsed, err := strconv.ParseUint(val, 10, 64)
				if err != nil {
					panic(fmt.Sprintf("cannot convert %q (value: %v) into uint", ftName, rvs))
				}
				num = parsed
			}
		}
		if fv.OverflowUint(num) {
			panic(fmt.Sprintf("value overflow for %q (value: %v)", ftName, rvs))
		}
		fv.SetUint(num)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var num int64
		switch val := rvs.(type) {
		case int64:
			num = val
		case int, int32, int16, int8:
			num = reflect.ValueOf(val).Int()
		case uint, uint64, uint32, uint16, uint8:
			num = int64(reflect.ValueOf(val).Uint())
		case string:
			if val != "" {
				parsed, err := strconv.ParseInt(val, 10, 64)
				if err != nil {
					panic(fmt.Sprintf("cannot convert %q (value: %v) into int", ftName, rvs))
				}
				num = parsed
			}
		}
		if fv.OverflowInt(num) {
			panic(fmt.Sprintf("value overflow for %q (value: %v)", ftName, rvs))
		}
		fv.SetInt(num)

	case reflect.Float32, reflect.Float64:
		var num float64
		switch val := rvs.(type) {
		case float64:
			num = val
		case float32:
			num = float64(val)
		case int, int64:
			num = float64(reflect.ValueOf(val).Int())
		case uint, uint64:
			num = float64(reflect.ValueOf(val).Uint())
		case string:
			if val != "" {
				bitSize := 32
				if vk == reflect.Float64 {
					bitSize = 64
				}
				parsed, err := strconv.ParseFloat(val, bitSize)
				if err != nil {
					panic(fmt.Sprintf("cannot convert %q (value: %v) into float", ftName, rvs))
				}
				num = parsed
			}
		}
		if fv.OverflowFloat(num) {
			panic(fmt.Sprintf("value overflow for %q (value: %v)", ftName, rvs))
		}
		fv.SetFloat(num)

	case reflect.Slice, reflect.Array:
		if str, ok := rvs.(string); ok && str != "" {
			parts := strings.Split(str, ",")
			slice := reflect.MakeSlice(fv.Type(), len(parts), len(parts))
			for i := 0; i < len(parts); i++ {
				elem := slice.Index(i)
				reflectValueFiller(elem, elem.Kind(), ftName, strings.TrimSpace(parts[i]))
			}
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
