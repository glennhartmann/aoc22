package common

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSliceSum(t *testing.T) {
	tests := []struct {
		name      string
		intIn     []int
		floatIn   []float32
		wantInt   int
		wantFloat float32
	}{
		{
			name:      "empty -> 0",
			intIn:     nil,
			floatIn:   nil,
			wantInt:   0,
			wantFloat: 0.0,
		},
		{
			name:      "single element -> single element's value",
			intIn:     []int{5},
			floatIn:   []float32{3.14159},
			wantInt:   5,
			wantFloat: 3.14159,
		},
		{
			name:      "actual summing",
			intIn:     []int{5, 6, 7, 8, 23534745},
			floatIn:   []float32{3.14159, 4, 8, 2.46357, 4.347234, 462345.3},
			wantInt:   23534771,
			wantFloat: 462367.252,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotInt := SliceSum[int](test.intIn)
			gotFloat := SliceSum[float32](test.floatIn)
			if gotInt != test.wantInt || gotFloat != test.wantFloat {
				t.Errorf("SliceSum() = {int: %d, float: %f}, want {int: %d, float: %f}", gotInt, gotFloat, test.wantInt, test.wantFloat)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		{
			name: "empty -> zero value for slice, not applicable for non-slice",
			vals: nil,
			want: 0,
		},
		{
			name: "single val -> the single val for slice, not applicable for non-slice",
			vals: []int{5},
			want: 5,
		},
		{
			name: "2 vals -> the greater val",
			vals: []int{5, 6},
			want: 6,
		},
		{
			name: "2 vals but reversed -> still the greater val",
			vals: []int{6, 5},
			want: 6,
		},
		{
			name: "3 vals -> the greatest val",
			vals: []int{5, 6, 7},
			want: 7,
		},
		{
			name: "3 vals, different order -> still the greatest val",
			vals: []int{5, 7, 6},
			want: 7,
		},
		{
			name: "3 vals, another different order -> still the greatest val",
			vals: []int{7, 5, 6},
			want: 7,
		},
		{
			name: "N vals",
			vals: []int{3, 2, 34573, 2, 5, 4574, 2345234, 5},
			want: 2345234,
		},
		{
			name: "N vals with negatives",
			vals: []int{3, 2, 34573, -2, 5, -4574, -2345234, 5},
			want: 34573,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotSlice := SliceMax[int](test.vals)
			if gotSlice != test.want {
				t.Errorf("SliceMax() = %d, want %d", gotSlice, test.want)
			}

			if len(test.vals) > 1 {
				got := Max[int](test.vals[0], test.vals[1], test.vals[2:]...)
				if got != test.want {
					t.Errorf("Max() = %d, want %d", got, test.want)
				}
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		{
			name: "empty -> zero value for slice, not applicable for non-slice",
			vals: nil,
			want: 0,
		},
		{
			name: "single val -> the single val for slice, not applicable for non-slice",
			vals: []int{5},
			want: 5,
		},
		{
			name: "2 vals -> the lesser val",
			vals: []int{5, 6},
			want: 5,
		},
		{
			name: "2 vals but reversed -> still the lesser val",
			vals: []int{6, 5},
			want: 5,
		},
		{
			name: "3 vals -> the least val",
			vals: []int{5, 6, 7},
			want: 5,
		},
		{
			name: "3 vals, different order -> still the least val",
			vals: []int{5, 7, 6},
			want: 5,
		},
		{
			name: "3 vals, another different order -> still the least val",
			vals: []int{7, 5, 6},
			want: 5,
		},
		{
			name: "N vals",
			vals: []int{3, 2, 34573, 2, 5, 4574, 2345234, 5},
			want: 2,
		},
		{
			name: "N vals with negatives",
			vals: []int{3, 2, 34573, -2, 5, -4574, -2345234, 5},
			want: -2345234,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotSlice := SliceMin[int](test.vals)
			if gotSlice != test.want {
				t.Errorf("SliceMin() = %d, want %d", gotSlice, test.want)
			}

			if len(test.vals) > 1 {
				got := Min[int](test.vals[0], test.vals[1], test.vals[2:]...)
				if got != test.want {
					t.Errorf("Min() = %d, want %d", got, test.want)
				}
			}
		})
	}
}

func TestFjoin(t *testing.T) {
	tests := []struct {
		name      string
		vals      []float32
		decPlaces int
		sep       string
		want      string
	}{
		{
			name:      "empty separator",
			vals:      []float32{3, 4.2, 86.34, 4325.24363452},
			decPlaces: 3,
			sep:       "",
			want:      "3.0004.20086.3404325.244",
		},
		{
			name:      "empty main string",
			vals:      []float32{3, 4.2, 86.34, 4325.24363452},
			decPlaces: -1,
			sep:       "SEP",
			want:      "SEPSEPSEP",
		},
		{
			name:      "empty both",
			vals:      []float32{3, 4.2, 86.34, 4325.24363452},
			decPlaces: -1,
			sep:       "",
			want:      "",
		},
		{
			name:      "normal example",
			vals:      []float32{3, 4.2, 86.34, 4325.24363452, -1},
			decPlaces: 4,
			sep:       ", ",
			want:      "3.0000, 4.2000, 86.3400, 4325.2437, -1.0000",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Fjoin[float32](test.vals, test.sep, func(e float32) string {
				if test.decPlaces == -1 {
					return ""
				}
				fmtStr := "%." + fmt.Sprintf("%d", test.decPlaces) + "f"
				return fmt.Sprintf(fmtStr, e)
			})
			if got != test.want {
				t.Errorf("Fjoin() = %q, want %q", got, test.want)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		name      string
		intIn     int64
		floatIn   float64
		wantInt   int64
		wantFloat float64
	}{
		{
			name:      "0",
			intIn:     0,
			floatIn:   0.0,
			wantInt:   0,
			wantFloat: 0.0,
		},
		{
			name:      "negative",
			intIn:     -236345,
			floatIn:   -23532.523,
			wantInt:   236345,
			wantFloat: 23532.523,
		},
		{
			name:      "positive",
			intIn:     236345,
			floatIn:   23532.523,
			wantInt:   236345,
			wantFloat: 23532.523,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotInt := Abs[int64](test.intIn)
			gotFloat := Abs[float64](test.floatIn)
			if gotInt != test.wantInt || gotFloat != test.wantFloat {
				t.Errorf("Abs() = {%d, %f}, want {%d, %f}", gotInt, gotFloat, test.wantInt, test.wantFloat)
			}
		})
	}
}

func TestFmax(t *testing.T) {
	tests := []struct {
		name string
		vals []string
		want int
	}{
		{
			name: "empty -> zero value for slice, not applicable for non-slice",
			vals: nil,
			want: 0,
		},
		{
			name: "single val -> the single val for slice, not applicable for non-slice",
			vals: []string{"5"},
			want: 5,
		},
		{
			name: "2 vals -> the greater val",
			vals: []string{"5", "6"},
			want: 6,
		},
		{
			name: "2 vals but reversed -> still the greater val",
			vals: []string{"6", "5"},
			want: 6,
		},
		{
			name: "3 vals -> the greatest val",
			vals: []string{"5", "6", "7"},
			want: 7,
		},
		{
			name: "3 vals, different order -> still the greatest val",
			vals: []string{"5", "7", "6"},
			want: 7,
		},
		{
			name: "3 vals, another different order -> still the greatest val",
			vals: []string{"7", "5", "6"},
			want: 7,
		},
		{
			name: "N vals",
			vals: []string{"3", "2", "34573", "2", "5", "4574", "2345234", "5"},
			want: 2345234,
		},
		{
			name: "N vals with negatives",
			vals: []string{"3", "2", "34573", "2", "5", "4574", "-2345234", "5"},
			want: 34573,
		},
	}

	f := func(e string) int {
		i, err := strconv.Atoi(e)
		if err != nil {
			t.Fatal(err)
		}
		return i
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotSlice := FsliceMax[string, int](test.vals, f)
			if gotSlice != test.want {
				t.Errorf("FsliceMax() = %d, want %d", gotSlice, test.want)
			}

			if len(test.vals) > 1 {
				got := Fmax[string, int](f, test.vals[0], test.vals[1], test.vals[2:]...)
				if got != test.want {
					t.Errorf("Fmax() = %d, want %d", got, test.want)
				}
			}
		})
	}
}

func TestFmin(t *testing.T) {
	tests := []struct {
		name string
		vals []string
		want int
	}{
		{
			name: "empty -> zero value for slice, not applicable for non-slice",
			vals: nil,
			want: 0,
		},
		{
			name: "single val -> the single val for slice, not applicable for non-slice",
			vals: []string{"5"},
			want: 5,
		},
		{
			name: "2 vals -> the lesser val",
			vals: []string{"5", "6"},
			want: 5,
		},
		{
			name: "2 vals but reversed -> still the lesser val",
			vals: []string{"6", "5"},
			want: 5,
		},
		{
			name: "3 vals -> the least val",
			vals: []string{"5", "6", "7"},
			want: 5,
		},
		{
			name: "3 vals, different order -> still the least val",
			vals: []string{"5", "7", "6"},
			want: 5,
		},
		{
			name: "3 vals, another different order -> still the least val",
			vals: []string{"7", "5", "6"},
			want: 5,
		},
		{
			name: "N vals",
			vals: []string{"3", "2", "34573", "2", "5", "4574", "2345234", "5"},
			want: 2,
		},
		{
			name: "N vals with negatives",
			vals: []string{"3", "2", "34573", "-2", "5", "-4574", "-2345234", "5"},
			want: -2345234,
		},
	}

	f := func(e string) int {
		i, err := strconv.Atoi(e)
		if err != nil {
			t.Fatal(err)
		}
		return i
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotSlice := FsliceMin[string, int](test.vals, f)
			if gotSlice != test.want {
				t.Errorf("FsliceMin() = %d, want %d", gotSlice, test.want)
			}

			if len(test.vals) > 1 {
				got := Fmin[string, int](f, test.vals[0], test.vals[1], test.vals[2:]...)
				if got != test.want {
					t.Errorf("Fmin() = %d, want %d", got, test.want)
				}
			}
		})
	}
}

