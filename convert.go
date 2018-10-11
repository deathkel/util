package util

import (
    "strings"
    "math"
)

const CODE62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const CODE_LENTH = 62

var EDOC = map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "a": 10, "b": 11, "c": 12, "d": 13, "e": 14, "f": 15, "g": 16, "h": 17, "i": 18, "j": 19, "k": 20, "l": 21, "m": 22, "n": 23, "o": 24, "p": 25, "q": 26, "r": 27, "s": 28, "t": 29, "u": 30, "v": 31, "w": 32, "x": 33, "y": 34, "z": 35, "A": 36, "B": 37, "C": 38, "D": 39, "E": 40, "F": 41, "G": 42, "H": 43, "I": 44, "J": 45, "K": 46, "L": 47, "M": 48, "N": 49, "O": 50, "P": 51, "Q": 52, "R": 53, "S": 54, "T": 55, "U": 56, "V": 57, "W": 58, "X": 59, "Y": 60, "Z": 61,}

/**
 * 编码 整数 为 base62 字符串
 */
func Encode(number int) string {
    return DecimalToAny(number, 62)
}

/**
 * 解码字符串为整数
 */
func Decode(str string) (result int) {
    return AnyToDecimal(str, 62)
}

//10进制转任意进制，num最大16位int，n最大62
func DecimalToAny(num, n int) (string) {
    if num == 0 {
        return "0"
    }
    
    result := make([]byte, 0)
    for num > 0 {
        remain := num % n
        result = append(result, CODE62[remain])
        num = num / n
    }
    return string(result)
}

//任意进制转回10进制，n最大62
func AnyToDecimal(num string, n int) (result int) {
    str := strings.TrimSpace(num)
    for k, v := range []byte(str) {
        result += EDOC[string(v)] * int(math.Pow(float64(n), float64(k)))
    }
    
    return result
}
