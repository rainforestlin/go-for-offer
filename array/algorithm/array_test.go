package algorithm

import (
	"math"
	"testing"
)

func TestFindNumbersSumEqualN(t *testing.T) {
	tests := []struct {
		input  []int
		sum    int
		output []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			6,
			[]int{0, 4},
		},
	}
	for _, test := range tests {
		actual := FindNumbersSumEqualN(test.input, test.sum)
		if !compareSliceLen(actual, test.output) {
			t.Errorf("FindNumbersSumEqualN error")
		}
	}
}

func TestFindContinuousSequence(t *testing.T) {
	tests := []struct {
		input  int
		output [][]int
	}{
		{100, [][]int{{9, 10, 11, 12, 13, 14, 15, 16}, {18, 19, 20, 21, 22}}},
	}
	for _, test := range tests {
		actual := FindContinuousSequence(test.input)
		for sliceIndex, slice := range actual {
			if !compareSliceLen(slice, test.output[sliceIndex]) {
				t.Errorf("FindContinuousSequence error")
			}
		}
	}
}

func TestMaxSubArraySum(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{2, 4, -7, 5, 2, -1, 2, -4, 3}, 8},
	}
	for _, test := range tests {
		actual := MaxSubArraySum(test.input)
		if actual != test.output {
			t.Errorf("MaxSubArraySum error")
		}
	}
}

func compareSliceLen(sliceA, sliceB []int) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}
	if (sliceA == nil) != (sliceB == nil) {
		return false
	}
	for index, value := range sliceB {
		if value != sliceA[index] {
			return false
		}
	}
	return true

}

func TestFindMedian(t *testing.T) {
	inputs := []struct {
		input  []int
		output float64
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			4.0,
		},{
			[]int{7,6,5,4,3,2,1},
			4.0,
		},
	}
	for _, test := range inputs {
		heap := Constructor()
		for _, num := range test.input {
			heap.AddNum(num)
		}
		result := heap.FindMedian()
		if !isEqual(result, test.output) {
			t.Errorf("FindMedian error")
		}

	}
}

const MIN = 0.000001

func isEqual(a float64, b float64) bool {
	if a > b {
		return math.Dim(a, b) < MIN
	} else {
		return math.Dim(b, a) < MIN
	}
}

func Test(t *testing.T) {
	inputs := []struct {
		input  []int
		result int
	}{
		{[]int{1, 1, 1, 1, 12, 3, 4, 5, 6, 1, 1, 1}, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 8}, -1},
		{[]int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2}, -1},
	}
	for _, input := range inputs {
		result := MoreThanHalfNumByCount(input.input)
		actual := input.result
		if actual != result {
			t.Errorf("More than half num err,get %d,actual %d", result, actual)
		}
	}
}
