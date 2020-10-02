package httpretryafter

import (
	"errors"
	"strconv"
	"time"
)

var (
	// ErrNegativeSecondsNotAllowed is parsing error that represents seconds value is negative.
	// The seconds value in Retry-After must be positive.
	ErrNegativeSecondsNotAllowed = errors.New("negative seconds not allowed")

	// ErrInvalidFormat is parsing error that represents given Retry-After neither valid seconds nor valid HTTP date.
	ErrInvalidFormat = errors.New("Retry-After value must be seconds integer or HTTP date string")

	nowFunc = func() time.Time { return time.Now() }
)

// Parse tries to parse the value as seconds or HTTP date.
func Parse(retryAfter string) (time.Time, error) {
	if dur, err := ParseSeconds(retryAfter); err == nil {
		now := nowFunc()
		return now.Add(dur), nil
	}
	if dt, err := ParseHTTPDate(retryAfter); err == nil {
		return dt, nil
	}
	return time.Time{}, ErrInvalidFormat
}

// ParseSeconds parses the value as seconds.
func ParseSeconds(retryAfter string) (time.Duration, error) {
	seconds, err := strconv.ParseInt(retryAfter, 10, 64)
	if err != nil {
		return time.Duration(0), err
	}
	if seconds < 0 {
		return time.Duration(0), ErrNegativeSecondsNotAllowed
	}
	return time.Second * time.Duration(seconds), nil
}

// ParseHTTPDate parses the value as HTTP date.
func ParseHTTPDate(retryAfter string) (time.Time, error) {
	parsed, err := time.Parse(time.RFC1123, retryAfter)
	if err != nil {
		return time.Time{}, err
	}
	return parsed, nil
}
