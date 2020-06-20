package unixtime

// (unix)time.Time Borrowed largely from https://github.com/pieterclaerhout/example-json-unixtimestamp
// Differences include:
// - Additional forwarded methods

// Note:
// Duration is not included in/does not come from the above mentioned package

import (
	"math"
	"strconv"
	"time"
)

// Time defines a timestamp encoded as epoch seconds in JSON
type Time time.Time

func Now() Time {
	return Time(time.Now())
}

// MarshalJSON is used to convert the timestamp to JSON
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

// UnmarshalJSON is used to convert the timestamp from JSON
func (t *Time) UnmarshalJSON(s []byte) (err error) {
	r := string(s)
	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(q, 0)
	return nil
}

// Add returns the time t+d.
func (t Time) Add(d Duration) Time {
	return Time(time.Time(t).Add(time.Duration(d)))
}

// Sub returns the duration t-u. If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration, the maximum (or minimum) duration
// will be returned.
// To compute t-d for a duration d, use t.Add(-d).
func (t Time) Sub(u Time) Duration {
	return Duration(time.Time(t).Sub(time.Time(u)))
}

// Unix returns t as a Unix time, the number of seconds elapsed
// since January 1, 1970 UTC. The result does not depend on the
// location associated with t.
func (t Time) Unix() int64 {
	return time.Time(t).Unix()
}

// Time returns the JSON time as a time.Time instance in UTC
func (t Time) Time() time.Time {
	return time.Time(t).UTC()
}

// String returns t as a formatted string
func (t Time) String() string {
	return t.Time().String()
}

type Duration time.Duration

// MarshalJSON is used to convert the timestamp to JSON
func (d Duration) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(math.Ceil(time.Duration(d).Seconds())), 10)), nil
}

// UnmarshalJSON is used to convert the timestamp from JSON
func (d *Duration) UnmarshalJSON(s []byte) (err error) {
	r := string(s)
	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Duration)(d) = time.Second * time.Duration(q)
	return nil
}
