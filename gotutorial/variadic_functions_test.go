package gotutorial_test

import (
	"fmt"
	"testing"
)

func variadicExample(s ...string) {
	if len(s) > 4 {
		fmt.Println(s[0])
		fmt.Println(s[3])
	} else {
		for ss, vv := range s {
			fmt.Println(ss, vv)
		}
	}

}

func TestVariablesFunction(t *testing.T) {
	variadicExample("hell0", "world", "ddd")
}
