package popsort

func PopsortASC(list []int) []int {
    lenList := len(list)
    for i := lenList - 1; i > 0; i-- {
        for y := 0; y < i; y++ {
            if (list[y+1] < list[y] ) {
                tmp := list[y]
                list[y] = list[y+1]
                list[y+1] = tmp
            }
        }
    }
    return list
}

func PopsortDESC(list []int) []int  {
    lenList := len(list)
    for i := lenList - 1; i > 0; i-- {
        for y := 0; y < i; y++ {
            if (list[y+1] > list[y] ) {
                tmp := list[y]
                list[y] = list[y+1]
                list[y+1] = tmp
            }
        }
    }
    return list
}


