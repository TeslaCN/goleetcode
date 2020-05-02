package p438_1

import "math/big"

var FindAnagrams = findAnagrams

var hash = []*big.Int{
	big.NewInt(2),
	big.NewInt(3),
	big.NewInt(5),
	big.NewInt(7),
	big.NewInt(11),
	big.NewInt(13),
	big.NewInt(17),
	big.NewInt(19),
	big.NewInt(23),
	big.NewInt(29),
	big.NewInt(31),
	big.NewInt(37),
	big.NewInt(41),
	big.NewInt(43),
	big.NewInt(47),
	big.NewInt(53),
	big.NewInt(59),
	big.NewInt(61),
	big.NewInt(67),
	big.NewInt(71),
	big.NewInt(73),
	big.NewInt(79),
	big.NewInt(83),
	big.NewInt(89),
	big.NewInt(97),
	big.NewInt(101),
}

func findAnagrams(s string, p string) []int {
	sLen, pLen := len(s), len(p)
	if sLen < 1 || pLen < 1 || sLen < pLen {
		return []int{}
	}
	result := make([]int, 0, sLen-pLen)
	hashP := big.NewInt(1)
	for i := 0; i < pLen; i++ {
		hashP = hashP.Mul(hashP, hash[p[i]-0x61])
	}
	hashCode := big.NewInt(1)
	for i := 0; i < pLen; i++ {
		hashCode = hashCode.Mul(hashCode, hash[s[i]-0x61])
	}
	if hashCode.Cmp(hashP) == 0 {
		result = append(result, 0)
	}
	for i := 0; i < sLen-pLen; i++ {
		hashCode = hashCode.Div(hashCode, hash[s[i]-0x61])
		hashCode = hashCode.Mul(hashCode, hash[s[i+pLen]-0x61])
		if hashCode.Cmp(hashP) == 0 {
			result = append(result, i+1)
		}
	}
	return result
}

/*
执行用时 : 248 ms , 在所有 Go 提交中击败了 7.34% 的用户
内存消耗 : 5.3 MB , 在所有 Go 提交中击败了 100.00% 的用户
*/
