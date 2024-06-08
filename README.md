# Apem : Config Bundle
A simple Go package that helps you manage your web-api project's configuration in certain areas. The areas that are currently covered are:
- http server (net/http)
- database server (gorm mysql, gorm postgres)
- memory storage server (redis)
- logger (zerolog, zap)
- language (karincake)

## Usage and Explanation
The concept is pretty simple:
- Set the configuration's values through `config.yml` file (can be changed to other name)
- Call apem's main fuction (`apem.Run()`) along with providing the `http.Handler` and adapter's object (we just call it that way for now) you want to use.

For example, to create a very basic http server for your api that prints `Hello World`, you can create a `config.yml` file with the following content:
```
httpConf:
  host: 127.0.0.1
  port: 8100
```
Note: the default configuration file name is `./config.yml`. You can specify the name through flag `-config-file` when calling the application, for example: `go run . -config-file=my-config.yaml`.

Prepare your handler
```
func createHandlers() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	r.HandleFunc("GET /some-path", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Still hello world!!"))
	})

	return r
}
```

Then in your go main package use `apem` as follows
```
package main

import (
	a "github.com/karincake/apem"
	l "github.com/karincake/apem/logger-zerolog"
)

func main() {
    // Your starting point wher you supply the adapters, here in the following
    // example it supplies logger adapter since it is required by apem itself.
	a.Run(createHandlers(), &l.O) // &l.O is the logger's object
}
```

Note that logger will always be required to be supplied since you have to decice which logger you want to use.

As matter of fact, there is another package that is always being used: `net/http`, but since `apem` only uses standard library then you don't have to supply.

Another thing worth to note is that, since it uses standard library `net/http`, you can use router package such as `chi` to manage your route.

For more example please visit: `www.github.com/karincake/apem-usage-samples`

## The Adapters
Adapters are just packages that help you in applying the configuration into other packages, such as `net/http`, `gorm/mysql`, or `redis`. Each adapter implements interface of each area.

It read the configurations and stores them in an exported objects of a struct name `O` of each adapter which you supply, then creates an instance in an exported variable named `I` of ach adapter, which you can use everywhere later on.

In the earlier example, the `Run` function is supplied with `loger-zerolog` object (`l.O`) which will be used to store the configuration.


## Covered Areas
### App
Just your app information

Configuration structure with sample:
```
fullName: Apem
codeName: apem
version: 0.0.1
env: development
```

There are some configurations with default value as listed below:
|Key|Value|
|---|---|
|FullName|Apem Instance|
|Version|1.0.0|
|Env|development|

### Http Server
Covered packages:
- `net/htpp`, standard library

Configuration structure with sample:
```
httpConf:
  host: 127.0.0.1
  port: 8100
```

### Logger
Covered packages:
- `https://github.com/rs/zerolog`
- `https://github.com/uber-go/zap`


Configuration structure with sample (mysql):
```
loggerConf:
  mode:
  level:
  output:
```

Due to the needs of standardization for the logger because of it's being used by the core, the interface also has several methods for logging purpose you can use as well listed as follows:
- `Debug()`
- `Info()`
- `Warning()`
- `Error()`
- `Panic()`
- `Fatal()`
- `Bool(string, bool)`
- `Int(string, int)`
- `Int8(string, int8)`
- `Int16(string, int16)`
- `Int32(string, int32)`
- `Int64(string, int64)`
- `Uint(string, uint)`
- `Uint8(string, uint8)`
- `Uint16(string, uint16)`
- `Uint32(string, uint32)`
- `Uint64(string, uint64)`
- `String(string, string)`
- `Send()`

### Database Server
Covered packages:
- `https://github.com/go-gorm/gorm` using `https://github.com/go-gorm/mysql`
- `https://github.com/go-gorm/gorm` using `https://github.com/go-gorm/postgres`

Configuration structure with sample (with mysql dsn example):
```
dbConf:
  dsn: acoount:password@tcp(127.0.0.1:3306)/my-database?charset=utf8mb4&parseTime=True&loc=Local
  maxOpenConns: 5
  maxIdleConns: 5
  maxIdleTime: 100
```

### Memory Storage
Covered packages:
- `https://github.com/redis/go-redis`

Configuration structure with sample:
```
msConf:
  dsn: 127.0.0.1:6379
```

### Language
- `https://github.com/karincake/lang`

Configuration structure with sample (mysql):
```
langConf
  active: en
  path: ./lang
  fileName: data.json
```

## Always Used Areas
- Http Server, due to its main purpose: web-api.
- Logger, as it is used by the package itself, required in the main function call.
