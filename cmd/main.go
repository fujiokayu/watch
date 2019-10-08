package main

import (
	"github.com/fujiokayu/watch"
)

func main() {
	watch.Start()
	watch.Watch()
	watch.GetLapTime()
	watch.Stop()
	watch.GetDuration()
	watch.GetLapTime()
}
