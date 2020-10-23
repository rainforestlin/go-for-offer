package algorithm

import "testing"

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
