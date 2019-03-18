package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"ocr_test/manager/configManager"
	"ocr_test/manager/logManager"
	"ocr_test/manager/serverManager"
)

/*
	Seetacloud API server
*/
func main() {

	//get configs
	conf := configManager.GetConf()

	//set release mode
	gin.SetMode(conf.ServerMode)

	if conf.ServerMode == gin.ReleaseMode {
		//redirect the ouput to access log file
		gin.DefaultWriter = logManager.GetLogFileWriter()
	} else {
		res, _ := json.Marshal(conf)
		fmt.Println("Configuration set: ", string(res))
	}

	//Initial the server
	serverManager.Initial()

	//Load the router
	serverManager.LoadRouter()

	//create server and start to listen the port
	serverManager.CreateServer(conf.Port)

	err := serverManager.StartToListen()
	if err != nil {
		logManager.LogError(err.Error(), false)
	}
}
