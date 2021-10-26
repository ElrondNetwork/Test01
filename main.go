package main

import (
	"fmt"

	"github.com/ElrondNetwork/Test02-renamed/components"
)

func main() {
	fmt.Println("app Test01")

	c := components.NewComponent("test-name", 37)
	c.PrintName()
	c.PrintPoints()
}
