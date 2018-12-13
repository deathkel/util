package queue

import (
    "time"
    "context"
    "fmt"
)

type Queue interface {
    Add(key interface{}, data interface{}) bool
    Get(context.Context, func(key interface{}, data interface{}) bool) bool
    Close()
    Open()
}

type event struct {
    key       interface{}
    data      interface{}
    createdAt int64
}

type Q struct {
    status bool                        //状态，是否启用
    length int                         //快速获取队列长度
    keyMap map[interface{}]interface{} //用于去重
    events []*event                    //队列内容
}

//构造函数
func NewQ() *Q {
    return &Q{
        status:true,
        keyMap: make(map[interface{}]interface{}),
        length: 0,
    }
}

//加入一个事件到队尾
func (q *Q) Add(key interface{}, data interface{}) bool {
    if q.status == false {
        //防止新数据写入
        fmt.Println("queue is stopped")
        return false
    }
    
    if find := q.keyMap[key]; find != nil {
        fmt.Println("duplicate key", key)
        return false
    }
    
    e := &event{
        key:       key,
        data:      data,
        createdAt: time.Now().Unix(),
    }
    
    q.keyMap[key] = true
    q.events = append(q.events, e)
    q.length ++
    fmt.Println("add key", key, "add data", data, "length",q.length)
    return true
}

//循环消费
func (q *Q) Get(ctx context.Context, callback func(key interface{}, data interface{}) bool) {
    childCtx, cancel := context.WithCancel(ctx)
    defer cancel()
    for {
        if q.status == false{
            cancel()
            fmt.Println("Close func stop Get...")
            return
        }
        
        select {
        case <-ctx.Done():
            cancel()
            fmt.Println("ctx stop Get...")
            return
        default:
            res := q.GetOne(childCtx, callback)
            fmt.Println("getOne res", res)
        }
    }
}

//消费一个事件
func (q *Q) GetOne(ctx context.Context, callback func(key interface{}, data interface{}) bool) (result bool) {
    var e *event

block:
    for {
        if q.status == false {
            fmt.Println("close func stop getOne...")
            return false
        }
        
        select {
        case <-ctx.Done():
            fmt.Println("ctx stop getOne...")
            return false
        default:
            if q.length > 0 {
                e = q.events[0]
                e, q.events = q.events[0], q.events[1:]
                delete(q.keyMap, e.key)
                q.length --
                break block
            }
            time.Sleep(time.Millisecond)
        }
    }
    
    defer func() {
        if err := recover(); err != nil {
            q.Add(e.key, e.data)
            result = false
        }
    }()
    
    callbackRes := callback(e.key, e.data)
    if !callbackRes {
        q.Add(e.key, e.data)
        return false
    }
    
    return true
}

//关闭队列
func (q *Q) Close() {
    if q.status {
        q.status = false
    }
}

//启用队列
func (q *Q) Open() {
    if !q.status {
        q.status = true
    }
}
