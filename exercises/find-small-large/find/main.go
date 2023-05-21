package find

func Findsmall(n []int) int {
	xx := n[0]
	for _, v := range n {
		if v < xx {
			xx = v
		}
	}
	return xx
}

func Findlarge(n []int) int {
	xx := n[0]
	for _, v := range n {
		if v > xx {
			xx = v
		}
	}
	return xx
}
