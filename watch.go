package watch

import (
	"errors"
	"fmt"
	"time"
)

// Watch is the struct of Watch
type Watch struct {
	Sp time.Time
	Ep time.Time
	Lp time.Time
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
	instance = &Watch{
		Sp: now,
		Lp: now,
	}

	return instance
}

// Stop function will stop this watch
func (watch *Watch) Stop() {
	watch.Ep = time.Now()
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
	if watch.Lp.IsZero() {
		return errors.New("Error GetLapTime(): watch isn't started")
	}

	temp := time.Now()
	fmt.Println(temp.Sub(watch.Lp))
	watch.Lp = temp

	return nil
}

// GetDuration is showing duration between start and end
func (watch *Watch) GetDuration() error {
	if watch.Sp.IsZero() {
		return errors.New("Error GetDuration(): watch isn't started")
	}
	if watch.Ep.IsZero() {
		return errors.New("Error GetDuration(): watch isn't stoped")
	}

	fmt.Println(watch.Ep.Sub(watch.Sp))
	return nil
}
