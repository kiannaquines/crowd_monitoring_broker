package main

import (
	_ "github.com/go-sql-driver/mysql"
	"parser/utils"
)

func main() {
	utils.InitializeDatabase()
	utils.MqttClientInit()
}
