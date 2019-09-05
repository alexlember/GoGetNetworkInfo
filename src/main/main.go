package main

import (
	"../httpserver"
	"fmt"
)

func main() {
	fmt.Println("main")

	go httpserver.Start()

	fmt.Println("after go start")

	var input string
	fmt.Scanln(&input)
}
