package main

import (
	"api/utils"
	"fmt"
)

func main() {
	var a int = 2
	fmt.Println(a)
	var str = "Wasan Yang Age 27"
	fmt.Println(str)
	fmt.Println(utils.ReplaceSpaceWithDash(str))
}
