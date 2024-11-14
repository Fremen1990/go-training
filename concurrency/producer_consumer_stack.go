package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type Stack struct {
	data     []int
	mutex    sync.Mutex
	notEmpty *sync.Cond
	notFull  *sync.Cond
	maxSize  int
}

func NewStack(maxSize int) *Stack {
	s := &Stack{
		data:    []int{},
		maxSize: maxSize,
	}
	s.notEmpty = sync.NewCond(&s.mutex)
	s.notFull = sync.NewCond(&s.mutex)
	return s
}

func (s *Stack) Push(value int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for len(s.data) >= s.maxSize {
		fmt.Println("Producer waiting: stack is full")
		s.notFull.Wait()
	}

	s.data = append(s.data, value)
	fmt.Println("Produced:", value)
	s.notEmpty.Signal()
}

func (s *Stack) Pop() (int, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for len(s.data) == 0 {
		fmt.Println("Consumer waiting: stack is empty")
		s.notEmpty.Wait()
	}

	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	fmt.Println("Consumed:", value)
	s.notFull.Signal()
	return value, true
}

func Producer(s *Stack, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 20; i++ {
		time.Sleep(time.Millisecond * 100)
		s.Push(i)
	}
}

func Consumer(s *Stack, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 20; i++ {
		time.Sleep(time.Millisecond * 150)
		s.Pop()
	}
}

func RunProducerConsumerStack() {
	stack := NewStack(3)
	var wg sync.WaitGroup

	wg.Add(1)
	go Producer(stack, &wg)
	wg.Add(1)
	go Consumer(stack, &wg)

	wg.Wait()
	fmt.Println("All tasks complete.")
}
