package linklist

import "errors"

// Insert
// i 是节点索引，从0开始
// node是需要插入的节点
func (linlist *LinkList) Insert(i int, node *ListNode) error {
	//发现unit 和int是不能比较大小的 会报类型不匹配
	//校验参数有效性
	if i < 0 || node == nil || i > linlist.length {
		return errors.New("节点为空或者是越界")
	}
	// 从head通过循环依次定位到索引i，获取单向列表的头部信息
	curNode := (*linlist).head
	// 并从头开始通过for循环定位i
	for j := 0; j < i; j++ {
		curNode = curNode.GetNext()
	}
	//将需要新插入的节点挂接到定位到的节点next上，构成新的链
	node.next = curNode.GetNext()
	curNode.next = node
	linlist.length++
	return nil
}
