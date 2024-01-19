package gotutorial

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestFlowControl(t *testing.T) {
	// first ： if condition
	var s = "Japan"
	x := true
	if x {
		fmt.Println(s)
	}

	//second：
	x2 := 100
	if x2 == 100 {
		fmt.Println("Japan")
	} else {
		fmt.Println("Canada")
	}

	x3 := rand.Intn(150)
	//third
	if x3 <= 50 {
		fmt.Println("Germany")
	} else if 50 < x3 && x3 <= 100 {
		fmt.Println("Japan")
	} else {
		fmt.Println("Canada")
	}

	//third
	if x := 100; x == 100 {
		fmt.Println("Germany")
	}
}

func TestForLoop(t *testing.T) {
	k := 1
	for ; k <= 10; k++ {
		fmt.Println(k)
	}

	k = 1
	for k <= 10 {
		fmt.Println(k)
		k++
	}

	for k := 1; ; k++ {
		fmt.Println(k)
		if k == 10 {
			break
		}
	}

	// Example 1
	strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
	for index, element := range strDict {
		fmt.Println("Index :", index, " Element :", element)
	}

	// Example 2
	for key := range strDict {
		fmt.Println(key)
	}

	// Example 3
	for _, value := range strDict {
		fmt.Println(value)
	}

	// Example 4
	i := 5
	for {
		fmt.Println("Hello")
		if i == 10 {
			break
		}
		i++
	}
}

func TestRangeLoopOverString(t *testing.T) {

	for range "Hello" {
		fmt.Println("Hello")
	}

	for v := range "Hello" {
		fmt.Println(v, "Hello")
	}

	for v, c := range "Hello" {
		fmt.Println(v, c, string(c), "Hello")
	}
}
