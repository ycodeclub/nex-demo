package main

import (
	"log"

	"github.com/ycodeclub/nex-demo/ioc"
)

func main() {
	if err := ioc.Run(); err != nil {
		log.Fatal(err)
	}
}
