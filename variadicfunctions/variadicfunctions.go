package main

import "fmt"

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, " found at index ", i, " in ", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, " not found in ", nums)
	}
	fmt.Printf("\n")
}

func variadicParams() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)
}

func sliceAsAVariadicParams() {
	nums := []int{89, 90, 85}
	find(89, nums...)
}

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func gotcha() {
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}

func main() {
	fmt.Println("1. variadic parameters")
	variadicParams()
	fmt.Println("2. Passing a slice to a variadic function")
	sliceAsAVariadicParams()
	fmt.Println("3. variadic parameter as paramter")
	gotcha()
}
