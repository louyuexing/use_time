package use_time

import (
	"time"
)

// UseTime is a struct for tracking and managing time-related information.
// It contains the origin time, a function to get the current time, and the beginning time for calculating time differences.
type UseTime struct {
	origin time.Time        // The origin time, usually set when the instance is created.
	g      func() time.Time // A function to get the current time, for obtaining the current time in a customizable way.
	b      time.Time        // The beginning time, used as a reference for calculating time differences.
}

// New creates and returns a pointer to a UseTime instance.
// Parameter origin: The origin time, marking the starting point for time tracking.
func New(origin time.Time) *UseTime {
	// Initialize a UseTime instance, setting the t function to return the current time.
	u := UseTime{
		g: func() time.Time {
			return time.Now()
		},
		origin: origin,
	}
	// Set the beginning time to the origin time.
	u.b = u.origin
	// Return the pointer to the UseTime instance.
	return &u
}

// Step returns the time duration since the last call to Step or Cross.
// It updates the internal beginning time to the current time.
func (u *UseTime) Step() time.Duration {
	// Calculate the time difference since the last beginning time.
	d := time.Since(u.b)
	// Update the beginning time to the current time.
	u.b = u.g()
	// Return the calculated time difference.
	return d
}

// Cross returns the time duration since the origin time.
func (u *UseTime) Cross() time.Duration {
	// Calculate the time difference since the origin time.
	return time.Since(u.origin)
}
