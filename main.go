package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/wiput1999/go-tutorial/slice"

	"github.com/wiput1999/go-tutorial/frog"
)

func main() {

	// Frog Jump Problem
	println(frog.Jump(10, 85, 30))
	println(frog.Jump(15, 200, 45))
	println(frog.Jump(5, 1000, 55))

	// Error Handling
	f, err := os.Open("go.mod")
	if err != nil {
		log.Fatal(err)
	}

	if b, err := ioutil.ReadAll(f); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(b))
	}

	// Array
	s := []int{1, 2, 3, 4, 5}

	if s == nil {
		fmt.Println("It's nil")
	}

	s = append(s, 6, 7)

	fmt.Println(len(s), cap(s), s)

	// for _, v := range s {
	// 	fmt.Println(v)
	// }

	fmt.Println(s[2:])
	fmt.Println(append(s, s[2:5]...))

	// Slice Problem
	fmt.Println(slice.Slice([]string{"a", "b", "c"}, []string{"1", "2", "3"}))

	// Map
	// m := make(map[string]string)
	m := map[string]string{
		"one":   "หนึ่ง",
		"two":   "สอง",
		"three": "สาม",
	}

	if m == nil {
		fmt.Println("It's nil")
	}

	// m["one"] = "หนึ่ง"

	for k, v := range m {
		fmt.Println(k, v)
	}

	if v, ok := m["four"]; ok {
		fmt.Printf("%q\n", v)
	} else {
		fmt.Println("Empty key")
	}

	varidic()

	// First-Class func
	var fn = func(n int) int {
		return n * n
	}

	printPower(4, fn)

	inc, cFn := closureCounter()
	fmt.Println(cFn())
	fmt.Println(cFn())
	inc()
	fmt.Println(cFn())

	rect := rectangle{3, 4}
	fmt.Println(rect.area())

}

type String string

func (s String) toString() string {
	return string(s)
}

type Int int

func (i Int) toString() string {
	return strconv.Itoa(int(i))
}

type xxxFunc func(int) func() int

func (x xxxFunc) Int() int {
	fn := x(12)
	return fn()
}

type rectangle struct {
	width  int
	length int
}

func (rec rectangle) area() int {
	return rec.width * rec.length
}

type counterFunc func() int
type increaseFunc func()

func closureCounter() (increaseFunc, counterFunc) {
	var i int
	return func() { i++ },
		func() int {
			return i
		}
}

type powerFunc func(int) int

func printPower(n int, fn powerFunc) {
	fmt.Println(fn(n))
}

func varidic(a ...int) {
	for _, v := range a {
		fmt.Println(v)
	}
}
