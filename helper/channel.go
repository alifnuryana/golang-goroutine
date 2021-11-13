package helper

import (
	"fmt"
	"strconv"
	"time"
)

func CreateChannel(name string) {
	// membuat channel dan menutup diakhir function
	var channel = make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(time.Second * 2)
		channel <- name
		fmt.Println("selesai mengirim data ke channel")
	}()

	// mengirimkan data dari channel ke variabel data
	var data = <-channel
	fmt.Println("Dari Channel " + data)
}

func GiveMeResponese(channel chan string) {
	channel <- "Alif Nuryana"
}

// function hanya untuk bisa memasukan data kedalam channel
func OnlyIn(channel chan<- string, name string) {
	channel <- name
}

// function hanya untuk bisa mengambil data dari channel
func OnlyOut(channel <-chan string) {
	var data = <-channel
	fmt.Println(data)
}

func ChannelRange(n int, channel chan<- string) {
	defer close(channel)
	for i := 0; i < n; i++ {
		channel <- "Perulangan ke " + strconv.Itoa(i)
	}
}
