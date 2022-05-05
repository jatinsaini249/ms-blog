package main

import "github.com/jatinsaini249/ms-blog/routes"

func main() {
	router := routes.Initialize()
	router.Run(":4000")
}
