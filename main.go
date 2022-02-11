package main

import (
	"go-trading/app/controllers"
	"go-trading/config"
	"go-trading/utils"
	"log"
)

// func main() {
// 	utils.LoggingSettings(config.Config.LogFile)
// 	fmt.Println(models.DbConnection)
// 	controllers.StreamIngestionData()
// 	controllers.StartWebServer()
// }

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())
}
