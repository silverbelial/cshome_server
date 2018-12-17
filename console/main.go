package main

import (
	"fmt"

	"github.com/silverbelial/cshome_server/conf"
	"github.com/silverbelial/cshome_server/console/controllers"
)

func main() {
	conf.InitialConf(conf.ConfTypeConsole)

	r := controllers.InitialConsoleRouters()

	r.Run(fmt.Sprintf(":%d", conf.GetListeningPort()))
}
