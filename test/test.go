package main

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"runtime"
	"strings"
)

func add(x ...int) (result int) {
	for _, num := range x {
		result += num
	}
	return
}

func looper() {
	for i := 1; i <= 25; i++ {
		fmt.Println(i)
	}
}

func cond() {
	r := rand.Intn(100)
	if r < 50 {
		fmt.Printf("Less than 50: %d\n", r)
	} else if r == 50 {
		fmt.Println("Equal to 50")
	} else {
		fmt.Printf("Greater than 50: %d\n", r)
	}
}

func condSwitch() {
	r := rand.Intn(100)
	switch {
	case r < 50:
		fmt.Println("Less than 50:", r)
	case r == 50:
		fmt.Println("Equal to 50:", r)
	case r > 50:
		fmt.Println("More than 50:", r)
	}
}

func sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		fmt.Println(z)
		z -= (z*z - x) / (2 * z)
		if (z*z - x) < 0.0001 {
			fmt.Println(z)
			return z
		}
	}
	return z
}

func getOS() {
	fmt.Println("OS is:", runtime.GOOS)
}

func deferring() {
	fmt.Println("Counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}

type Point struct {
	X, Y int
}

func prntPnt(x, y int) {
	v := Point{x, y}
	p := &v
	fmt.Println("Pointer to point:", p)
	fmt.Println("Point X:", p.X)
	fmt.Println("Point Y:", p.Y)
}

func Pic(dx, dy int) (a [][]uint8) {
	a = make([][]uint8, dy)
	for i := range a {
		a[i] = make([]uint8, dx)
		for j := range a[i] {
			a[i][j] = uint8(i * j)
		}
	}
	return
}

func wordCount(s string) (results map[string]int) {
	results = make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		results[word]++
	}
	return
}

func clozures() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fib() func() int {
	prev := 0
	current := 1
	return func() int {
		temp := prev
		prev = current
		current = temp + current
		return temp
	}
}

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	if err != nil {
		return n, err
	}
	for i := range b {
		b[i] = rot13(b[i])
	}
	return n, nil
}

func rot13(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return (b-'A'+13)%26 + 'A'
	} else if b >= 'a' && b <= 'z' {
		return (b-'a'+13)%26 + 'a'
	}
	return b
}

func addMsg(msg string, c chan string) {
	c <- msg
}

func sendOnly(c chan<- string, msg string) {
	c <- msg
}

func receiveAndSend(cS <-chan string, cR chan<- string) {
	msg := <-cS
	cR <- msg
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}

func main() {
}
