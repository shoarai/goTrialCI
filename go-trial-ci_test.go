package gotrialci_test

import (
	"os"
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

func TestIsPathExist(t *testing.T) {
	path := "test.txt"
	os.Create(path)
	if !gotrialci.IsPathExist(path) {
		t.Errorf("IsPathExist(%q) = %v, want %v", path, false, true)
	}
	os.Remove(path)
}
