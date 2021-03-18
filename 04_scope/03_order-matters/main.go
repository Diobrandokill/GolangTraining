package main

import "fmt"

func main() {
	// fmt.Println(x)
	fmt.Println(y)
	//x := 42
}

var y = 42

// in closure: if var defined after call then it won't work
// however, if var defined at package scope and defined after the main func, it still can be called correctly