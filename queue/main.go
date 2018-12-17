package main

import (
    "github.com/deathkel/util/queue/queue"
    "context"
    "fmt"
    "time"
    "math/rand"
)

type bar struct{
    v int
}


func main() {
    q := queue.NewQ(1)
    for i := 0; i < 1000; i++ {
        go func() {
            ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
            for {
                a := rand.Int()
                res := q.Add(ctx, a)
                if res == false{
                    time.Sleep(time.Second)
                }else{
                    time.Sleep(time.Millisecond)
                }
            }
        }()
    }
    
    for i:=0; i< 50; i++{
        //5个消费者
        go func() {
            ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
            q.Get(ctx, func(data interface{}) bool {
                fmt.Println( " data:", data)
                return false
            })
        }()
    }
    time.Sleep(time.Second * 5)
    fmt.Println("5s after")
    q.Close()
    time.Sleep(time.Second * 2)
}
