package appa

type AppConf struct {
	CodeName string
	FullName string `yaml:"fullName"`
	Env      string
	Version  string
}
