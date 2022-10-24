package main

import route "fgd-golang/task/satu/routes"

func main() {
	var PORT = ":8080"

	route.StartServer().Run(PORT)
}
