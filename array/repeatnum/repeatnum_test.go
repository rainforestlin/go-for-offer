package main

import "testing"

func TestFindRepeatNum(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{
			[]int{2, 3, 1, 0, 2, 5},
			2,
		},
		{
			[]int{1, 2, 3, 4, 5, 1},
			1,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 11},
			11,
		},
	}
	for _, test := range tests {
		resAgain := findRepeatNum(test.input)
		if resAgain != test.output {
			t.Errorf("expected get %d,but get %d", test.output, resAgain)
		}
		res := findRepeatNumByMap(test.input)
		if res != test.output {
			t.Errorf("by map expected get %d,but get %d", test.output, res)
		}
		res = findRepeatNumBySort(test.input)
		if res != test.output {
			t.Errorf("by sort expected get %d,but get %d", test.output, res)
		}

	}
}
