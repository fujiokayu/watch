package watch_test

import (
	"fmt"
	"testing"

	"github.com/fujiokayu/watch"
)

var bar = "----------------------------"

func ExampleWatch() {
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

func TestSuccess_1(t *testing.T) {
	fmt.Println("TestSuccess_1: " + bar)

	myWatch := watch.Start()
	watch.Now()

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
	if err != nil {
		// this function should be fail
		fmt.Println(err)
	}
}
