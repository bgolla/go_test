package main

import (
	"fmt"

	simplepb "github.com/simplesteph/protobuf-example-go/src/simple"
	log "github.com/sirupsen/logrus"
)

func main() {
	doSimple()
}
func doSimple() {
	//  a := example_simple.SimpleMessage{}
	fmt.Printf("this is printf")
	log.Info("abc")
	a := simplepb.SimpleMessage{}
	fmt.Print(a)
}
