package gotrialci

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
