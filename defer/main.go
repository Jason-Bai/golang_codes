package main

import (
	"fmt"
	"sync"
)

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

func deferExample() {
	nums := []int{78, 109, 2, 563, 300}
	largest(nums)
}

type person struct {
	firstName, lastName string
}

func (p person) fullName() {
	fmt.Printf(" %s %s\n", p.firstName, p.lastName)
}

func deferredMethods() {
	p := person{
		firstName: "John",
		lastName:  "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome")
}

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}

func argumentsEvaluation() {
	a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)
}

func stackOfDefers() {
	name := "Naveen"
	fmt.Printf("Orignal String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
}

type rect struct {
	length, width int
}

func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		// wg.Done()
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		// wg.Done()
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
	// wg.Done()
}

func practicalUseOfDefer() {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

func main() {
	fmt.Println("1. defer example")
	deferExample()
	fmt.Println("2. defered method")
	deferredMethods()
	fmt.Println("3. arguments evaluation")
	argumentsEvaluation()
	fmt.Println("4. stack of defers")
	stackOfDefers()
	fmt.Println()
	fmt.Println("5. Practical use of defer")
	practicalUseOfDefer()
}
