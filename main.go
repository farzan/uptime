package main

import (
	"UptimeMonitor/domain/monitoring"
	_ "UptimeMonitor/routers"
	"github.com/astaxie/beego"
)

func main() {
	// Running monitoring goroutines
	monitoring.Run()

	beego.Run()
}
