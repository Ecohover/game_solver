package main

import (
	"fmt"
	"game_solver/internal/helper/filereader"
)

var config = ""

func main() {

	fmt.Println(filereader.ReadFile("game_solver/configs/config.json"))
}
