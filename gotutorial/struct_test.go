package slice

import (
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {
	type rectangle struct {
		length  float64
		breadth float64
		color   string
	}

	fmt.Println(rectangle{10.5, 25.10, "red"})
}

func TestStructAssign(t *testing.T) {
	type rectangle struct {
		length  int
		breadth int
		color   string

		geometry struct {
			area      int
			perimeter int
		}
	}

	var rect rectangle // dot notation
	rect.length = 10
	rect.breadth = 20
	rect.color = "Green"

	rect.geometry.area = rect.length * rect.breadth
	rect.geometry.perimeter = 2 * (rect.length + rect.breadth)

	fmt.Println(rect)
	fmt.Println("Area:\t", rect.geometry.area)
	fmt.Println("Perimeter:", rect.geometry.perimeter)

}
