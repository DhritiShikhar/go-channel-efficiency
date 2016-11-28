package main

import (
	"fmt"
	"time"
)

func timetrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", name, elapsed)
}

func main() {
	defer timetrack(time.Now(), "Rules")
	finished := make(chan bool)

	go func() {
		ok := r1()
		if ok {
			finished <- true
			fmt.Println("Done!")
		}
	}()
	go func() {
		ok := r2()
		if ok {
			finished <- true
			fmt.Println("Done!")
		}
	}()
	go func() {
		ok := r3()
		if ok {
			finished <- true
			fmt.Println("Done!")
		}
	}()
	go func() {
		ok := r4()
		if ok {
			finished <- true
			fmt.Println("Done!")
		}
	}()
	go func() {
		ok := r5()
		if ok {
			finished <- true
			fmt.Println("Done!")
		}
	}()
	go func() {
		ok := r6()
		if ok {
			finished <- true
			fmt.Println("Done!")
		}
	}()

	for {
		select {
		case isFinished := <-finished:
			fmt.Println("Done: ", isFinished)
			return
		}
	}
}

func r1() bool {
	time.Sleep(1000)
	return false
}
func r2() bool {
	time.Sleep(1000)
	return false
}
func r3() bool {
	time.Sleep(1000)
	return false
}
func r4() bool {
	time.Sleep(1000)
	return false
}
func r5() bool {
	time.Sleep(1000)
	return false
}
func r6() bool {
	time.Sleep(1000)
	return false
}
