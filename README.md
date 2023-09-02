# Apem : API Bundle
A simple bundle that helps you manage the configuration to build http ReST-API server using Go.

## Goal
To provide instance of each feature so it will be ready to use by adjusting the configuration according to your environment. 

## The Features
Basic features, which is automatically apply the configuration on the start, including:
- Language, built in
- Logger, utilizing zap (https://github.com/uber-go/zap)
- HTTP Server, utilizing standard version, net/http (https://pkg.go.dev/net/http)

Additional features, which can be loaded when needed simply by importing and mark as _ :
- Database, for now we have gorm (https://github.com/go-gorm/gorm)
- Memory Storage, for now we have redis (github.com/go-redis/redis)

All of the features are designed as a separate package except the HTTP Sever, since it is the core of the package

## Usage
Adjust the configuration of each feature in the config.yml according to your need. 

Due toe Http server's job that handle request and response, which is related to the business logic, you need to provide the handler first. The package uses standard Http/Net package from golang therefore you can utilize any framework that compatible with it or just create from scratch.

As for the Lang, it provides the very basic data that is used by several responses that are need by the bundle. Most of the time you are going to need to provide the data yourself, and place it inside a file according to the configuration on the config.yml file. For example with the following configuration:
```
langConf
  active: en
  path: ./lang
  fileName: data.json
```
you will need an accessible file in `./lang/en/data.json` relative to the main package

## Default Values
There are some configurations with default value as listed below:
|Key|Value|
|---|---|
|FullName|Apem Instance|
|Version|1.0.0|
|Env|development|
|LangConf.Active|en|
|LangConf.Path|./lang|
|LangConf.FileName|data.json|
|HttpConf.Host|127.0.0.1|
|HttpConf.Port|8000|
