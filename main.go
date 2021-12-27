package main

import (
	"github.com/edwardsuwirya/wmbTableMgmt/config"
)

func main() {
	appConfig := config.NewConfig()
	appConfig.RunMigration()
	appConfig.StartEngine()
}
