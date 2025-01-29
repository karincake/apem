package apem

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/karincake/apem/appa"
	"github.com/karincake/apem/dba"
	"github.com/karincake/apem/httpa"
	"github.com/karincake/apem/loggera"
	lo "github.com/karincake/apem/loggero"
	"github.com/karincake/apem/msa"
)

type apemCfg struct {
	AppCfg         *appa.AppCfg          `yaml:"appCfg"`
	LoggerCfg      *loggera.LoggerCfg    `yaml:"loggerCfg"`
	DbCfg          *dba.DbCfg            `yaml:"dbCfg"`
	MsCfg          *msa.MsCfg            `yaml:"msCfg"`
	HttpCfg        *httpa.HttpCfg        `yaml:"httpCfg"`
	RateLimiterCfg *httpa.RateLimiterCfg `yaml:"rateLimiterCfg"`
}

func (obj *apemCfg) initCfg() {
	CfgFile = "./config.yml"
	flag.StringVar(&CfgFile, "config-file", "./config.yml", "Cfgiguration path (default=./config.yaml)")
	yamlFile, err := os.ReadFile(CfgFile)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = yaml.Unmarshal(yamlFile, obj)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	lo.I.Print("Loaded config successfully")
}
