package bitmap

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestBitmap(t *testing.T) {
	a := NewBitMap(3000)
	a.Set(30)
	val, err := a.Get(30)
	if err != nil {
		t.Fail()
	}
	if !val {
		t.Fail()
	}

	err = a.Unset(30)
	if err != nil {
		t.Fail()
	}
	if val, _ = a.Get(30); val {
		t.Fail()

	}
	err = a.Unset(30)
	if err != nil {
		t.Fail()
	}
	if val, _ = a.Get(30); val {
		t.Fail()

	}

	a.Set(2999)
	val, err = a.Get(2999)
	if err != nil {
		t.Fail()
	}
	if !val {
		t.Fail()
	}

	err = a.Set(3000)
	if err == nil {
		t.Fail()
	}
}

func TestMisc(t *testing.T) {
	a := 0b00011101
	a &= ^(1 << 4)
	print(a)
}

func timed(f func()) {
	s1 := time.Now().UnixNano()
	f()
	fmt.Printf("Time used: %v ms\n", (time.Now().UnixNano()-s1)/1000)
}

func TestMem(t *testing.T) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	m1 := m.Alloc

	SIZE := 300000000
	a := make([]bool, SIZE)
	print("=====Using boolean array: =====\n")
	timed(func() {
		for i := 0; i < SIZE; i++ {
			a[i] = true
		}
	})
	runtime.ReadMemStats(&m)
	m2 := m.Alloc - m1
	fmt.Printf("Memory Used: %v\n", m2)

	print("=====Using BitMap=====\n")
	timed(func() {
		a2 := NewBitMap(SIZE)
		for i := 0; i < SIZE; i++ {
			a2.UnsafeSet(i)
		}
	})
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory Used: %v\n", m.Alloc-m2)
}
