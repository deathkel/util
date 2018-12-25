package main

import "fmt"

func main() {
    fmt.Println(search([]int{0}, 0))
}

func search(nums []int, target int) int {
    lenNums := len(nums)
    if lenNums > 0 {
        p := nums[lenNums-1]
        breakPoint := foo(nums, 0, lenNums-1, target, p)
        return breakPoint
    }else{
        return -1
    }
}

func foo(nums []int, left, right, target, p int) int {
    if nums[left] == target {
        return left
    }
    if nums[right] == target {
        return right
    }
    if (left + 1 == right || left == right)  && (nums[left] != target && nums[right] != target) {
        return -1
    }
    
    midIndex := (right + left) / 2
    mid := nums[midIndex]
    if target > p {
        //左半
        if mid > p {
            //左半段
            if mid > target {
                right = midIndex
            } else {
                left = midIndex
            }
        } else {
            //右半段
            right = midIndex
        }
    } else {
        //右半
        if mid > p {
            //左半段
            left = midIndex
        } else {
            //右半段
            if mid > target {
                right = midIndex
            } else {
                left = midIndex
            }
        }
    }
    return foo(nums, left, right, target, p)
}
