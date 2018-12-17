package queue

import (
    "testing"
    "context"
    "strconv"
    "time"
    "fmt"
)

func Test_hash(t *testing.T) {
    
    data := []interface{}{1,
        "1",
        1.234,
        int64(123),
        struct {
            v1 interface{}
            v2 string
        }{"123", "123"},
        &struct {
            v1 interface{}
            v2 string
        }{"123","123"},
        []int{1, 2, 3},
        &[]int{1, 2, 3},
        []string{"1", "2"},
        &[]string{"1", "2"},
        map[string]string{"a": "a", "b": "b"},
        &map[string]string{"a": "a", "b": "b"},
        map[string]interface{}{"a": "a", "b": "b"},
        &map[string]interface{}{"a": "a", "b": "b"},
    }
    
    for _key := range data {
        res := hash(data[_key])
        if res == "" {
            t.Error("get hash fail")
        }
    }
}

func TestNewQ(t *testing.T) {
    q := NewQ()
    if q == nil {
        t.Error("new Q fail")
    }
}

func TestQ_Add(t *testing.T) {
    q := NewQ()
    for i := 0; i < 5; i++ {
        res := q.Add(i + 10)
        if !res {
            t.Error("add fail")
        }
    }
    
    if q.length != 5 {
        t.Error("length wrong")
    }
    
    if len(q.events) != 5 {
        t.Error("events num wrong")
    }
}

func TestQ_getOne(t *testing.T) {
    q := NewQ()
    
    //push 1 event to queue
    res := q.Add(10)
    if !res {
        t.Error("add fail")
    }
    
    //callback返回false，任务放回队尾
    ch := make(chan bool, 1)
    go func() {
        ctx1, _ := context.WithCancel(context.Background())
        res = q.getOne(ctx1, func(data interface{}) bool {
            return false
        })
        ch <- res
    }()
    fmt.Println(q.unionMap)
    if <-ch {
        t.Error("Get return true, should return false")
    }
    if q.length != 1 {
        t.Error("callback repush data to queue fail, length should be 1, get " + strconv.Itoa(q.length))
    }
    
    //callback返回true,队列消费正常
    go func() {
        ctx2, _ := context.WithCancel(context.Background())
        res = q.getOne(ctx2, func(data interface{}) bool {
            return true
        })
        ch <- res
    }()
    
    if ! <- ch {
        t.Error("Get return false, should return true")
    }
    if q.length != 0 {
        t.Error("queue length error shoud be 0, get " + strconv.Itoa(q.length))
    }
    
    //消费者取消
    go func() {
        ctx3, _ := context.WithTimeout(context.Background(), time.Second)
        res = q.getOne(ctx3, func(data interface{}) bool {
            t.Error("should block")
            return false
        })
        ch <- res
    }()
    
    
    if <- ch {
        t.Error("consumer timeout shoud return false")
    }
}

func TestQ_Get(t *testing.T) {
    q := NewQ()
    
    for i := 0; i < 5; i++ {
        q.Add(i)
    }
    
    ch := make(chan bool)
    go func() {
        defer func() {
            ch <- false
        }()
        i := 0
        ctx3, _ := context.WithTimeout(context.Background(), time.Second)
        q.Get(ctx3, func(data interface{}) bool {
            if i > 5 {
                t.Error("should be lock")
                ch <- true
            }
            i++
            return true
        })
    }()
    
    <-ch
    return
}

func TestQ_Close(t *testing.T) {
    q := NewQ()
    for i := 0; i < 5; i++ {
        q.Add(i)
    }
    
    q.Close()
    
    if q.status != false {
        t.Error("Colse fail")
    }
}
