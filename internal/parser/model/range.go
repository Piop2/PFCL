package model

import (
	"errors"
	"math"
)

type RangeValue interface {
	~int64 | ~float64
}

type Range[T RangeValue] struct {
	Min T
	Max T
}

func NewRange[T RangeValue](min T, max T) (*Range[T], error) {
	if min > max {
		return nil, errors.New("invalid range: min > max")
	}
	return &Range[T]{Min: min, Max: max}, nil
}

// Helper constructors for infinite ranges: FullRangeInt64, FllRangeFloat64

func FullRangeInt64() *Range[int64] {
	return &Range[int64]{Min: math.MinInt64, Max: math.MaxInt64}
}

func FullRangeFloat64() *Range[float64] {
	return &Range[float64]{Min: math.Inf(-1), Max: math.Inf(1)}
}

// Contains checks if value is inside the range.
func (r Range[T]) Contains(v T) bool {
	return v >= r.Min && v <= r.Max
}
