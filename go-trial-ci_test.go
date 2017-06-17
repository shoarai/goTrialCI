package gotrialci_test

import (
	"testing"

	gotrialci "github.com/shoarai/goTrialCI"
)

func TestAdd(t *testing.T) {
	for _, test := range []struct {
		vals []int
		want int
	}{
		{[]int{1, -1}, 1},
		// {[]int{1, -1}, 2}, // Test fail
	} {
		actual := gotrialci.AddPlus(test.vals...)
		if actual != test.want {
			t.Errorf("AddPlus(%v) = %d, want %d", test.vals, actual, test.want)
		}
	}
}
