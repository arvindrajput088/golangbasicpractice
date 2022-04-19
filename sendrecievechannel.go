/*making a channel
c := make(chan int)
● putting values on a channel
c <- 42
● taking values off of a channel
<-c
● buffered channels
c := make(chan int, 4)*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
