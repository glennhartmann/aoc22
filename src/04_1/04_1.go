package main

import (
	"fmt"
	"io"
)

func main() {
	var total int64
	for {
		var fstStart, fstEnd, lastStart, lastEnd int64
		n, err := fmt.Scanf("%d-%d,%d-%d\n", &fstStart, &fstEnd, &lastStart, &lastEnd)
		if err == io.EOF {
			break
		}
		if n != 4 || err != nil {
			panic("invalid input")
		}

		if isFullyContained(fstStart, fstEnd, lastStart, lastEnd) || isFullyContained(lastStart, lastEnd, fstStart, fstEnd) {
			total++
		}
	}
	fmt.Printf("%d\n", total)
}

func isFullyContained(xStart, xEnd, yStart, yEnd int64) bool {
	return xStart <= yStart && xEnd >= yEnd
}
