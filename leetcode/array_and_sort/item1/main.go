package main

import (
    "fmt"
    //"sort"
    "sort"
    "strconv"
)

func main() {
    //fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
    fmt.Println(threeSum([]int{12, 0, 3, -14, 5, -11, 11, -5, -2, -1, 6, -7, -10, 1, 4, 1, 1, 9, -3, 6, -15, 0, 6, 1, 6, -12, 3, 7, 11, -6, -8, 0, 9, 3, -7, -1, 7, -10, 1, 13, -4, -7, -9, -7, 9, 3, 1, -13, -3, 13, 8, -11, -9, -8, -3, 4, -13, 7, -11, 5, -14, -4, -9, 10, 6, -9, -6, -9, -12, 11, -11, -9, 11, -5, 0, -3, 13, -14, -1, -13, 7, -7, 14, 5, 0, -4, -6, -6, -11, -2, 14, -10, 2, 12, 8, -7, -11, -13, -9, 14, 5, -5, -9, 1, -2, 6, 5, -12, -1, -10, -9, -9, -10, 12, 11}))
}

func threeSum(nums []int) [][]int {
    
    fMap := map[int]int{}
    zMap := map[int]int{}
    fList := []int{}
    zList := []int{}
    for _, v := range nums {
        if v < 0 {
            fMap[v] += 1
            fList = append(fList, v)
        } else {
            zMap[v] += 1
            zList = append(zList, v)
        }
    }
    
    result := [][]int{}
    unionMap := map[string]bool{}
    //负数0个
    if zero := zMap[0]; zero >= 3 {
        result = append(result, []int{0, 0, 0})
    }
    
    //负数1个
    for i := 0; i < len(zList); i++ {
        for j := i + 1; j < len(zList); j++ {
            r := -(zList[i] + zList[j])
            if _, ok := fMap[r]; ok {
                item := []int{r, zList[i], zList[j]}
                sort.Ints(item)
                key := getKey(item)
                if !unionMap[key] {
                    unionMap[key] = true
                    result = append(result, item)
                }
            }
        }
    }
    
    //负数2个
    for i := 0; i < len(fList); i++ {
        for j := i + 1; j < len(fList); j++ {
            r := -(fList[i] + fList[j])
            if _, ok := zMap[r]; ok {
                item := []int{r, fList[i], fList[j]}
                sort.Ints(item)
                key := getKey(item)
                if !unionMap[key] {
                    unionMap[key] = true
                    result = append(result, item)
                }
            }
        }
    }
    
    return result
}

func getKey(a []int) string {
    key := ""
    for _, v := range a {
        key += strconv.Itoa(v) + ","
    }
    return key
}
