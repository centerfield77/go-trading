package main

import (
	"fmt"
	"go-trading/bitflyer"
	"go-trading/config"
	"go-trading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)

	order := &bitflyer.Order{
		ProductCode:     config.Config.ProductCode,
		ChildOrderType:  "MARKET",
		Side:            "BUY",
		Size:            0.01,
		MinuteToExpires: 1,
		TimeInForce:     "GTC",
	}
	res, _ := apiClient.SendOrder(order)
	fmt.Println(res.ChildOrderAcceptanceID)

	//i := "JRF20181012-144016-140584"
	//params := map[string]string{
	//	"product_code": config.Config.ProductCode,
	//	"child_order_acceptance_id": i,
	//}
	//r, _ := apiClient.ListOrder(params)
	//fmt.Println(r)

}
