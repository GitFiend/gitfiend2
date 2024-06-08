package shared

import (
	"crypto/rand"
	"fmt"
	"time"
)

// PseudoUuid
// Note - NOT RFC4122 compliant
func PseudoUuid() (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)

	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	}

	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return uuid
}

/*
Elapsed

Call this like:

defer Elapsed("timerName")()
*/
func Elapsed(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("\n%s took %v\n", name, time.Since(start))
	}
}

func Find[T any](slice []T, f func(value T) bool) (T, bool) {
	for _, val := range slice {
		if f(val) {
			return val, true
		}
	}
	return *new(T), false
}

func Map[T any, V any](slice []T, f func(value T) V) []V {
	res := make([]V, len(slice))
	for i, v := range slice {
		res[i] = f(v)
	}
	return res
}

//type Vec[T any] []T
//
//func (v Vec[T]) Map() {}
