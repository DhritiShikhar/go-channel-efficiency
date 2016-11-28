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
	ok := r1()
	if ok {
		fmt.Println("Done!")
	}

	ok = r2()
	if ok {
		fmt.Println("Done!")
	}

	ok = r3()
	if ok {
		fmt.Println("Done!")
	}

	ok = r4()
	if ok {
		fmt.Println("Done!")
	}

	ok = r5()
	if ok {
		fmt.Println("Done!")
	}

	ok = r6()
	if ok {
		fmt.Println("Done!")
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
