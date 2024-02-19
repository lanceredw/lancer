package main

import (
	"fmt"
	"lancer/cmd"
	"lancer/global"
	"lancer/migrate/method"
	"lancer/migrate/migratefile"
)

func main() {

	cmd.StartMigrate()

	method.MigrateRun(&migratefile.InitTable{}) //init table

	method.Init()

	//method.MigrateRun(&migratefile.User20231221155800{})
	//method.MigrateRun(&migratefile.User20240102171241{})

	fmt.Println(global.DB.Name())

}
