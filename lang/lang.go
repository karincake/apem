package lang

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	l "github.com/karincake/ambeng/lepet"
	"go.uber.org/zap"
	"golang.org/x/exp/maps"

	lz "github.com/karincake/apem/loggerzap"
)

// Configuration type that is used by the core
type LangConf struct {
	Active   string
	Path     string
	FileName string
}

// Instance of the language
var I *l.LangData = &l.LangData{}

func Init(conf LangConf) {
	I.Active = conf.Active
	I.SetList("en", defaultList)

	jsonFile, err := os.Open(fmt.Sprintf("%v/%v/%v", conf.Path, conf.Active, conf.FileName))
	if err != nil {
		lz.I.Fatal("failed to load source file. " + err.Error())
	}
	defer jsonFile.Close()

	var myLI l.LangItem = l.LangItem{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &myLI)
	maps.Copy(I.List["en"], myLI)
	lz.I.Info("instantiation", zap.String("feature", "lang"), zap.String("source", "built-in"), zap.String("status", "done"))
}
