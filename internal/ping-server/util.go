package ping_server

import (
	_ "embed"
	"time"
)

// StatusOptions is the options provided as query parameters to the status route.
type StatusOptions struct {
	Query   bool
	Timeout time.Duration
}

// PointerOf returns a pointer of the argument passed.
func PointerOf[T any](v T) *T {
	return &v
}

// Contains returns true if the array contains the value.
func Contains[T comparable](arr []T, v T) bool {
	for _, value := range arr {
		if value == v {
			return true
		}
	}

	return false
}

// Map applies the provided map function to all of the values in the array and returns the result.
func Map[I, O any](arr []I, f func(I) O) []O {
	result := make([]O, len(arr))

	for i, v := range arr {
		result[i] = f(v)
	}

	return result
}