func TestLongest(t *testing.T) {
	tests := []struct {
		name       string
		strs       []string
		slices     [][]int
		wantStrs   int
		wantSlices int
	}{
		{
			name:       "empty",
			strs:       nil,
			slices:     nil,
			wantStrs:   0,
			wantSlices: 0,
		},
		{
			name:       "one item - and it's empty",
			strs:       []string{""},
			slices:     [][]int{nil},
			wantStrs:   0,
			wantSlices: 0,
		},
		{
			name:       "multiple",
			strs:       []string{"a", "", "asdfsadf", "4"},
			slices:     [][]int{{1}, nil, {}, {1, 2, 3, 4}, {5}},
			wantStrs:   8,
			wantSlices: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotStrs := Longest[string](test.strs)
			gotSlices := Longest[int, []int](test.slices)

			if gotStrs != test.wantStrs || gotSlices != test.wantSlices {
				t.Errorf("Longest() = {%d, %d}, want {%d, %d}", gotStrs, gotSlices, test.wantStrs, test.wantSlices)
			}
		})
	}
}

func TestPadding(t *testing.T) {
	tests := []struct {
		name        string
		padder      string
		repititions int
		want        string
	}{
		{
			name:        "empty padder",
			padder:      "",
			repititions: 500,
			want:        "",
		},
		{
			name:        "0 repititions",
			padder:      "asdf",
			repititions: 0,
			want:        "",
		},
		{
			name:        "1 repitition",
			padder:      "asdf",
			repititions: 1,
			want:        "asdf",
		},
		{
			name:        "multiple repitition",
			padder:      "asdf",
			repititions: 5,
			want:        "asdfasdfasdfasdfasdf",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Padding(test.padder, test.repititions)
			if got != test.want {
				t.Errorf("Padding(%q, %d) = %q, want %q", test.padder, test.repititions, got, test.want)
			}
		})
	}
}

