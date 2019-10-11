package watch

import (
	"fmt"
	"time"
)

// Duration is exported gloval value
type Duration int64

var sp time.Time
var ep time.Time

// Watch : mock function
func Watch() {
	fmt.Println(time.Now())
}

// Start : mock function
func Start() {
	sp = time.Now()
}

// Stop : mock function
func Stop() {
	ep = time.Now()
}

// GetLapTime : show time since start
func GetLapTime() {
	fmt.Println(time.Since(sp))
}

// GetDuration : show duration between start and end
func GetDuration() {
	fmt.Println(ep.Sub(sp))
}
