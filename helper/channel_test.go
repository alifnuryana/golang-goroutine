package helper

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	CreateChannel("Alif")
}

func TestGiveMeResponese(t *testing.T) {
	var channel = make(chan string)
	defer close(channel)

	go GiveMeResponese(channel)

	var data = <-channel
	fmt.Println(data)
}

func TestChannelInOut(t *testing.T) {
	var channel = make(chan string)
	defer close(channel)

	go OnlyIn(channel, "Alif Nuryana")
	go OnlyOut(channel)

	// mengirimkan data channel kedalam variabel
	var data = <-channel

	// mengecek apakah data sudah sampai dan sesuai
	require.Equal(t, "Alif Nuryana", data)

	time.Sleep(time.Second * 2)
}

func TestChannelRange(t *testing.T) {
	var channel = make(chan string)

	go ChannelRange(100, channel)

	for c := range channel {
		fmt.Println(c)
	}

}
