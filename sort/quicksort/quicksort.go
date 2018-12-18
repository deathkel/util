package quicksort

import "fmt"

func QuicksortDESC(list []int) []int {
    lenList := len(list)
    if lenList <= 1 {
        return list
    }
    
    standard := list[0]
    start := 1
    end := lenList - 1
    for start < end {
        for list[end] < standard && start < end {
            end --
        }
        
        for list[start] > standard && start < end {
            start ++
        }
        
        if list[end] > list[start] {
            swap(list, start, end)
            start ++
            end --
        }
    }
    
    if list[start] > standard {
        list[0] = list[start]
        list[start] = standard
    } else {
        list[0] = list[start-1]
        list[start-1] = standard
    }
    
    list = append(QuicksortDESC(list[:start]), QuicksortDESC(list[start:])...)
    return list
}

func QuicksortASC(list []int) []int {
    fmt.Println("")
    maxList := len(list)
    if maxList <= 1 {
        //一个数
        return list
    }
    
    start := 1
    end := maxList - 1
    standard := list[0]
    for start < end {
        for list[end] > standard && start < end {
            end --
        }
        
        for list[start] < standard && start < end {
            start ++
        }
        if list[end] < list[start] {
            swap(list, start, end)
            start ++
            end --
        }
    }
    if list[start] < standard {
        list[0] = list[start]
        list[start] = standard
    } else {
        list[0] = list[start-1]
        list[start-1] = standard
    }
    
    list = append(QuicksortASC(list[:start]), QuicksortASC(list[start:])...)
    
    return list
}

func ThreeWayQuickSort(a []int, left, right int) []int {
    threeWayQuickSort(a)
    return a
}

//三路快速排序
func threeWayQuickSort(a []int) {
    if len(a) <= 1 {
        return
    }
    
    i, p, left := 0, 0, 0
    right := len(a) - 1
    j, q := right-1, right-1
    pivot := right
    
    for {
        for i <= j && a[i] <= a[pivot] {
            if a[i] == a[pivot] {
                swap(a, i, p)
                p++
            }
            i ++
        }
        
        for j >= i && a[j] >= a[pivot] {
            if a[j] == a[pivot] {
                swap(a, j, q)
                q --
            }
            j--
        }
        if i >= j {
            break
        }
        swap(a, i, j)
        i++
        j--
    }
    // 此时得到的顺序
    // |等于|小于|大于|等于|
    // left p  i j  q  right
    //将等于的值交换到中间
    //处理的时候用的后面一个值，此处需要退1
    i--
    p--
    for p >= left {
        swap(a, p, i)
        p--
        i--
    }
    j++
    q++
    for q <= right {
        swap(a, q, j)
        q++
        j++
    }
    
    //得到
    //|小于|等于|大于|
    //递归排序 小于和大于的值
    threeWayQuickSort(a[:i+1])
    threeWayQuickSort(a[j-1:])
}

func swap(a []int, i, j int) {
    if i != j {
        temp := a[i]
        a[i] = a[j]
        a[j] = temp
    }
}
