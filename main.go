package main

import (
	"fmt"
)

func main() {
	fmt.Println("main of chatbot")
	cb := &ChatBot{}
	cb.Start()
}
