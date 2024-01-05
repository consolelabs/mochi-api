package sliceutils

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test_FindDuplications(t *testing.T) {
	type testCase[T comparable] struct {
		arg  []T
		want []T
	}
	var emptySlice []int
	testCases := []testCase[int]{
		{arg: []int{1, 2, 3}, want: emptySlice},
		{arg: []int{1, 2, 3, 2}, want: []int{2}},
		{arg: []int{6, 77, 1, 6, 52, 12, 1}, want: []int{1, 6}},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test_FindDuplications #%d", i+1), func(t *testing.T) {
			got := FindDuplications(tc.arg)
			sort.Slice(got, func(i, j int) bool {
				return got[i] < got[j]
			})
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("FindDuplications() = %v, want %v", got, tc.want)
			}
		})
	}
}
