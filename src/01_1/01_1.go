package main

import (
	"fmt"
	"io"

	"github.com/glennhartmann/aoc22/src/common"
)

func main() {
	calorieses := make([]int64, 0, 500)
	var curr, sum int64
	for {
		n, err := fmt.Scanln(&curr)
		if err == io.EOF {
			break
		}
		if n == 0 {
			calorieses = append(calorieses, sum)
			sum = 0
			continue
		} else if err != nil {
			panic(fmt.Sprintf("got %d arguments, expected 0 or 1", n))
		}
		sum += curr
	}
	fmt.Printf("%d\n", common.SliceMax(calorieses))
}
