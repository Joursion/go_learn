package main

import (
	"fmt"
)

type Student struct {
	id int
	name string
}

type Node struct {
	Student
	Next *Node
}

func (head *Node) Create() *Node {
	head = nil
	return head
}

func (p *Node) traverse() {
	for p != nil {
		fmt.Println("%d %s\n", p.id, p.name)
		p = p.Next
	}
}

func (newNode *Node) insert(head *Node) *Node {
	var p0, p1, p2 *Node
	p0 = newNode
	p1 = head
	p2 = p1.Next
	if head == nil {
		p0.Next = nil
		head = p0
	} else {
		// if p0.id <= p1.id {
		// 	p0.Next = p1
		// 	head = p0
		// }
		// while (p2 != nil && p2.id < p0.id) {
		// 	p2 = p2.Next
		// }
		// if (p2.id <= p0.id) {
		// 	p2.Next = p0
		// } else {

		// }

		for (p0.id > p1.id) && p1.Next != nil {
			p2 = p1
			p1 = p1.Next
		}
		if (p0.id > p1.id) {
			p1.Next = p0
		} else {
			if head == p1 {
				head = p0
				p0.Next = p1
			} else {
				p2.Next = p0
				p0.Next = p1
			}
		}
	}
	return head
}

func (ptr *Node) Delete(head *Node) *Node {
	var p1, p2 *Node
	p1 = head

	for ptr.id != p1.id && p1.Next != nil {
		p2 = p1
		p1 = p1.Next
	}

	if p1.id == ptr.id {
		if p1 == head {
			head = p1.Next
		} else {
			p2.Next = p1.Next
		}
	} else {
		fmt.Println("Node not found")
	}
	return head
}

func main() {
	var head *Node
	stu1 := Node{Student{100, "Amy"}, nil}
	head = head.Create()
	head = stu1.insert(head)
}

func test1(a ...interface{}) {
	
}