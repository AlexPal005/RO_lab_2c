package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count_first  int
	count_second int
}

func main() {
	start()
	time.Sleep(2 * time.Second)
}
func start() {
	ch := make(chan string)
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go smoker_one(ch1)
	go smoker_two(ch2)
	go smoker_three(ch3)
	go broker(ch)
	for i := 1; i < 11; i++ {
		first := <-ch
		second := <-ch
		ch1 <- first
		ch1 <- second
		ch2 <- first
		ch2 <- second
		ch3 <- first
		ch3 <- second
	}
}
func smoker_one(ch chan string) {
	//element := "tobacco"
	for i := 1; i < 11; i++ {
		first := <-ch
		second := <-ch

		if first == "paper" && second == "matches" {
			fmt.Println("The first smokes...")
			time.Sleep(2 * time.Second)
		}
	}
}
func smoker_two(ch chan string) {
	//element := "paper"
	for i := 1; i < 11; i++ {
		first := <-ch
		second := <-ch

		if first == "tobacco" && second == "matches" {
			fmt.Println("The second smokes...")
			time.Sleep(2 * time.Second)
		}
	}
}
func smoker_three(ch chan string) {
	//element := "matches"
	for i := 1; i < 11; i++ {
		first := <-ch
		second := <-ch

		if first == "tobacco" && second == "paper" {
			fmt.Println("The third smokes...")
			time.Sleep(2 * time.Second)
		}
	}
}
func broker(ch1 chan string) {
	arr := []string{"tobacco", "paper", "matches"}
	number := new(Counter)
	number.count_first = 0
	number.count_second = 1
	for i := 1; i < 11; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println("Broker : ", arr[number.count_first], " ", arr[number.count_second])
		ch1 <- arr[number.count_first]
		ch1 <- arr[number.count_second]
		if number.count_second == 1 && number.count_first == 0 {
			number.count_first = 1
			number.count_second = 2
		} else if number.count_second == 2 && number.count_first == 1 {
			number.count_first = 0
			number.count_second = 2
		} else if number.count_second == 2 && number.count_first == 0 {
			number.count_first = 0
			number.count_second = 1
		}
	}
}
