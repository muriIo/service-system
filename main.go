package main

import (
	"errors"
	"fmt"
)

type Priority int

const (
	Normal Priority = iota
	Urgency
	Emergency
)

type Node struct {
	value    string
	next     *Node
	priority Priority
}

type Queue struct {
	front  *Node
	rear   *Node
	length int
}

func (queue *Queue) isEmpty() bool {
	return queue.front == nil
}

func (queue *Queue) enqueue(value string, priority Priority) {
	newNode := &Node{value: value, priority: priority}

	if queue.isEmpty() {
		queue.front = newNode
		queue.rear = newNode
		queue.length = 1

		return
	}

	queue.length++

	if newNode.priority == Normal {
		queue.rear.next = newNode
		queue.rear = newNode

		return
	}

	currentNode := queue.front

	for currentNode.next != nil && currentNode.next.priority >= priority {
		currentNode = currentNode.next
	}

	newNode.next = currentNode.next
	currentNode.next = newNode

	if newNode.next == nil {
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

	queue.length--

	return dequeuedValue, nil
}

func (queue *Queue) peek() (*Node, error) {
	if queue.isEmpty() == true {
		return nil, errors.New("Queue is empty")
	}

	return queue.front, nil
}

func main() {
	queue := &Queue{}

	queue.enqueue("josefina", Emergency)
	queue.enqueue("josé", Normal)
	queue.enqueue("marcos", Emergency)
	queue.enqueue("maria", Urgency)
	queue.dequeue()
	queue.enqueue("cláudio", Urgency)
	queue.dequeue()

	front, error := queue.peek()

	if error != nil {
		return
	}

	fmt.Println(front)
	fmt.Println(queue.length)
}
