package apem

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"go.uber.org/zap"

	lz "github.com/karincake/apem/loggerzap"
)

type LangConf struct {
	Active  string
	SrcPath string
}

type langItem map[string]string
type langData struct {
	Active string
	list   map[string]langItem
}

var I *langData // instance

func New() langData {
	return langData{}
}

func Init(conf LangConf) {
	I = &langData{}
	I.Active = conf.Active
	I.list = map[string]langItem{}

	jsonFile, err := os.Open(conf.SrcPath)
	if err != nil {
		lz.I.Fatal("failed to load source file. " + err.Error())
	}
	defer jsonFile.Close()

	var myLI langItem = langItem{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &myLI)
	I.list["en"] = myLI
	lz.I.Info("instantiation", zap.String("feature", "lang"), zap.String("source", "built-in"), zap.String("status", "done"))
}

func (a *langData) Add(name string) {
	_, ok := a.list[name]
	if !ok {
		a.list[name] = langItem{}
	}
}

func (a *langData) AddMsg(code string, message string, opt ...string) {
	lang := a.Active
	if len(opt) > 0 {
		a.Add(opt[1])
		lang = opt[1]
	}
	a.list[lang][code] = message
}

func (a *langData) AddMsgList(list langItem, opt ...string) {
	lang := a.Active
	if len(opt) > 0 {
		a.Add(opt[1])
		lang = opt[1]
	}

	for k, v := range list {
		a.list[lang][k] = v
	}
}

func (a *langData) Msg(k string, opt ...string) string {
	lang := a.Active
	if len(opt) > 0 {
		a.Add(opt[1])
		lang = opt[1]
	}

	if msg, ok := a.list[lang][k]; !ok {
		return "** warning: usage of unlisted code **"
	} else {
		return msg
	}
}
