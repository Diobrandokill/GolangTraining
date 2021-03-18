package main

import (
	winniepooh "GolangTraining/02_package/icomefromalaska"
	"fmt"
	"GolangTraining/02_package/stringutil"
	// someAlias "github.com/GoesToEleven/GolangTraining/02_package/icomefromalaska"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println(stringutil.MyName)
	fmt.Println(winniepooh.BearName)
	fmt.Println(stringutil.Reverse(winniepooh.BearName))
	fmt.Println(stringutil.Reverse(stringutil.MyName))
}
