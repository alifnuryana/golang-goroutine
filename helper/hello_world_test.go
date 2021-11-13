package helper

import "testing"

func TestGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go HelloWorld(i)
	}
}
