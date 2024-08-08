package main

import "fmt"

func main() {

	// Initialize Queues
	q1 := Queue{"Q1", []int{}, 0}
	q2 := Queue{"Q2", []int{}, 0}

	// Initialize Consumers reading from Q1
	c1q1 := Consumer{"Consumer-1-Q1", q1.topIndex, 1}
	c2q1 := Consumer{"Consumer-2-Q1", q1.topIndex, 1}
	c3q1 := Consumer{"Consumer-3-Q1", q1.topIndex, 1}

	// Initialize Consumers reading from Q2
	c1q2 := Consumer{"Consumer-1-Q2", q2.topIndex, 1}
	c2q2 := Consumer{"Consumer-2-Q2", q2.topIndex, 1}

	// Initialize Producer
	p := Producer{"Producer-1"}

	p.Produce(1, &q1)
	p.Produce(2, &q1)
	p.Produce(3, &q1)
	consumeAndPrintFromQueue(&q1, &c1q1)

	p.Produce(4, &q1)
	consumeAndPrintFromQueue(&q1, &c2q1)
	consumeAndPrintFromQueue(&q1, &c3q1)

	p.Produce(5, &q1)
	p.Produce(6, &q1)
	consumeAndPrintFromQueue(&q1, &c2q1)

	p.Produce(7, &q1)
	consumeAndPrintFromQueue(&q1, &c2q1)
	consumeAndPrintFromQueue(&q1, &c3q1)
	consumeAndPrintFromQueue(&q1, &c1q1)

	p.Produce(101, &q2)
	p.Produce(102, &q2)
	consumeAndPrintFromQueue(&q2, &c1q2)

	p.Produce(103, &q2)
	consumeAndPrintFromQueue(&q2, &c1q2)
	consumeAndPrintFromQueue(&q2, &c2q2)
}

// Read from Queue, and prints the slice of consumed values
func consumeAndPrintFromQueue(q *Queue, consumer *Consumer) {
	fmt.Println(fmt.Sprintf("%v consumes from Queue:%v | Call-%v => ", consumer.name, q.name, consumer.calledTimes), consumer.Consume(q))
}

// Define the struct for Consumer
type Consumer struct {
	name        string
	topIndex    int
	calledTimes int
}

// Define 'Consume' method
func (c *Consumer) Consume(q *Queue) []int {
	readval := []int{}
	size := len(q.items)
	indx := c.topIndex
	for indx <= size-1 {
		readval = append(readval, q.Pop(indx))
		indx++
	}
	c.topIndex = indx
	c.calledTimes = c.calledTimes + 1
	return readval
}

// Define the struct for Producer
type Producer struct {
	name string
}

// Define 'Produce' method
func (p *Producer) Produce(val int, q *Queue) {
	fmt.Println(fmt.Sprintf("%v ===> Queue:%v | %v", p.name, q.name, val))
	q.Push(val)
}

// Define the struct for Queue
type Queue struct {
	name     string
	items    []int
	topIndex int
}

// Push to Queue
func (q *Queue) Push(val int) {
	q.items = append(q.items, val)
}

// Pull from Queue
func (q *Queue) Pop(indx int) int {
	popVal := q.items[indx]
	return popVal
}
