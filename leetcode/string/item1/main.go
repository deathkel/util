package main

import (
    "fmt"
    //"math"
)

func main() {
    find, l := lengthOfLongestSubstring("abcabcdbb")
    fmt.Println(string(find), l)
}

func lengthOfLongestSubstring(s string) (find []rune, maxLength int) {
    unionMap := map[rune]bool{}
    i, j := 0, 0
    lenS := len(s)
    for i < lenS && j < lenS {
        v := rune(s[j])
        if _, ok := unionMap[v]; !ok {
            unionMap[v] = true
            j++
            if j-i > maxLength {
                maxLength = j - i
                find = []rune(s[i:j])
            }
        } else {
            delete(unionMap, rune(s[i]))
            i++
        }
    }
    return
}
