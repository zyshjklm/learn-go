package main

import "fmt"

// Student info
type Student struct {
	ID   int
	Name string
}

// Node for link
type Node struct {
	Val  Student
	Next *Node
	// need use *(pointer)
	// if use Node, error: invalid recursive type Node
}

func main() {
	nodeA := Node{Val: Student{ID: 1, Name: "jungle"}}
	nodeB := Node{Val: Student{ID: 2, Name: "tom"}}
	nodeC := Node{Val: Student{ID: 3, Name: "jack"}}
	nodeD := Node{Val: Student{ID: 4, Name: "alice"}}

	nodeA.Next = &nodeB
	nodeB.Next = &nodeC
	nodeC.Next = &nodeD

	p := &nodeA
	printList(p)
	fmt.Println("---")
	p = reverseList(p)
	printList(p)
}

func reverseList(p *Node) *Node {
	// | pre | current | next | ......
	var pre *Node
	for p != nil {
		// 暂存 p的下一个节点
		next := p.Next
		// 让当前节点p的下一个Node为pre
		p.Next = pre
		// 更新pre为当前节点
		pre = p
		// 继续迭代， 变更p为其原来的next
		p = next
	}
	// 此时p为nil，应当返回pre
	return pre
}

func printList(p *Node) {
	for p != nil {
		node := p
		fmt.Printf("%+v\n", node)
		p = node.Next
	}
}
