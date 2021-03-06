package gotrialci

import "os"

// AddPlus adds only plus numbers
func AddPlus(n ...int) int {
	var num int
	for _, nn := range n {
		if nn < 0 {
			continue
		}
		num += nn
	}
	return num
}

// IsPathExist returns whether a path exists or not.
func IsPathExist(path string) bool {
	_, e := os.Stat(path)
	return e == nil
}