func TestPadToLeft(t *testing.T) {
	tests := []struct {
		name   string
		str    string
		padder string
		chars  int
		want   string
	}{
		{
			name:   "empty str",
			str:    "",
			padder: "a",
			chars:  1,
			want:   "a",
		},
		{
			name:   "want no chars",
			str:    "",
			padder: "a",
			chars:  0,
			want:   "",
		},
		{
			name:   "chars < len(str)",
			str:    "asdf",
			padder: "p",
			chars:  1,
			want:   "asdf",
		},
		{
			name:   "chars == len(str)",
			str:    "asdf",
			padder: "p",
			chars:  4,
			want:   "asdf",
		},
		{
			name:   "chars > len(str)",
			str:    "asdf",
			padder: "p",
			chars:  8,
			want:   "ppppasdf",
		},
		{
			name:   "chars > len(str) but padder doesn't divide evenly",
			str:    "asdf",
			padder: "pie",
			chars:  8,
			want:   "pieasdf",
		},
		{
			name:   "chars > len(str) and multi-char padder divides evenly",
			str:    "asdf",
			padder: "pies",
			chars:  8,
			want:   "piesasdf",
		},
		{
			name:   "chars > len(str) but padder doesn't divide evenly (but the other way)",
			str:    "asdf",
			padder: "pies2",
			chars:  8,
			want:   "asdf",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := PadToLeft(test.str, test.padder, test.chars)
			if got != test.want {
				t.Errorf("PadToLeft(%q, %q, %d) = %q, want %q", test.str, test.padder, test.chars, got, test.want)
			}
		})
	}
}

