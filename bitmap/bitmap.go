package bitmap

import (
	"errors"
)

const ONE = uint64(1)

type BitMap struct {
	data []uint64
	size int
}

func (bm *BitMap) Size(s int) int {
	return bm.size
}

func (bm *BitMap) Get(i int) (bool, error) {
	if i >= bm.size {
		return false, errors.New("")
	}
	return bm.UnsafeGet(i), nil
}

func (bm *BitMap) UnsafeGet(i int) bool {
	return (bm.data[i>>6] & (ONE << (i & 63))) > 0
}

func (bm *BitMap) Set(i int) error {
	if i >= bm.size {
		return errors.New("")
	}
	bm.UnsafeSet(i)
	return nil
}

func (bm *BitMap) UnsafeSet(i int) {
	bm.data[i>>6] |= (ONE << (i & 63))
}

func (bm *BitMap) Unset(i int) error {
	if i >= bm.size {
		return errors.New("")
	}
	bm.UnsafeUnset(i)
	return nil
}

func (bm *BitMap) UnsafeUnset(i int) {
	if bm.UnsafeGet(i) {
		bm.data[i>>6] &= (^(ONE << (i & 63)))
	}
}

func NewBitMap(size int) *BitMap {
	return &BitMap{
		data: make([]uint64, size>>6+1),
		size: size,
	}
}
