package main

import "fmt"

func main() {
	// variables
	a := 2
	b := &a
	*b = 3 // a = 3;
	c := &a

	d := new(int)
	*d = 12
	*c = *d // c = 12 -> a = 12
	*d = 13

	c = d
	*c = 14 // c = 14 -> d = 14, a = 12

	// array (размер массива, это часть типа данных)
	var a1 [3]int
	fmt.Printf("default array: %v\n", a1)

	a3 := [4]int{1, 2, 3}
	fmt.Printf("initialized array: %v\n", a3)

	//slice
	var buf0 []int
	buf1 := []int{}
	buf2 := []int{42}

	buf3 := []string{"est", "ax", "xer"}

	buf4 := make([]string, 2)
	buf4[0] = "ty"
	buf4[1] = "ne"

	buf5 := buf3[0:]

	fmt.Println(buf0, buf1, buf2)
	fmt.Println(buf4)
	fmt.Println(buf5)

	buf5 = append(buf5, buf4...)

	fmt.Println(buf5)

	var bufLen, bufCap = len(buf4), cap(buf5)

	fmt.Printf("buf buf4: %d, cap buf5: %d", bufLen, bufCap)

	// slice, extended operations
	buf6 := []int{1, 2, 3, 4, 5}
	fmt.Println(buf6)

	// get slice,
	sl1 := buf6[1:4] // [2,3,4]
	sl2 := buf6[:2]  // [1,2]
	sl3 := buf6[2:]  // [3,4,5]

	fmt.Println(sl1, sl2, sl3)

	sl2[0] = 7
	fmt.Println(buf6)

	// map
	var user = map[string]string{
		"name":     "Karl",
		"lastName": "Ivanov",
	}

	// with empty of cap
	profile := make(map[string]string, 10)

	mapLength := len(user)

	fmt.Printf("%d %+v\n", mapLength, profile)

	delete(user, "lastName")

	fmt.Printf("%v\n", user)
}
