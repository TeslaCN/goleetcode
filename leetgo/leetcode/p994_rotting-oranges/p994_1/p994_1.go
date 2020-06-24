package p994_1

func orangesRotting(grid [][]int) int {
	minute := 0

	rot := make(map[[2]int]struct{})
	fresh := make(map[[2]int]struct{})
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			switch grid[i][j] {
			case 0:
			case 1:
				fresh[[2]int{i, j}] = struct{}{}
			case 2:
				rot[[2]int{i, j}] = struct{}{}
			}
		}
	}

	for len(rot) > 0 && len(fresh) > 0 {
		newRot := make(map[[2]int]struct{})
		for pos := range rot {
			x, y := pos[0], pos[1]
			down := [2]int{x + 1, y}
			if _, exists := fresh[down]; exists {
				newRot[down] = struct{}{}
				delete(fresh, down)
			}
			up := [2]int{x - 1, y}
			if _, exists := fresh[up]; exists {
				newRot[up] = struct{}{}
				delete(fresh, up)
			}
			right := [2]int{x, y + 1}
			if _, exists := fresh[right]; exists {
				newRot[right] = struct{}{}
				delete(fresh, right)
			}
			left := [2]int{x, y - 1}
			if _, exists := fresh[left]; exists {
				newRot[left] = struct{}{}
				delete(fresh, left)
			}
			delete(rot, pos)
		}
		for pos := range newRot {
			rot[pos] = struct{}{}
		}
		minute++
	}

	if len(fresh) > 0 {
		return -1
	} else {
		return minute
	}
}
