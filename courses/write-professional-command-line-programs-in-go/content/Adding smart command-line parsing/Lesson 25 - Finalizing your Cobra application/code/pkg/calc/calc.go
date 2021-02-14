package calc

const (
	maxUint=^uint(0)
	maxInt = int(maxUint >> 1)
	minInt = -maxInt - 1
)

func checkAdd(a, b int) {
	if a > 0 {
		if b > maxInt - a {
			panic("overflow!")
		}
	} else {
		if b < minInt - a {
			panic("underflow!")
		}
	}
}

func checkSub(a, b int) {
	if a > 0 {
		if b < minInt + a {
			panic("overflow!")
		}
	} else {
		if b < a - maxInt {
			panic("underflow!")
		}
	}
}

func Add(a, b int, check bool) int {
	if check {
		checkAdd(a, b)
	}

	return a + b
}

func Subtract(a, b int, check bool) int {
	if check {
		checkSub(a, b)
	}
	return a - b
}

