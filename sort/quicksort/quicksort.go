package quicksort

import "fmt"

func QuicksortDESC(list []int) []int  {
    lenList := len(list)
    if lenList <= 1 {
        return  list
    }
    
    standard := list[0]
    start := 1
    end := lenList - 1
    for start < end {
        for list[end] < standard {
            end --
        }
        
        for list[start] > standard {
            start ++
        }
        
        if end < start {
            t := list[start]
            list[start] = list[end]
            list[end] = t
            start ++
            end --
        }
    }
    
    if list[start] > standard{
        list[0] = list[start]
        list[start] = standard
    }else{
        list[0] = list[start - 1]
        list[start - 1] = standard
    }
    
    list = append(QuicksortDESC(list[:start]), QuicksortDESC(list[start:])...)
    return list
}

func QuicksortASC(list []int) []int {
    fmt.Println("")
    println("")
    maxList := len(list)
    if maxList <= 1 {
        //一个数
        return list
    }
    
    start := 1
    end := maxList - 1
    standard := list[0]
    for start < end {
        for list[end] > standard && start < end{
            end --
        }
        
        for list[start] < standard && start < end {
            start ++
        }
        if list[end] < list[start] {
            t := list[start]
            list[start] = list[end]
            list[end] = t
            start ++
            end --
        }
    }
    if list[start] < standard{
        list[0] = list[start]
        list[start] = standard
    }else{
        list[0] = list[start - 1]
        list[start - 1] = standard
    }

    list = append(QuicksortASC(list[:start]), QuicksortASC(list[start :])...)
    
    return list
}

