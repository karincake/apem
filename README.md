# Apem : API Bundle
A simple bundle that helps you manage the configuration to build http ReST-API server using Go.

## Goal
To provide instance of each feature so it will be ready to use by adjusting the configuration according to your environment. 

## The Features
Basic features, which is automatically apply the configuration on the start, including:
- Language, built in
- Logger, utilizing zap
- HTTP Server, utilizing standard version, net/http

Additional features, which can be loaded when needed simply by importing and mark as _ :
- Database, for now we have gorm
- Memory Storage, for now we have redis
- Mailer, for now we have go mail

All of the features are designed as a separate package except the HTTP Server, since it is the very core of the package

## Usage
Each feature only needs configuration in the config.yml, except for the Http Server and Lang. Http server is related to the business logic therefore you need to provide the handler first. Since it uses standard Http/Net package from golang, you can utilize any framework that compatible with it.

As for the Lang, it only provides the very basic data that is used by several responses that are need by the bundle. So you have to provide the handler before running it.