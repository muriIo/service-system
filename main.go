package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value string
	next  *Node
}

type Queue struct {
	front *Node
	rear  *Node
}

func (queue *Queue) isEmpty() bool {
	return queue.front == nil
}

func (queue *Queue) enqueue(value string) {
	newNode := &Node{value: value}

	if queue.isEmpty() == true {
		queue.front = newNode
		queue.rear = newNode
	} else {
		queue.rear.next = newNode
		queue.rear = newNode
	}
}

func (queue *Queue) dequeue() (*Node, error) {
	if queue.isEmpty() == true {
		return nil, errors.New("Queue is already empty")
	}

	dequeuedValue := queue.front
	queue.front = queue.front.next

	if queue.front == nil {
		queue.rear = nil
	}

	return dequeuedValue, nil
}

func (queue *Queue) peek() (string, error) {
	if queue.isEmpty() == true {
		return "", errors.New("Queue is empty")
	}

	return queue.front.value, nil
}

func main() {
	queue := &Queue{}

	fmt.Println(queue.isEmpty())
	fmt.Println(queue.peek())
	fmt.Println(queue.dequeue())

	queue.enqueue("josé")
	queue.enqueue("maria")
	queue.enqueue("cláudio")
	queue.enqueue("josefina")

	fmt.Println(queue.peek())
	fmt.Println(queue.dequeue())
	fmt.Println(queue.peek())
	fmt.Println(queue.dequeue())
	fmt.Println(queue.dequeue())
	fmt.Println(queue.dequeue())
	fmt.Println(queue.dequeue())
}
