package gotutorial

import (
	"fmt"
	"testing"
)

var done = make(chan bool)
var doneL = make(chan bool, 3)

func goroutineWrite() {
	fmt.Println("goroutineWrite start")
	fmt.Println("goroutineWrite over")
	done <- true
}

func TestUNBlockChanWrite(t *testing.T) {
	go goroutineWrite()
	<-done
	fmt.Println("test case over!")
	done <- true //永远阻塞在这里
	fmt.Println("block forever")
}

func TestUNBlockChanRead(t *testing.T) {
	<-done
	fmt.Println("block forever")
}

func blockGoroutineWrite() {
	fmt.Println("blockGoroutineWrite start")
	fmt.Println("blockGoroutineWrite over")
	doneL <- true
}

func TestBlockChanWrite(t *testing.T) {
	go blockGoroutineWrite()
	<-doneL
	fmt.Println("test case over!")
	doneL <- true //不会阻塞在这里
	fmt.Println("nonblock")
}
