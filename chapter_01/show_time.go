package main

import (
	"fmt"
	"time"
)

func displayTime() {
	fmt.Println(time.Now())
}

// func countdown(stop time.Duration) {
// 	time.AfterFunc(stop, func() {
// 		fmt.Println(time.Now())
// 	})
// }
