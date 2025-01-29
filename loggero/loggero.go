package loggero

import (
	"log"
	"os"
)

var I *log.Logger = log.New(os.Stdout, "", log.LstdFlags)
