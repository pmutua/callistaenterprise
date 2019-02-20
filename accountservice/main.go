package main

import (
	"fmt"

	"github.com/pmutua/callistaenterprise/goblog/accountservice/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Start %v\n", appName)
	service.StartWebServer("6767") //Hardcode port number
}
