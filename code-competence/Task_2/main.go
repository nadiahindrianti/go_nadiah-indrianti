package main

import (
	"Task_2/configs"
	"Task_2/routes"
)

func main() {
	// create a new echo instance
	configs.InitDB()
	configs.InitialMigration()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8888"))
}
