package watch_test

import (
	"fmt"
	"testing"

	"github.com/fujiokayu/watch"
)

var bar = "----------------------------"

func ExampleWatch() {
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

func TestSuccess_1(t *testing.T) {
	fmt.Println("TestSuccess_1: " + bar)

	myWatch := watch.Start()

	err := myWatch.GetLapTime()
	if err != nil {
		t.Fatal("failed test")
	}

	myWatch.Stop()

	err = myWatch.GetDuration()
	if err != nil {
		t.Fatal("failed test")
	}

	err = myWatch.GetLapTime()
	if err != nil {
		t.Fatal("failed test")
	}
}

func TestFailed_1(t *testing.T) {

	fmt.Println("TestFailed_1 : " + bar)
	myWatch := watch.Start()

	err := myWatch.GetDuration()
	if err == nil {
		// this function should be fail
		t.Fatal("failed test")
	}
}
