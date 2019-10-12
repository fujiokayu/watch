package watch

import (
	"errors"
	"fmt"
	"time"
)

// Watch is struct of Watch
type Watch struct {
	sp time.Time
	ep time.Time
	lp time.Time
}

// manage instance as a singleton
var instance *Watch

// Now is showing date and time
func Now() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))
}

// Start function initialize watch
func Start() *Watch {
	if instance != nil {
		return instance
	}

	now := time.Now()
	myWatch := &Watch{
		sp: now,
		lp: now,
	}

	return myWatch
}

// Stop function stop this watch
func (watch *Watch) Stop() {
	watch.ep = time.Now()
}

/*
func (watch *Watch) Reset() error {
	if instance == nil {
		return errors.New("Error Reset(): watch isn't initialized")
	}

	now := time.Now()
	watch = &Watch{
		sp: now,
		lp: now,
		ep: time.Time,
	}
}
*/

// GetLapTime is showing the lap time
func (watch *Watch) GetLapTime() error {
	if watch.lp.IsZero() {
		return errors.New("Error GetLapTime(): watch isn't started")
	}

	temp := time.Now()
	fmt.Println(temp.Sub(watch.lp))
	watch.lp = temp

	return nil
}

// GetDuration is showing duration between start and end
func (watch *Watch) GetDuration() error {
	if watch.sp.IsZero() {
		return errors.New("Error GetDuration(): watch isn't started")
	}
	if watch.ep.IsZero() {
		return errors.New("Error GetDuration(): watch isn't stoped")
	}

	fmt.Println(watch.ep.Sub(watch.sp))
	return nil
}
