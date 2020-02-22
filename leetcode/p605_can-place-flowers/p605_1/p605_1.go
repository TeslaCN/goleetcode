package p605_1

var CanPlaceFlowers = canPlaceFlowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n <= 0 {
		return true
	}
	flowerbed = append(flowerbed, 0, 0)
	copy(flowerbed[1:], flowerbed[:len(flowerbed)-2])
	for i := 1; n > 0 && i < len(flowerbed)-1; i++ {
		if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n--
		}
	}
	return n <= 0
}
