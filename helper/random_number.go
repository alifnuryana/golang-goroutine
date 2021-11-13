package helper

import (
	"math/rand"
)

// func channel in for generate random number
func longTimeRequest(channel chan<- int32) {
	channel <- rand.Int31n(100)
}

// func sum square
func sumSquare(a, b int32) int32 {
	return a*a + b*b
}
