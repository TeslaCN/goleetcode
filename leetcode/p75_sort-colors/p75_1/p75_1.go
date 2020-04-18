package p75_1

var SortColors = sortColors

func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}

	left, right, i := 0, len(nums)-1, -1
	for left < right {
		for left < right && nums[left] == 0 {
			left++
		}
		for right >= left && nums[right] == 2 {
			right--
		}
		if left < right {
			if nums[left] == 2 || nums[right] == 0 {
				nums[left], nums[right] = nums[right], nums[left]
			} else if nums[left]&nums[right] == 1 {
				if i == -1 {
					i = left + 1
				}
				for i <= right {
					switch nums[i] {
					case 0:
						nums[i], nums[left] = nums[left], nums[i]
						left++
					case 2:
						nums[i], nums[right] = nums[right], nums[i]
						right--
					default:
						i++
					}
				}
				if i >= right {
					return
				}
			}
		}
	}
}

/*
执行用时 : 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
内存消耗 : 2 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
