package main

import (
	"./entity"
	"flag"
	"fmt"
)

var (
	stopper		= make(chan bool)
)

func main() {
	flag.Parse()
	token := flag.Arg(0)
	_,err := entity.New(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening...")
	<-stopper
	return
}

