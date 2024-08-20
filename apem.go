package apem

import (
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/karincake/apem/appa"
	"github.com/karincake/apem/dba"
	"github.com/karincake/apem/httpa"
	"github.com/karincake/apem/loggera"
	"github.com/karincake/apem/msa"
	"gopkg.in/yaml.v3"

	hs "github.com/karincake/apem/http-std"
)

type extCall func()

var extCalls []extCall

var App *apemCfg
var CfgFile string
var CfgContent []byte

// init
func init() {
	App = &apemCfg{
		AppCfg: &appa.AppCfg{
			CodeName: "apem",
			FullName: "Apem Instance",
			Version:  "0.0.1",
			Env:      "development",
		},
		LoggerCfg: &loggera.LoggerCfg{},
		DbCfg:     &dba.DbCfg{},
		MsCfg:     &msa.MsCfg{},
		HttpCfg:   &httpa.HttpCfg{},
	}
	App.initCfg()
}

// app start the App
func Run(h http.Handler, m ...any) {
	readines := 0
	loggerIdx := -1

	initExtCall()

	for i := range m {
		if myModule, ok := m[i].(loggera.LoggerItf); ok {
			readines++
			loggerIdx = i
			myModule.Init(App.LoggerCfg, App.AppCfg)
		} else if myModule, ok := m[i].(dba.DbItf); ok {
			myModule.Init(App.DbCfg, App.AppCfg)
		} else if myModule, ok := m[i].(msa.MsItf); ok {
			myModule.Init(App.MsCfg, App.AppCfg)
		}
	}

	if readines < 1 {
		log.Fatal("Please make sure App.Run supplied by mandatory adapters implementation. Missing: `loggera`.")
	}

	hs.O.Init(App.HttpCfg, &h, App.AppCfg, m[loggerIdx].(loggera.LoggerItf))
}

func ParseCfg(cfg any) {
	var err error
	if CfgContent == nil {
		CfgContent, err = os.ReadFile(CfgFile)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}

	oriVal := reflect.ValueOf(cfg).Interface()
	err = yaml.Unmarshal(CfgContent, oriVal)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

func ParseSingleCfg(cfg any) {
	// get the file content
	var baseCfg map[string]any = map[string]any{}
	content, err := os.ReadFile("./config.yml")
	if err != nil {
		log.Fatalf("%v", err)
	}

	// parse file content
	err = yaml.Unmarshal(content, baseCfg)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// get cfg name
	cv := reflect.ValueOf(cfg)
	for cv.Kind() == reflect.Pointer || cv.Kind() == reflect.Interface || cv.Kind() == reflect.Bool {
		cv = cv.Elem()
	}
	ctn := firstLetterToLower(cv.Type().Name())

	// non struct cant be filled
	if cv.Kind() != reflect.Struct {
		panic("input requires struct type")
	}

	ct := cv.Type()
	values := baseCfg[ctn].(map[string]any)
	for i := 0; i < cv.NumField(); i++ {
		// identify field type and value of the field
		ft := ct.Field(i)
		fv := cv.Field(i)

		key := keyOrYamlTag(ft.Name, ft.Tag.Get("yaml"))
		value, ok := values[key]
		if !ok {
			continue
		}

		ftName := ft.Name
		fvKind := fv.Kind()
		var err error
		if fvKind != reflect.Pointer {
			if fvKind == reflect.Bool {
				boolValue := "false"
				if value.(bool) {
					boolValue = "true"
				}
				reflectValueFiller(fv, fvKind, ftName, boolValue)
			} else if fvKind == reflect.Slice || fvKind == reflect.Array {
				sliceValue := joinInterfaceSlice(value.([]interface{}))
				reflectValueFiller(fv, fvKind, ftName, sliceValue)
			} else {
				reflectValueFiller(fv, fvKind, ftName, value.(string))
			}

		} else {
			reflectPointerValueFiller(fv, fv.Type().Elem().Kind(), ftName, value.(string))
		}
		if err != nil {
			panic("input can not be parsed at field: " + ftName)
		}
	}
}

func RegisterExtCall(funcToCall extCall) {
	extCalls = append(extCalls, funcToCall)
}

func initExtCall() {
	log.Print("Executing extra calls")
	for _, init := range extCalls {
		init()
	}
}
