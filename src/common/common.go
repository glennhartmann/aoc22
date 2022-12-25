package common

import (
	"constraints"
	"strings"
)

type Real interface {
	constraints.Integer | constraints.Float
}

func SliceSum[T Real](slice []T) T {
	var sum T
	for _, n := range slice {
		sum += n
	}
	return sum
}

func SliceMax[T Real](slice []T) T {
	return FsliceMax(slice, func(e T) T { return e })
}

func SliceMin[T Real](slice []T) T {
	return FsliceMin(slice, func(e T) T { return e })
}

func Max[T Real](a, b T, rest ...T) T {
	return SliceMax(append(rest, a, b))
}

func Min[T Real](a, b T, rest ...T) T {
	return SliceMin(append(rest, a, b))
}

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

func Abs[T Real](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func FsliceMax[T any, R Real](slice []T, f func(e T) R) R {
	var max R
	for i, e := range slice {
		n := f(e)
		if i == 0 || n > max {
			max = n
		}
	}
	return max
}

func FsliceMin[T any, R Real](slice []T, f func(e T) R) R {
	var min R
	for i, e := range slice {
		n := f(e)
		if i == 0 || n < min {
			min = n
		}
	}
	return min
}

func Fmax[T any, R Real](f func(e T) R, a, b T, rest ...T) R {
	return FsliceMax(append(rest, a, b), f)
}

func Fmin[T any, R Real](f func(e T) R, a, b T, rest ...T) R {
	return FsliceMin(append(rest, a, b), f)
}

func Longest(s []string) int {
	return FsliceMax(s, func(e string) int { return len(e) })
}

func Padding(p string, r int /* repititions */) string {
	var sb strings.Builder
	for i := 0; i < r; i++ {
		sb.WriteString(p)
	}
	return sb.String()
}

// if (len(s) - r) % len(p) != 0, this won't be aligned. Usually best to stick with len(p) = 1
func PadToLeft(s, p string, c int /* characters, not repititions */) string {
	return padToPadding(s, p, c) + s
}

// if (len(s) - r) % len(p) != 0, this won't be aligned. Usually best to stick with len(p) = 1
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
