package queue

import (
    "time"
    "context"
    "fmt"
    "sync"
    "reflect"
    "crypto/md5"
    "encoding/base64"
    "sensitive/sensitive_vendor/github.com/gin-gonic/gin/json"
    "math"
)

type Queue interface {
    Add(data interface{}) bool
    Get(context.Context, func(data interface{}) bool) bool
    Close()
    Open()
}

type event struct {
    hashKey   string
    data      interface{}
    createdAt int64
}

type Q struct {
    status bool                        //状态，是否启用
    lock sync.RWMutex                  //锁
    length int                         //快速获取队列长度
    unionMap map[string][]*interface{} //hash table
    events []*event                    //队列内容
    sleepRankLimit int                 //消费者休眠等级限制
}

//构造函数
func NewQ() *Q {
    return &Q{
        status:true,
        unionMap: make(map[string][]*interface{}),
        length: 0,
        sleepRankLimit: 10,
    }
}

// 计算hash值
// 不支持 map[interface{}]
func hash(i interface{}) string {
    b, err := json.Marshal(i)
    if err != nil {
        fmt.Println("json marshal fail")
        return ""
    }
    
    h := md5.New()
    return base64.StdEncoding.EncodeToString(h.Sum(b))
}

//是否重复
func (q *Q) isDuplicate(hashStr string, data interface{}) bool {
    if find := q.unionMap[hashStr]; find != nil {
        for _key := range find {
            if reflect.ValueOf(find[_key]).Kind() == reflect.Ptr {
                if reflect.DeepEqual(*find[_key], data){
                    return true
                }
            }else{
                if reflect.DeepEqual(find[_key], data){
                    return true
                }
            }
        }
    }
    return false
}

func (q *Q) enqueue(hashStr string, e *event) {
    q.lock.Lock()
    q.events = append(q.events, e)
    q.unionMap[hashStr] = append(q.unionMap[hashStr], &e.data)
    q.length ++
    q.lock.Unlock()
}

func (q *Q) dequeue() *event {
    q.lock.Lock()
    if q.length == 0 {
        q.lock.Unlock()
        return nil
    }
    e := q.events[0]
    e, q.events = q.events[0], q.events[1:]
    if find := q.unionMap[e.hashKey ]; find != nil {
        for _key := range find {
            if reflect.DeepEqual(find[_key], e.data){
                q.unionMap[e.hashKey] = append(q.unionMap[e.hashKey ][:_key], q.unionMap[e.hashKey][_key:]...)
                break
            }
        }
    }
    q.length --
    q.lock.Unlock()
    return e
}

//加入一个事件到队尾
func (q *Q) Add(data interface{}) bool {
    if q.status == false {
        //防止新数据写入
        fmt.Println("queue is stopped")
        return false
    }
    
    //去重
    hashStr := hash(data)
    if hashStr == "" {
        fmt.Println("get hash string fail")
        return false
    }
    
    if q.isDuplicate(hashStr, data) {
        fmt.Println("duplicate data")
        return false
    }
    e := &event{
        hashKey:       hashStr,
        data:      data,
        createdAt: time.Now().Unix(),
    }
    
    q.enqueue(hashStr, e)
    
    fmt.Println("add data", data,"length",q.length)
    return true
}

//循环消费
func (q *Q) Get(ctx context.Context, callback func(data interface{}) bool) {
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
func (q *Q) GetOne(ctx context.Context, callback func(data interface{}) bool) (result bool) {
    var e *event
    var sleepRank int = 1
    
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
            e = q.dequeue()
            if e != nil{
                break block
            }
            
            //逐渐增加sleep时间，提高性能
            time.Sleep(time.Millisecond * time.Duration(math.Pow(2, float64(sleepRank))))
            if sleepRank < q.sleepRankLimit {
                sleepRank += 1
            }
        }
    }
    
    defer func() {
        if err := recover(); err != nil {
            q.Add(e.data)
            result = false
        }
    }()
    
    callbackRes := callback(e.data)
    if !callbackRes {
        q.Add(e.data)
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
