package util

// BinaryToIntArray converts a binary number as a string to an array, where
// the index represents the bit number. The '0' character will be treated as
// 0, all other characters will be treated as 1
//
// For Example
//      1) "00101" -> [1, 0, 1, 0, 0]
func BinaryToIntArray(bin string) []int {
	var ret []int = make([]int, len(bin))
	for i, c := range bin {
		switch c {
		case '0':
			ret[len(bin)-1-i] = 0
		default:
			ret[len(bin)-1-i] = 1
		}
	}
	return ret
}

// IntArrayToBinary converts an indexed array of binary value to a string
// representation.
//
// For Example
//      1) [1, 0, 1, 0, 0] -> "00101"
func IntArrayToBinary(vals []int) string {
	bin := ""
	for i := len(vals) - 1; i >= 0; i-- {
		switch vals[i] {
		case 0:
			bin += "0"
		default:
			bin += "1"
		}
	}
	return bin
}
