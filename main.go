package main

import (
	"UptimeMonitor/domain"
	_ "UptimeMonitor/routers"
	"github.com/astaxie/beego"
)

func main() {
	// Running monitoring goroutines
	domain.Run()

	beego.Run()
}
