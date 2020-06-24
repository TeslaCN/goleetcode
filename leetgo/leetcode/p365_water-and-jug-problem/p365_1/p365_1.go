package p365_1

func canMeasureWater(x int, y int, z int) bool {
	switch {
	case z == 0 || x+y == z:
		return true
	case x <= 0 && y <= 0 || x+y < z:
		return false
	case x <= 0:
		return z == y
	case y <= 0:
		return z == x
	}
	if x < y {
		x, y = y, x
	}
	for y != 0 {
		x, y = y, x%y
	}
	return z%x == 0
}
