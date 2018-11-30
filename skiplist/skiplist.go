package skiplist

import "math/rand"

type Node struct {
    forward []*Node
    key     int
    value   interface{}
}

type SkipList struct {
    header   *Node
    footer   *Node
    length   int
    MaxLevel int
}

func NewNode(lv int, key int, value interface{}) *Node {
    return &Node{
        make([]*Node, lv),
        key,
        value,
    }
}

func NewSkipList() *SkipList {
    sl := &SkipList{}
    sl.header = nil
    sl.footer = nil
    sl.length = 0
    sl.MaxLevel = 0
    return sl
}

func (sl *SkipList) Search(key int) bool {
    current := sl.header
    
    for i := sl.MaxLevel; i >= 0; i-- {
        for current.forward[i].key < key {
            current = current.forward[i]
        }
    }
    
    current = current.forward[0]
    if current.key == key {
        return true
    } else {
        return false
    }
}

func (sl *SkipList) randomLevel() int {
    return rand.Intn(sl.MaxLevel)
}


func (sl *SkipList) Insert(key int, value interface{}) bool {
    x := sl.header
    if x == nil {
        //insert first node
        sl.header = NewNode(sl.randomLevel(), key, value)
        sl.footer = sl.header
        return true
    }
    
    update := make([]*Node, sl.MaxLevel)
    for i := sl.MaxLevel; i >= 0; i-- {
        for x.forward[i].key <= key {
            x = x.forward[i]
        }
        update[i] = x
    }
    
    //key对应的node已存在
    if x.key == key {
        x.value = value
        return false
    }
    
    nodeLevel := sl.randomLevel()
    node := NewNode(nodeLevel, key, value)
    //插入新节点
    for i := nodeLevel; i >= 0; i-- {
        node.forward[i] = update[i].forward[i]
        update[i].forward[i] = node
    }
    
    //是否更改footer
    if key > sl.footer.key {
        sl.footer = node
    }
    
    return true
}
