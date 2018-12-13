package popsort

import "testing"

func Test_popsortASC(t *testing.T)  {
    a := []int{1,23,2,42,142,1,7,53}
    res := PopsortASC(a)
    println(res)
    t.Error(res)
}


func Test_popsortDESC(t *testing.T)  {
    a := []int{1,23,2,42,142,1,7,53}
    res := PopsortDESC(a)
    println(res)
    t.Error(res)
}