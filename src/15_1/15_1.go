package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/glennhartmann/aoclib/common"
)

type point struct {
	x, y int
}

type rect struct {
	min, max point
}

func (r rect) expand(min, max point) rect {
	return rect{
		min: point{common.Min(r.min.x, min.x), common.Min(r.min.y, min.y)},
		max: point{common.Max(r.max.x, max.x), common.Max(r.max.y, max.y)},
	}
}

func (r rect) String() string {
	indent := " "
	topLeft := fmt.Sprintf("(%d, %d)", r.min.x, r.min.y)
	topRight := fmt.Sprintf("(%d, %d)", r.max.x, r.min.y)
	bottomLeft := fmt.Sprintf("(%d, %d)", r.min.x, r.max.y)
	bottomRight := fmt.Sprintf("(%d, %d)", r.max.x, r.max.y)

	if len(topLeft) > len(bottomLeft) {
		bottomLeft = spaces(len(topLeft)-len(bottomLeft)) + bottomLeft
	} else {
		topLeft = spaces(len(bottomLeft)-len(topLeft)) + topLeft
	}
	indent += spaces(len(topLeft))

	return fmt.Sprintf(`%s       %s
%s+---+
%s|   |
%s|   |
%s+---+
%s       %s`, topLeft, topRight, indent, indent, indent, indent, bottomLeft, bottomRight)
}

func spaces(n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString(" ")
	}
	return sb.String()
}

type sensor struct {
	loc, beacon point
	rng         int
}

func main() {
	sensors := make([]sensor, 0, 50)
	r := rect{
		min: point{math.MaxInt64, math.MaxInt64},
		max: point{math.MinInt64, math.MinInt64},
	}
	for {
		var sx, sy, bx, by int
		n, err := fmt.Scanf("Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d\n", &sx, &sy, &bx, &by)
		if err == io.EOF || err == io.ErrUnexpectedEOF { // why unexpected?
			break
		}
		if n != 4 || err != nil {
			panic(err)
		}

		s := sensor{loc: point{sx, sy}, beacon: point{bx, by}}
		s.rng = dist(s.loc, s.beacon)
		sensors = append(sensors, s)

		r = r.expand(point{s.loc.x - s.rng, s.loc.y - s.rng}, point{s.loc.x + s.rng, s.loc.y + s.rng})
	}

	log.Printf("window of interest:")
	log.Printf("\n%s\n", r.String())

	//logTestGrid(sensors)

	const row = 2000000
	total := 0
	for i := r.min.x; i <= r.max.x; i++ {
		if cannotExist(sensors, point{i, row}) {
			total++
		}
	}
	log.Printf("total: %d", total)
}

// manhattan distance
func dist(src, dst point) int {
	return common.Abs(dst.x-src.x) + common.Abs(dst.y-src.y)
}

func isSensor(sensors []sensor, loc point) bool {
	return slices.ContainsFunc(sensors, func(s sensor) bool {
		return s.loc == loc
	})
}

func isBeacon(sensors []sensor, loc point) bool {
	return slices.ContainsFunc(sensors, func(s sensor) bool {
		return s.beacon == loc
	})
}

func cannotExist(sensors []sensor, loc point) bool {
	for _, s := range sensors {
		if loc != s.beacon && isWithinRange(s, loc) {
			return true
		}
	}
	return false
}

func isWithinRange(s sensor, loc point) bool {
	return dist(s.loc, loc) <= s.rng
}

func logTestGrid(sensors []sensor) {
	log.Print("test grid:")
	tg := make([][]byte, 61)
	for i := range tg {
		tg[i] = make([]byte, 61)
		for j := range tg[i] {
			r := i - 30
			c := j - 30

			if isSensor(sensors, point{c, r}) {
				tg[i][j] = 'S'
			} else if isBeacon(sensors, point{c, r}) {
				tg[i][j] = 'B'
			} else if cannotExist(sensors, point{c, r}) {
				tg[i][j] = '#'
			} else {
				tg[i][j] = '.'
			}
		}
		log.Printf("%s", string(tg[i]))
	}
	log.Print("")
}
