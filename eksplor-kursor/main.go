package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {

	for i := 0; i < 5; i++ {

		robotgo.Move(300, 300)

		time.Sleep(1 * time.Second)

		robotgo.Move(700, 300)

		time.Sleep(1 * time.Second)
	}
}
