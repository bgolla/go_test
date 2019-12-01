package main

import (
	"fmt"
	"github.com/bgolla/go_sample/abc/example_simple"
	log "github.com/sirupsen/logrus"
)

func main() {
	doSimple()
}
func doSimple() {
	//  a := example_simple.SimpleMessage{}
	fmt.Printf("this is printf")
	log.Info("abc")
	a := simple.SimpleMessage{}
	fmt.Print(a)
}
