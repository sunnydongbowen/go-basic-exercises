package linklist

// ListNode
type ListNode struct {
	data int
	next *ListNode
}

// CreateListNode 快速构建一个值为value的*ListNode节点
func CreateListNode(value int) *ListNode {
	return &ListNode{
		data: value,
		next: nil,
	}
}

// GetNext 获取下一个节点的值，不要用this这种写法命名接收器
func (listnode *ListNode) GetNext() *ListNode {
	return listnode.next
}

func (listnode *ListNode) GetValue() int {
	return listnode.data
}
