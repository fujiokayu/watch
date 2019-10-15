package main

import (
	"fmt"

	"github.com/fujiokayu/watch"
)

func main() {

	// show current date time
	watch.Now()

	//start to measure the time
	myWatch := watch.Start()

	// first calling of GetLapTime(), will be shown the time since Start()
	err := myWatch.GetLapTime()
	if err != nil {
		fmt.Println(err)
	}

	// stop to measure the time
	myWatch.Stop()

	// show the time since Start() untill Stop()
	err = myWatch.GetDuration()
	if err != nil {
		fmt.Println(err)
	}

	// show the time since previous GetLapTime()
	err = myWatch.GetLapTime()
	if err != nil {
		fmt.Println(err)
	}
}
