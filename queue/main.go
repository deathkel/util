package main

import (
    "github.com/deathkel/util/queue/queue"
    "context"
    "fmt"
    "time"
)

type bar struct{
    v int
}


func main() {
    q := queue.NewQ()
    for i := 0; i < 10; i++ {
        data := &bar{i}
        q.Add(data)
        q.Add(data)
    }
    
    for i:=0; i< 5; i++{
        //5个消费者
        go func() {
            ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
            q.Get(ctx, func(data interface{}) bool {
                fmt.Println( " data:", data)
                return true
            })
        }()
    }
    q.Close()
    time.Sleep(time.Second * 5)
}
