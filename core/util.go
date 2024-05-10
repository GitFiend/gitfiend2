package core

import (
	"crypto/rand"
	"fmt"
	"time"
)

type Number interface {
	int | uint | float64 | float32
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

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
