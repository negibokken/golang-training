package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount 2.6.2 example
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// MyPopCount Practice 2.3
func MyPopCount(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>(uint(i)*8))])
	}
	return sum
}

// DirtyPopCount Practice 2.4
func DirtyPopCount(x uint64) int {
	sum := 0
	for i := 0; x>>uint(i) != 0; i++ {
		sum += int(x >> uint(i) & 1)
	}
	return sum
}

// BitClearPopCount Practice 2.5
func BitClearPopCount(x uint64) int {
	sum := 0
	for x != 0 {
		x = x & (x - 1)
		sum++
	}
	return sum
}
