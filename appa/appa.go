package appa

type AppCfg struct {
	CodeName string `yaml:"codeName"`
	FullName string `yaml:"fullName"`
	Env      string
	Version  string
}
