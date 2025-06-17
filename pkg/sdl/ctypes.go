package sdl

import (
	"math"
	"unsafe"
)

type CSlice[T any] struct {
	start *T
}

func newCSlice[T any](ptr *T) CSlice[T] {
	return CSlice[T]{
		start: ptr,
	}
}

func newCSliceRaw[T any](ptr uintptr) CSlice[T] {
	return CSlice[T]{
		start: (*T)(unsafe.Pointer(ptr)),
	}
}

func (s CSlice[T]) Iterate(yield func(it *T) bool) {
	it := s.start

	for {
		if it == nil {
			break
		}

		if !yield(it) {
			break
		}

		it = (*T)(unsafe.Pointer(uintptr(unsafe.Pointer(it)) + unsafe.Sizeof(it)))
	}
}

func (s CSlice[T]) GetLength() int {
	return s.GetLengthMax(math.MaxInt)
}

func (s CSlice[T]) GetLengthMax(maxLen int) (len int) {
	it := s.start

	for len < maxLen {
		if it == nil {
			break
		}

		len += 1
		it = (*T)(unsafe.Pointer(uintptr(unsafe.Pointer(it)) + unsafe.Sizeof(it)))
	}

	return len
}

func (s CSlice[T]) Collect() []T {
	return s.CollectMax(math.MaxInt)
}

func (s CSlice[T]) CollectMax(maxLen int) (res []T) {
	sz := s.GetLengthMax(maxLen)
	res = make([]T, 0, sz)

	for it := range s.Iterate {
		if len(res) == maxLen {
			break
		}

		res = append(res, *it)
	}

	return res
}
