package formdatahelper

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type keyVal struct {
	Key string
	Val string
}

// Uses standard net/http
// Any non primitive value of struct will be ignored due to its nature (key-value pairs of string)
func CopyToStruct(input any, r *http.Request) error {
	// identiy value and loop if its pointer until reaches non pointer
	inputV := reflect.ValueOf(input)

	// loop until we get what kind lays behind the input
	for inputV.Kind() == reflect.Pointer || inputV.Kind() == reflect.Interface {
		inputV = inputV.Elem()
	}

	// non struct cant be validated
	if inputV.Kind() != reflect.Struct {
		panic("input requires struct type")
	}

	// check each field
	// inputT := reflect.TypeOf(inputV.Interface()) // keep this for now
	inputT := inputV.Type()
	for i := 0; i < inputV.NumField(); i++ {
		// identify field type and value of the field
		ft := inputT.Field(i)
		fv := inputV.Field(i)
		if !fv.CanSet() {
			continue
		}

		key := ft.Tag.Get("json")
		if key != "" {
			keys := strings.Split(key, ",")
			if keys[0] != "" {
				key = keys[0]
			} else {
				key = ft.Name
			}
		} else {
			key = ft.Name
		}

		fName := ft.Name
		rv := r.FormValue(key)
		ftName := ft.Type.String()
		ftNameClean := strings.Trim(ftName, "*")
		switch {
		case ftName == "string":
			fv.SetString(rv)
		case ftName == "*string" && !fv.IsNil():
			reflect.Indirect(fv).SetString(rv)
		case ftName == "bool":
			if rv == "true" {
				fv.SetBool(true)
			} else {
				fv.SetBool(false)
			}
		case ftName == "*bool" && !fv.IsNil():
			if rv == "true" {
				reflect.Indirect(fv).SetBool(true)
			} else if rv == "false" {
				reflect.Indirect(fv).SetBool(false)
			}
		case len(ftNameClean) >= 3 && ftNameClean[:3] == "int": // bundle in one
			if rv != "" {
				rvVal, err := strconv.Atoi(rv)
				if err != nil {
					return fmt.Errorf("can not convert %s into number", fName)
				} else {
					if ftName[:1] != "*" {
						if fv.OverflowInt(int64(rvVal)) {
							return fmt.Errorf("value overflow for %s", fName)
						} else {
							fv.SetInt(int64(rvVal))
						}
					} else if !fv.IsNil() {
						if reflect.Indirect(fv).OverflowInt(int64(rvVal)) {
							return fmt.Errorf("value overflow for %s", fName)
						} else {
							reflect.Indirect(fv).SetInt(int64(rvVal))
						}
					}
				}
			} else if ftName[:1] != "*" {
				fv.SetInt(0)
			}
		case len(ftNameClean) >= 5 && ftNameClean[:5] == "float": // bundle in one
			if rv != "" {
				floatType := 32
				if ftName == "float64" {
					floatType = 64
				}
				rvVal, err := strconv.ParseFloat(rv, floatType)
				if err != nil {
					return fmt.Errorf("can not convert %s into number", fName)
				} else {
					if ftName[:1] != "*" {
						if fv.OverflowFloat(rvVal) {
							return fmt.Errorf("value overflow for %s", fName)
						} else {
							fv.SetFloat(rvVal)
						}
					} else if !fv.IsNil() {
						if reflect.Indirect(fv).OverflowFloat(rvVal) {
							return fmt.Errorf("value overflow for %s", fName)
						} else {
							reflect.Indirect(fv).SetFloat(rvVal)
						}
					}
				}
			} else if ftName[:1] != "*" {
				fv.SetFloat(0)
			}
		}
	}

	return nil
}
