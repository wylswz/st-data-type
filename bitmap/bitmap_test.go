package bitmap

import (
	"st/profile"
	"testing"
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

func TestMem(t *testing.T) {

	SIZE := 300000000
	a := make([]bool, SIZE)
	print("=====Using boolean array: =====\n")
	profile.TimeSpaced(func() {
		for i := 0; i < SIZE; i++ {
			a[i] = true
		}
	})()

	print("=====Using BitMap=====\n")
	profile.TimeSpaced(func() {
		a2 := NewBitMap(SIZE)
		for i := 0; i < SIZE; i++ {
			a2.UnsafeSet(i)
		}
	})()

}
