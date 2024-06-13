# Apem : Config Bundle
A simple Go package that helps you start up your web-api application by reading the configurations and apply the configuration to the `adadpters` (we just call it that way for now) to some external libraries.

Currently the configurations covers several area as follows:
- http server (net/http)
- database server (gorm mysql, gorm postgres)
- memory storage server (redis)
- logger (zerolog, zap)

All the read configuration is stored in an exported instance of `apemConf` named `App`.

## Usage and More Explanation
The concept is pretty simple:
- Set the configuration's values through `config.yml` file (can be changed to other name)
- Call apem's main fuction (`apem.Run()`) along with providing the `http.Handler` and adapter's object you want to use.

For example, to create a very basic http server for your api that write `Hello World`, you can create a `config.yml` file with the following content:
```yml
httpConf:
  host: 127.0.0.1
  port: 8100
```
Prepare your handler
```go
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
```go
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

Note:
- Adapter for `httpa` doesn't have to be supplied since it only has one adapter (`net/http`) and will always be used.
- Adapter for `loggera` have to be supplied since it is needed and you have to decice which logger you want to use.

## Extra Call
You can run extra logic by utilizing the `apem.RegisterExtCall` that will register functions that will be executed before the http server runs. You can register function as many as you want.

Example:
```go
import (
	a "github.com/karincake/apem"
)

func Create() http.Handler {
	a.App.RegisterExtrCall(myTings) // Register a function to be called before the http server start

	r.HandleFunc("/", (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello world!!"))
  })

	r := http.NewServeMux()
	return r
}

function myThings() {
  // do anything here
}

```
You can utilize the function anywhere as long as it is called before the `apem.Run`, for example: inside the `init()` function of a pakcage.

## The Adapters
Adapters are just packages that help you in applying the configuration into other packages, such as `net/http`, `gorm/mysql`, or `redis`. Each adapter implements interface of each area.

Apem reads the configurations and stores them in an exported objects of a struct name `O` of each adapter which you supply, then creates an instance in an exported variable named `I` of ach adapter, which you can use everywhere later on.

In the earlier example, the `Run` function is supplied with `loger-zerolog` object (`l.O`) which will be used to store the configuration.

## The Area
### App
Just your app information

Configuration structure with sample:
```yml
fullName: Apem
codeName: apem
version: 0.0.1
env: development
```

Some default values:
|Key|Value|
|---|---|
|CodeName|apem|
|FullName|Apem Instance|
|Version|0.0.1|
|Env|development|

### Http Server
Covered packages:
- `net/htpp`, standard library

Configuration structure with sample:
```yml
httpConf:
  host: 127.0.0.1
  port: 8100
```

### Logger
Covered packages:
- `https://github.com/rs/zerolog`
- `https://github.com/uber-go/zap`


Configuration structure with sample (mysql):
```yml
loggerConf:
  mode:
  level:
  output:
```

Due to the needs of standardization for the logger because of it's being used by the core, the interface also has several methods for logging purpose you can use as well listed as follows:
- `Debug()` - log level
- `Info()` - log level
- `Warning()` - log level
- `Error()` - log level
- `Panic()` - log level
- `Fatal()` - log level
- `Bool(string, bool)` - set key-val with type accordingly
- `Int(string, int)` - set key-val with type accordingly
- `Int8(string, int8)` - set key-val with type accordingly
- `Int16(string, int16)` - set key-val with type accordingly
- `Int32(string, int32)` - set key-val with type accordingly
- `Int64(string, int64)` - set key-val with type accordingly
- `Uint(string, uint)` - set key-val with type accordingly
- `Uint8(string, uint8)` - set key-val with type accordingly
- `Uint16(string, uint16)` - set key-val with type accordingly
- `Uint32(string, uint32)` - set key-val with type accordingly
- `Uint64(string, uint64)` - set key-val with type accordingly
- `String(string, string)` - set key-val with type accordingly
- `Send()` - send the log

### Database Server
Covered packages:
- `https://github.com/go-gorm/gorm` using `https://github.com/go-gorm/mysql`
- `https://github.com/go-gorm/gorm` using `https://github.com/go-gorm/postgres`

Configuration structure with sample (with mysql dsn example):
```yml
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
```yml
msConf:
  dsn: 127.0.0.1:6379
```

### Always Used Area
- Http Server, due to its main purpose: web-api.
- Logger, as it is used by the package itself, required in the main function call.

## Samples
- `https://github.com/karincake/apem-sample01-basic`

More sample will be coming