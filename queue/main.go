package main

import (
    "github.com/deathkel/util/queue/queue"
    "context"
    "fmt"
    "time"
)

func main() {
    q := queue.NewQ()
    for i := 0; i < 10; i++ {
        data := i + 10
        q.Add(i, data)
        q.Add(i, data)
    }
    
    for i:=0; i< 5; i++{
        //5个消费者
        go func() {
            ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
            q.Get(ctx, func(key interface{}, data interface{}) bool {
                fmt.Println("get key:", key, " data:", data)
                return true
            })
        }()
    }
    q.Close()
    time.Sleep(time.Second * 5)
}
