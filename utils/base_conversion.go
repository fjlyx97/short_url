package utils

import (
	"math"
	"strings"
)

var (
	baseMapEncode = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	baseMapDecode = map[string]int64{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "a": 10, "b": 11, "c": 12, "d": 13, "e": 14, "f": 15, "g": 16, "h": 17, "i": 18, "j": 19, "k": 20, "l": 21, "m": 22, "n": 23, "o": 24, "p": 25, "q": 26, "r": 27, "s": 28, "t": 29, "u": 30, "v": 31, "w": 32, "x": 33, "y": 34, "z": 35, "A": 36, "B": 37, "C": 38, "D": 39, "E": 40, "F": 41, "G": 42, "H": 43, "I": 44, "J": 45, "K": 46, "L": 47, "M": 48, "N": 49, "O": 50, "P": 51, "Q": 52, "R": 53, "S": 54, "T": 55, "U": 56, "V": 57, "W": 58, "X": 59, "Y": 60, "Z": 61}
)

// 10进制转62进制
func Base10ToBase62(number int64) string {
	builder := strings.Builder{}
	if number == 0 {
		builder.WriteByte('0')
	}
	for number > 0 {
		index := number % 62
		number /= 62
		builder.WriteByte(baseMapEncode[index])
	}
	// 反转得到正确的数值
	s := builder.String()
	l := builder.Len() - 1
	builder.Reset()
	for i, _ := range s {
		builder.WriteByte(s[l-i])
	}
	return builder.String()
}

func Base62ToBase10(number string) int64 {
	var result int64 = 0
	l := len(number) - 1
	for index, c := range []byte(number) {
		result += baseMapDecode[string(c)] * int64(math.Pow(62, float64(l-index)))
	}
	return result
}
