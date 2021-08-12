package profile

import (
	"fmt"
	"runtime"
	"time"
)

func Timed(f func()) func() {
	return func() {
		s1 := time.Now().UnixNano()
		f()
		fmt.Printf("Time used: %v ms\n", (time.Now().UnixNano()-s1)/1000)
	}
}

func Spaced(f func()) func() {
	return func() {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		start := m.Alloc
		f()
		runtime.ReadMemStats(&m)
		fmt.Printf("Memory used: %v bytes\n", m.Alloc-start)
	}
}

func TimeSpaced(f func()) func() {
	return Timed(Spaced(f))
}
