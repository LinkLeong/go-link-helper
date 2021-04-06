package main

import (
	"fmt"
	"go-link-helper/logging"
)

func main()  {
	logging.Setup()
	fmt.Println("11111")
	logging.Error("dddddd")
}
