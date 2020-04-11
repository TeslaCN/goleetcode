package p151_1

// Space: O(1)
func reverseWords(s string) string {
	b := []byte(s)
	n := len(b)
	if n <= 0 {
		return ""
	}
	src, target := 0, 0

	// skip blank prefix
	for ; src < n && b[src] == ' '; src++ {
	}

	blank := 0
	valid := 0
	for src < n {
		if b[src] == ' ' && blank > 0 {
			// duplicated
			src++
		} else {
			if b[src] == ' ' {
				blank++
			} else {
				blank = 0
			}
			// move
			b[target] = b[src]
			src++
			target++
			valid++
		}
	}
	if valid < 1 {
		return ""
	}
	if b[valid-1] == ' ' {
		valid--
	}
	b = b[:valid]
	n = len(b)

	leftStart, leftEnd, rightStart, rightEnd := 0, 0, n, n
	for leftEnd < rightStart {
		// find left end
		for leftEnd < n && b[leftEnd] != ' ' {
			leftEnd++
		}
		// find right start
		for rightStart > 0 && b[rightStart-1] != ' ' {
			rightStart--
		}

		leftLen, rightLen := leftEnd-leftStart, rightEnd-rightStart
		if leftLen == rightLen {
			d := rightStart - leftStart
			for i := leftStart; i < leftEnd; i++ {
				b[i], b[i+d] = b[i+d], b[i]
			}
		} else if leftLen < rightLen {
			sub := rightLen - leftLen
			for times := 0; times < sub; times++ {
				tmp := b[rightEnd-1]
				for i := rightEnd - 1; i > leftEnd; i-- {
					b[i] = b[i-1]
				}
				b[leftEnd] = tmp
			}
			d := rightStart - leftStart + sub
			for i := leftStart; i < leftEnd; i++ {
				b[i], b[i+d] = b[i+d], b[i]
			}
		} else {
			sub := leftLen - rightLen
			for times := 0; times < sub; times++ {
				tmp := b[leftStart]
				for i := leftStart; i < rightStart-1; i++ {
					b[i] = b[i+1]
				}
				b[rightStart-1] = tmp
			}
			d := rightStart - leftStart
			for i := rightStart; i < rightEnd; i++ {
				b[i], b[i-d] = b[i-d], b[i]
			}
		}
		leftStart = leftStart + rightLen + 1
		leftEnd = leftStart + 1
		rightEnd = rightEnd - leftLen - 1
		rightStart = rightEnd - 1
	}

	return string(b)
}

/*
执行用时 : 436 ms , 在所有 Go 提交中击败了 11.04% 的用户
内存消耗 : 3.5 MB , 在所有 Go 提交中击败了 93.75% 的用户
*/