func TestPadToRight(t *testing.T) {
	tests := []struct {
		name   string
		str    string
		padder string
		chars  int
		want   string
	}{
		{
			name:   "empty str",
			str:    "",
			padder: "a",
			chars:  1,
			want:   "a",
		},
		{
			name:   "want no chars",
			str:    "",
			padder: "a",
			chars:  0,
			want:   "",
		},
		{
			name:   "chars < len(str)",
			str:    "asdf",
			padder: "p",
			chars:  1,
			want:   "asdf",
		},
		{
			name:   "chars == len(str)",
			str:    "asdf",
			padder: "p",
			chars:  4,
			want:   "asdf",
		},
		{
			name:   "chars > len(str)",
			str:    "asdf",
			padder: "p",
			chars:  8,
			want:   "asdfpppp",
		},
		{
			name:   "chars > len(str) but padder doesn't divide evenly",
			str:    "asdf",
			padder: "pie",
			chars:  8,
			want:   "asdfpie",
		},
		{
			name:   "chars > len(str) and multi-char padder divides evenly",
			str:    "asdf",
			padder: "pies",
			chars:  8,
			want:   "asdfpies",
		},
		{
			name:   "chars > len(str) but padder doesn't divide evenly (but the other way)",
			str:    "asdf",
			padder: "pies2",
			chars:  8,
			want:   "asdf",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := PadToRight(test.str, test.padder, test.chars)
			if got != test.want {
				t.Errorf("PadToRight(%q, %q, %d) = %q, want %q", test.str, test.padder, test.chars, got, test.want)
			}
		})
	}
}

func TestPadToPadding(t *testing.T) {
	tests := []struct {
		name   string
		str    string
		padder string
		chars  int
		want   string
	}{
		{
			name:   "empty str",
			str:    "",
			padder: "a",
			chars:  1,
			want:   "a",
		},
		{
			name:   "want no chars",
			str:    "",
			padder: "a",
			chars:  0,
			want:   "",
		},
		{
			name:   "chars < len(str)",
			str:    "asdf",
			padder: "p",
			chars:  1,
			want:   "",
		},
		{
			name:   "chars == len(str)",
			str:    "asdf",
			padder: "p",
			chars:  4,
			want:   "",
		},
		{
			name:   "chars > len(str)",
			str:    "asdf",
			padder: "p",
			chars:  8,
			want:   "pppp",
		},
		{
			name:   "chars > len(str) but padder doesn't divide evenly",
			str:    "asdf",
			padder: "pie",
			chars:  8,
			want:   "pie",
		},
		{
			name:   "chars > len(str) and multi-char padder divides evenly",
			str:    "asdf",
			padder: "pies",
			chars:  8,
			want:   "pies",
		},
		{
			name:   "chars > len(str) but padder doesn't divide evenly (but the other way)",
			str:    "asdf",
			padder: "pies2",
			chars:  8,
			want:   "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := padToPadding(test.str, test.padder, test.chars)
			if got != test.want {
				t.Errorf("padToPadding(%q, %q, %d) = %q, want %q", test.str, test.padder, test.chars, got, test.want)
			}
		})
	}
}
