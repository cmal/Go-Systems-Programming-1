package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	Value int
	Next  *Node
}

func addNode(t *Node, v int) int {
	if root == nil { // test whether global root is nil
		t = &Node{v, nil}
		root = t // give root a initial root address, NOTE this will not be changed forever
		return 0
	}

	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return -1
	}

	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}

	return addNode(t.Next, v)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

var root = new(Node)

func oldMain() {
	fmt.Println(root)  // global root is a pointer
	root = nil // global root is now a 0 pointer
	traverse(root) // root is nil
	addNode(root, 1)
	addNode(root, 1)
	traverse(root)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 0)
	addNode(root, 0)
	traverse(root)
	addNode(root, 100)
	traverse(root)
}

func testFunc(t *Node) {
	fmt.Println("t: ", &t)
	fmt.Println("t: ", t)
	t = &Node{3, nil} // 结论： 这个操作改变了指针 t 所指向的地址，而对指针 t 原来指向的地址没有做任何操作
	fmt.Println("t: ", &t)
	fmt.Println("t: ", t)
}

func testFunc2(t Node) {
	fmt.Println("t: ", t)
	t = Node{3, nil}
	fmt.Println("t: ", t)
}

func testFunc3(t Node) {
	fmt.Println("t: ", t)
	t.Value = 4
	fmt.Println("t: ", t)
}


func newMain() {
	a := new(Node)
	fmt.Println("a: ", &a)
	testFunc(a) // 结论： 虽然这里传递了指针（引用），但函数内的操作并没有改变指针所指向的地址内的结构体的内容
	fmt.Println("a: ", &a)

	// b := new(Node) // prevent the error `use of untyped nil`
	// b = nil
	// or use:
	var b *Node
	
	fmt.Println("b: ", &b)
	testFunc(b)
	fmt.Println("b: ", &b)
}


func newMain2() {
	a := new(Node)
	b := Node{3, nil}
	fmt.Println(reflect.TypeOf(a)) // 指针
	fmt.Println(reflect.TypeOf(b)) // 结构体
	fmt.Println(reflect.TypeOf(&b)) // 指针
	fmt.Println(reflect.TypeOf(*a)) // 结构体
	fmt.Println(reflect.TypeOf(*&b)) // 结构体
	fmt.Println(reflect.TypeOf(*&*&*&b)) // 结构体

	// therefore, we can do this:
	testFunc(a)
	testFunc(&b)
	testFunc2(*a)
	testFunc2(b)
	testFunc3(*a)
	testFunc3(b)

	// 以上操作都不会影响 a 和 b 所代表的结构体内的值

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	// but we can not do this:
	// testFunc(&a) // a 是指针， &a 没有意义
	// testFunc(b) // b 是结构体，testFunc要求参数是指针
	
	
}

func main() {
	oldMain()
	// newMain()
	// newMain2()
}
