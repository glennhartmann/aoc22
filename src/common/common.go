// Package common contains a set of miscellaneous helper functions (generic,
// when applicable) that should maybe be part of Go's standard library.
package common

import (
	"strings"

	"golang.org/x/exp/constraints"
)

// Real is all non-complex number types.
type Real interface {
	constraints.Integer | constraints.Float
}

// ConvenientLenable is the set of non-map types than len() makes sense on.
type ConvenientLenable[T any] interface {
	~string | ~[]T
}

// Lenable is the set of types than len() makes sense on.
// maps make this really annoying because they require
// a second type constraint which can't be automatically
// inferred if you want to use one of the non-map
// lenable types.
type Lenable[T any, C comparable] interface {
	ConvenientLenable[T] | ~map[C]T
}

// SliceSum returns the total sum of a slice of Real numbers.
func SliceSum[T Real](slice []T) T {
	var sum T
	for _, n := range slice {
		sum += n
	}
	return sum
}

// SliceMax returns the maximum element of a slice of Ordered values.
func SliceMax[T constraints.Ordered](slice []T) T {
	return FsliceMax(slice, func(e T) T { return e })
}

// SliceMin returns the minimum element of a slice of Ordered values.
func SliceMin[T constraints.Ordered](slice []T) T {
	return FsliceMin(slice, func(e T) T { return e })
}

// Max returns the maximum element of the given Ordered values.
func Max[T constraints.Ordered](a, b T, rest ...T) T {
	return SliceMax(append(rest, a, b))
}

// Min returns the minimum element of the given Ordered values.
func Min[T constraints.Ordered](a, b T, rest ...T) T {
	return SliceMin(append(rest, a, b))
}

// Fjoin combines a slice of values into a single string separated by |sep|.
// This is the same as the built-in strings.Join() except that the caller needs
// to pass in a function to convert the input values into strings.
func Fjoin[T any](elems []T, sep string, str func(e T) string) string {
	var sb strings.Builder

	first := true
	for _, e := range elems {
		if !first {
			sb.WriteString(sep)
		}
		first = false
		sb.WriteString(str(e))
	}

	return sb.String()
}

// Abs returns the absolute value of the given Real number.
func Abs[T Real](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

// FsliceMax returns the maximum Ordered value that results from applying the
// given function to each individual element of a slice.
func FsliceMax[T any, R constraints.Ordered](slice []T, f func(e T) R) R {
	var max R
	for i, e := range slice {
		n := f(e)
		if i == 0 || n > max {
			max = n
		}
	}
	return max
}

// FsliceMin returns the minimum Ordered value that results from applying the
// given function to each individual element of a slice.
func FsliceMin[T any, R constraints.Ordered](slice []T, f func(e T) R) R {
	var min R
	for i, e := range slice {
		n := f(e)
		if i == 0 || n < min {
			min = n
		}
	}
	return min
}

// Fmax returns the maximum Ordered value that results from applying the given
// function to each of the given values.
func Fmax[T any, R constraints.Ordered](f func(e T) R, a, b T, rest ...T) R {
	return FsliceMax(append(rest, a, b), f)
}

// Fmin returns the minimum Ordered value that results from applying the given
// function to each of the given values.
func Fmin[T any, R constraints.Ordered](f func(e T) R, a, b T, rest ...T) R {
	return FsliceMin(append(rest, a, b), f)
}

// Longest returns the longest ConvenientLenable out of a slice.
func Longest[T any, L ConvenientLenable[T]](s []L) int {
	return FsliceMax(s, func(e L) int { return len(e) })
}

// Padding generates a padding string made up of |r| repeated copies of a
// string |p|.
func Padding(p string, r int /* repititions */) string {
	var sb strings.Builder
	for i := 0; i < r; i++ {
		sb.WriteString(p)
	}
	return sb.String()
}

// PadToLeft takes in a string, |s|, a padding string |p|, and a target number
// of characters, |c|. If len(s) >= c, it returns s. Otherwise, it returns a
// string consisting of some amount of repititions of |p|, followed by |s|,
// such that the total length of the return value is as close as possible to
// |c|.
//
// if (len(s) - r) % len(p) != 0, this won't be aligned. Usually best to stick
// with len(p) = 1
//
// Also, if len(p) == 0, this will crash.
func PadToLeft(s, p string, c int /* characters, not repititions */) string {
	return padToPadding(s, p, c) + s
}

// PadToRight takes in a string, |s|, a padding string |p|, and a target number
// of characters, |c|. If len(s) >= c, it returns s. Otherwise, it returns a
// string consisting of |s| followed by some amount of repititions of |p|, such
// that the total length of the return value is as close as possible to |c|.
//
// if (len(s) - r) % len(p) != 0, this won't be aligned. Usually best to stick
// with len(p) = 1
//
// Also, if len(p) == 0, this will crash.
func PadToRight(s, p string, c int /* characters, not repititions */) string {
	return s + padToPadding(s, p, c)
}

func padToPadding(s, p string, c int /* characters, not repititions */) string {
	e := c - len(s)
	if e <= 0 {
		return ""
	}

	return Padding(p, e/len(p))
}
