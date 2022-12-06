package main

import (
	"fmt"

	"github.com/glennhartmann/aoc22/src/01_1/getcalorieses"
	"github.com/glennhartmann/aoc22/src/common"
)

func main() {
	calorieses := getcalorieses.Get()
	fmt.Printf("%d\n", common.SliceMax(calorieses))
}
