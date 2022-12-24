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
	var max T
	for i, n := range slice {
		if i == 0 || n > max {
			max = n
		}
	}
	return max
}

func SliceMin[T Real](slice []T) T {
	var min T
	for i, n := range slice {
		if i == 0 || n < min {
			min = n
		}
	}
	return min
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
