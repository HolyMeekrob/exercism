// Package gigasecond has tools for working with time based on the gigasecond standard
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import (
	"fmt"
	"time"
)

// AddGigasecond takes a timestamp, adds one billion seconds to it, and returns the resulting timestamp
func AddGigasecond(t time.Time) time.Time {
	const oneBillion = 1_000_000_000
	// Drop the error because we know it works
	duration, _ := time.ParseDuration(fmt.Sprintf("%ds", oneBillion))
	return t.Add(duration)
}
