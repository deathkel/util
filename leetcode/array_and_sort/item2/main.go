package main

import "fmt"

func main() {
    input := [][]int{
        {0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
        {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
        {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
        {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
        {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
        {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
        {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
        {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
    }
    
    fmt.Println(maxAreaOfIsland(input))
}

func maxAreaOfIsland(grid [][]int) int {
    
    width := len(grid)
    height := len(grid[0])
    
    max := 0
    for i := 0; i < width; i++ {
        for j := 0; j < height; j++ {
            if grid[i][j] == 1 {
                num := deepSearch(grid, i, j)
                if num > max {
                    max = num
                }
            }
        }
    }
    return max
}

func deepSearch(grid [][]int, i, j int) int {
    if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && grid[i][j] == 1 {
        grid[i][j] = 0
        num := 1 + deepSearch(grid, i-1, j) + deepSearch(grid, i+1, j) + deepSearch(grid, i, j+1) + deepSearch(grid, i, j-1)
        return num
    } else {
        return 0
    }
}
