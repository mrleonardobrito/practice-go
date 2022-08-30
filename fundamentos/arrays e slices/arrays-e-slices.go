package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 = [5]int{1, 2, 3, 4, 5}
	var slice = []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(array1)
	fmt.Println(slice[6])

	var slice2 []int = array1[2:4]

	fmt.Println(reflect.TypeOf(array1))
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(slice2)

	// ARRAYS INTERNOS

	slice3 := make([]float32, 10, 15)
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	slice3 = append(slice3, 6)
	slice3 = append(slice3, 6)
	slice3 = append(slice3, 6)
	slice3 = append(slice3, 6)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
}
