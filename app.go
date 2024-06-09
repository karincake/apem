package apem

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/karincake/apem/appa"
	"github.com/karincake/apem/dba"
	"github.com/karincake/apem/httpa"
	"github.com/karincake/apem/langa"
	"github.com/karincake/apem/loggera"
	"github.com/karincake/apem/msa"
)

type extCall func()

type apemConf struct {
	Conf            *appa.AppConf
	LoggerConf      *loggera.LoggerConf    `yaml:"loggerConf"`
	LangConf        *langa.LangConf        `yaml:"langConf"`
	DbConf          *dba.DbConf            `yaml:"dbConf"`
	MsConf          *msa.MsConf            `yaml:"msConf"`
	HttpConf        *httpa.HttpConf        `yaml:"httpConf"`
	RateLimiterConf *httpa.RateLimiterConf `yaml:"rateLimiterConf"`
	extCalls        []extCall
}

func (a *apemConf) initConfig() {
	cfgFile := "./config.yml"
	flag.StringVar(&cfgFile, "config-file", "./config.yml", "Configuration path (default=./config.yaml)")
	yamlFile, err := os.ReadFile(cfgFile)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = yaml.Unmarshal(yamlFile, a)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Print("Loaded config successfully")
}

func (a *apemConf) RegisterExtCall(e extCall) {
	a.extCalls = append(a.extCalls, e)
}

func (a *apemConf) initExtCall() {
	for _, init := range a.extCalls {
		init()
	}
}
