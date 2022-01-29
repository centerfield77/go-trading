package main

import (
	"fmt"
	"go-trading/app/controllers"
	"go-trading/app/models"
	"go-trading/config"
	"go-trading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(models.DbConnection)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}
