package main

import (
	"fmt"

	"github.com/fujiokayu/watch"
)

func main() {

	myWatch := watch.Start()
	watch.Now()

	err := myWatch.GetLapTime()
	if err != nil {
		fmt.Println(err)
	}

	myWatch.Stop()

	err = myWatch.GetDuration()
	if err != nil {
		fmt.Println(err)
	}

	err = myWatch.GetLapTime()
	if err != nil {
		fmt.Println(err)
	}
}
