package linklist

import "fmt"

type LinkList struct {
	head   *ListNode // 链表的头，指向一个链表节点
	length int       // 当前链表的节点个数
}

// CreateLinkList 初始化单向链表
func CreateLinkList() *LinkList {
	return &LinkList{
		head:   &ListNode{0, nil},
		length: 0,
	}
}

// PrintLint 打印链表
func (linklist *LinkList) PrintLint() {
	cur := linklist.head.GetNext()
	for nil != cur {
		fmt.Printf("%v->", cur.GetValue())
		cur = cur.GetNext()
	}
	fmt.Println()
}
