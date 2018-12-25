package main


import (
    "strings"
    "fmt"
)

func main()  {
   
    fmt.Println( simplifyPath("/a/./ba/../../c/"))
}

func simplifyPath(path string) string {
    strArr := strings.Split(path, "/")
    stack := []string{}
    
    for _, str := range strArr {
        if str == ".." {
            //pop
            lenStack := len(stack)
            if lenStack > 0 {
                stack = stack[:lenStack - 1]
            }
            
        } else if str == "." || str == "" {
        
        }else{
            stack = append(stack, str)
        }
    }
    
    res := "/"
    lenStack := len(stack)
    for k, v := range stack {
        if k < lenStack - 1{
            res += v+ "/"
        }else{
            res += v
        }
        
    }
    
    return res
}