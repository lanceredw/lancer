package main

import "lancer/cmd"

func main() {

	//action when closing
	defer cmd.Clean()
	//start the application
	cmd.Start()

}
