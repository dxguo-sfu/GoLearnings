package basics

// package main

// import(
// 	"fmt"
// )

// func main() {
// 	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

// 	ch := make(chan int)
// 	var a int

// 	for _, v := range arr {
// 		go func(ch chan int, n int) {
// 			ch <- n
// 		}(ch, v)
// 	}
// 	// for i := 0; i < len(arr)+1; i++ {
// 	// for i := 0; i < len(arr)-4; i++ {
// 	for i := 0; i < len(arr); i++ {
// 		a = <-ch
// 		fmt.Println(a)
// 	}
// }

/*
		Reduced for concurrency
	First for loop puts the array in the channel concurrently
	All goroutines are fired
	Execution is allowed to continue until <- is needed

	Second loops prints the order that items have been added
	<- awaits on ch until there is a value to be drawn
	Deadlock occurs if all goroutines have finished but we still await with <-
	(Use the commented for loop line)
	Unawaited routines will not cause problems but may lead to unexpected errors
*/
