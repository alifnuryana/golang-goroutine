package helper

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandomNumber(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	//jika menggunakan 2 channel
	//var ra, rb = make(chan int32), make(chan int32)
	//defer close(ra)
	//defer close(rb)

	//menggunakan 1 channel
	var sisi = make(chan int32, 2)
	defer close(sisi)

	go longTimeRequest(sisi)
	go longTimeRequest(sisi)

	fmt.Println(sumSquare(<-sisi, <-sisi))
}
