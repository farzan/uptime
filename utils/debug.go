package debug

import (
	"fmt"
	"github.com/astaxie/beego"
)

const debugMode = "debugmonitoring"

var debug bool

func init() {
	debug = beego.AppConfig.DefaultBool(debugMode, false)
}

func Printf(format string, a ...interface{}) (n int, err error) {
	if debug {
		return fmt.Printf(format, a...)
	} else {
		return
	}
}

func Println(a ...interface{}) (n int, err error) {
	if debug {
		return fmt.Println(a...);
	} else {
		return
	}
}

func Print(a ...interface{}) (n int, err error) {
	if debug {
		return fmt.Print(a...);
	} else {
		return
	}
}