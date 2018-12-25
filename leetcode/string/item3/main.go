package main

import (
    "fmt"
    "strconv"
)

func main() {
    fmt.Println(restoreIpAddresses("25525511135"))
    fmt.Println(restoreIpAddresses("25510011135"))
}

func restoreIpAddresses(s string) []string {
    lenS := len(s)
    if lenS < 4 {
        return nil
    }
    all := restore(s, 4)
    return all
}

func restore(s string, k int) []string {
    //最后一段
    if k == 1 {
        if ok(s) {
            return []string{s}
        } else {
            return []string{}
        }
    }
    
    result := []string{}
    //前3段
    for i := 1; i <= 3; i++ {
        cur := s[:i]
        if ok(cur){
            res := restore(s[i:], k - 1)
            for _, item := range res {
                if item != "" {
                    result = append(result, cur + "." + item)
                }
            }
        }
    }
    return result
}

func ok(str string) bool {
    if str == "" {
        return false
    }
    intV, _ := strconv.Atoi(str)
    if intV <= 255 && intV >= 0 && strconv.Itoa(intV) == str {
        return true
    }
    return false
}
