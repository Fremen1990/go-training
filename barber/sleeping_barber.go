package barber

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BarberShop struct {
	customerChan chan int
	doneChan     chan bool
	quitChan     chan bool
	waitingSeats chan int
	barberSleep  bool
	open         bool
	wg           *sync.WaitGroup
}

// NewBarberShop creates a new BarberShop
func NewBarberShop(numSeats int, wg *sync.WaitGroup) *BarberShop {
	return &BarberShop{
		customerChan: make(chan int),
		doneChan:     make(chan bool),
		quitChan:     make(chan bool),
		waitingSeats: make(chan int, numSeats),
		barberSleep:  true,
		open:         true,
		wg:           wg,
	}
}

// barber function that cuts hair or sleeps
func (shop *BarberShop) barber() {
	for {
		select {
		case customerID := <-shop.customerChan:
			// Wake up if asleep and start a haircut
			if shop.barberSleep {
				fmt.Println("Barber wakes up.")
				shop.barberSleep = false
			}
			fmt.Printf("Barber starts cutting hair for Customer %d.\n", customerID)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)+100)) // Simulate haircut time
			fmt.Printf("Barber finished cutting hair for Customer %d.\n", customerID)
			shop.doneChan <- true // Notify customer that the haircut is done

		case <-shop.quitChan:
			// Barber goes home when the shop is closed and waiting room is empty
			if len(shop.waitingSeats) == 0 {
				fmt.Println("Barber goes home as shop is closed and no more customers.")
				shop.wg.Done() // Notify main routine that barber is done
				return
			}
		}
	}
}

// customer function to represent arriving customers
func (shop *BarberShop) customer(id int) {
	if !shop.open {
		fmt.Printf("Customer %d leaves as the shop is closed.\n", id)
		return
	}
	fmt.Printf("Customer %d arrives at the shop.\n", id)

	select {
	case shop.waitingSeats <- id:
		fmt.Printf("Customer %d waits in the waiting room.\n", id)
		shop.customerChan <- id // Notify barber of the customer’s arrival

		<-shop.doneChan // Wait for haircut to be done
		fmt.Printf("Customer %d leaves after haircut.\n", id)

	default:
		// No seats available
		fmt.Printf("Customer %d leaves as no seats are available.\n", id)
	}
}

// openShop opens the shop for a certain duration and then stops accepting new clients
func (shop *BarberShop) openShop(duration time.Duration) {
	fmt.Println("Barber Shop is now open.")
	time.AfterFunc(duration, func() {
		shop.open = false
		close(shop.quitChan) // Signal the barber to start closing routine
	})
}

func SleepingBarber() {

	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	// Initialize the barber shop with 3 waiting seats
	shop := NewBarberShop(3, &wg)
	wg.Add(5) // Barber routine

	for i := 1; i <= 5; i++ {
		go shop.barber()
		fmt.Printf("Barber %d is ready.\n", i)
	}

	// Start barber routine
	go shop.barber()

	// Open the shop for 5 seconds
	shop.openShop(5 * time.Second)

	// Simulate customer arrivals
	customerID := 1
	for shop.open {
		go shop.customer(customerID)
		customerID++
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)+100)) // Random interval for next customer
	}

	wg.Wait()
	fmt.Println("Barber Shop is closed.")
}
