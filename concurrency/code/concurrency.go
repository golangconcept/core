package main

import (
	"fmt"
	"time"
)

// func pinger(c chan string) {
// 	for i := 0; ; i++ {
// 		c <- "ping"
// 	}
// }
// func printer(c chan string) {
// 	for {
// 		msg := <-c
// 		fmt.Println(msg)
// 		time.Sleep(time.Second * 1)
// 	}
// }
// func ponger(c chan string) {
// 	for i := 0; ; i++ {
// 		c <- "pong"
// 	}
// }

//	func f(n int) {
//		for i := 0; i < 10; i++ {
//			fmt.Println(n, ":", i)
//			amt := time.Duration(rand.Intn(250))
//			time.Sleep(time.Millisecond * amt)
//		}
//		// fmt.Println("++++++++++++++++++++++++++++++++")
//	}

func infiniteCount(thing string) {
	for i := 1; ; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Second * 1)
	}
}

func countwithChannel(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Microsecond * 500)
	}
	close(c)
}
func main() {
	// for i := 0; i < 10; i++ {
	// 	go f(0)
	// }
	// var c chan string = make(chan string)

	// go pinger(c)
	// go ponger(c)
	// go printer(c)

	// var wg sync.WaitGroup
	// wg.Add(1)

	// go func() {
	// 	infiniteCount("dog")
	// 	wg.Done()
	// }()

	// wg.Wait()
	// c := make(chan string)
	// go countwithChannel("dog", c)

	// for msg := range c {
	// 	fmt.Println(msg)
	// }

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Microsecond * 500)
			c1 <- "every 500ms"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			c2 <- "every 2 seconds"
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
	// go infiniteCount("cat")
	// var input string
	// fmt.Scanln(&input)
}
