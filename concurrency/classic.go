package concurrency

import (
	"fmt"
	"sync"
	"time"
)

/*
func Run() {
	go printText("Hello")
	time.Sleep(1 * time.Second)
}

func printText(text string) {
	fmt.Println(text)
}
*/

//WaitGroup

/*func Run() {
	wg := sync.WaitGroup{}
	letters := []string{"a", "b", "c", "d", "e", "f", "g"}
	wg.Add(len(letters))
	for _, letter := range letters {
		go printText(letter, &wg)
	}
	wg.Wait()
	fmt.Println("Done")
}

func printText(text string, wg *sync.WaitGroup) {
	fmt.Println(text)
	wg.Done()
}*/

// Mutex

/*var counter = 0

func Run() {
	n := 1000000
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go increment(&wg, &mutex)
	}
	wg.Wait()
	fmt.Println(counter)
}

func increment(wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	counter += 1
}*/

//  RWMutex

/*type safeSlice[T any] struct {
	mutex sync.RWMutex
	data  []T
}

func (ss *safeSlice[T]) add(value T) {
	ss.mutex.Lock()
	ss.data = append(ss.data, value)
	ss.mutex.Unlock()
}

func (ss *safeSlice[T]) get(index int) (T, bool) {
	ss.mutex.RLock()
	defer ss.mutex.RUnlock()
	if index < 0 || index >= len(ss.data) {
		var empty T
		return empty, false
	}
	return ss.data[index], true
}

func (ss *safeSlice[T]) size() int {
	ss.mutex.RLock()
	defer ss.mutex.RUnlock()
	return len(ss.data)
}

func Run() {
	ss := safeSlice[int]{}

	go func() {
		ss.add(1)
		ss.add(2)
	}()

	go func() {
		fmt.Printf("Size: %d\n", ss.size())
	}()

	time.Sleep(1 * time.Second)
}*/

// Mutex + Signals

var (
	money                 = 100
	mutex                 = sync.Mutex{}
	moneyIsGraterThanZero = sync.NewCond(&mutex)
	spendValue            = 10
)

func spend() {
	for i := 1; i < 500; i++ {
		mutex.Lock()
		for money-spendValue < 10 {
			moneyIsGraterThanZero.Wait()
		}
		money -= spendValue
		fmt.Println("Spend: ", money)
		mutex.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Spend: Done")
}

func work() {
	for i := 1; i < 500; i++ {
		mutex.Lock()
		money += 5
		fmt.Println("New income, current value:", money)
		moneyIsGraterThanZero.Broadcast() // all threads
		//moneyIsGraterThanZero.Signal() // one random thread
		mutex.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Work: Done")
}

func Run() {
	go work()
	go spend()

	time.Sleep(10 * time.Second)
	fmt.Println("Current value:", money)
}
