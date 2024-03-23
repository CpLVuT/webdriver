package main

import (
	"fmt"
	"os"
	"webdriver_server/config"
	databases "webdriver_server/db"
)

func main() {
	configPath := "./origin.conf"
	if len(os.Args) == 3 {
		if os.Args[1] == "-c" {
			configPath = os.Args[2]
		}
	}
	config.CONF = config.InitConfig(configPath)

	databases.DB = databases.InitDB(config.CONF["db_user"], config.CONF["db_pass"], config.CONF["db_addr"], config.CONF["db_port"], config.CONF["db_name"])

	f, err := os.Create(config.CONF["log_file"])
	if err != nil {
		fmt.Println("log_file error", err)
		return
	}
	defer f.Close()

	f, err = os.OpenFile(config.CONF["log_file"], os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

}
