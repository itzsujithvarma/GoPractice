package main

import "fmt"

func main() {

	name := "sujith"
	nameaddr := &name
	fmt.Println(&nameaddr)
	myprint(nameaddr)
}

func myprint(addr *string) {
	fmt.Println(&addr)
}
