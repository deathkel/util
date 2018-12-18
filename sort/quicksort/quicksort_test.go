package quicksort

import (
    "reflect"
    "testing"
)

func TestQuicksortASC(t *testing.T) {
    type args struct {
        list []int
    }
    tests := []struct {
        name string
        args args
        want []int
    }{
        // TODO: Add test cases.
        {args:args{[]int{123, 1, 213, 21, 42, 1321}}, want: []int{1, 21, 42, 123, 213, 1321}},
        //{args:args{[]int{1,32,4,1,23,42,141,2}}, want: []int{1,1,2,4,23,32,42,141}},
        //{args:args{[]int{0,2,0}}, want: []int{0,0,2}},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := QuicksortASC(tt.args.list); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("QuicksortASC() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestQuicksortDESC(t *testing.T) {
    type args struct {
        list []int
    }
    tests := []struct {
        name string
        args args
        want []int
    }{
        // TODO: Add test cases.
        {args:args{[]int{123, 1, 213, 21, 42, 1321}}, want: []int{1321,213,123,42,21,1}},
        {args:args{[]int{1,32,4,1,23,42,141,2}}, want: []int{141,42, 32,23,4,2,1,1}},
        {args:args{[]int{0,2,0}}, want: []int{2,0,0}},
        {args:args{[]int{0,0,2}}, want: []int{2,0,0}},
        {args:args{[]int{0,2}}, want: []int{2,0}},
        {args:args{[]int{2,0}}, want: []int{2,0}},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := QuicksortDESC(tt.args.list); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("QuicksortASC() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestThreeWayQuickSort(t *testing.T) {
    
    type args struct {
        list []int
    }
    tests := []struct {
        name string
        args args
        want []int
    }{
        // TODO: Add test cases.
        {args:args{[]int{4, 1, 4, 4, 3, 5,8,3,3,4}}, want: []int{1,3,3,3,4,4,4,4,5,8}},
        {args:args{[]int{1,32,4,1,23,42,141,2}}, want: []int{1,1,2,4,23,32,42,141}},
        {args:args{[]int{0,2,0}}, want: []int{0,0,2}},
        {args:args{[]int{0,0,2}}, want: []int{0,0,2}},
        {args:args{[]int{2,0,0}}, want: []int{0,0,2}},
        {args:args{[]int{2,2,0}}, want: []int{0,2,2}},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ThreeWayQuickSort(tt.args.list, 0, len(tt.args.list) - 1); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ThreeWayQuickSort() = %v, want %v", got, tt.want)
            }
        })
    }
}