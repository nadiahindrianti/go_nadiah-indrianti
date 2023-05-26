package main

import (
	"Task_1/configs"
	"Task_1/routes"
)

func main() {
	// create a new echo instance
	configs.InitDB()
	configs.InitialMigration()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8888"))
}
