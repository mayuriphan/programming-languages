package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// =========================
	// VARIABLES
	// =========================

	var age int = 30
	var salary float64 = 50000.50
	var name string = "John"
	var isActive bool = true

	fmt.Println("Age:", age)
	fmt.Println("Salary:", salary)
	fmt.Println("Name:", name)
	fmt.Println("Active:", isActive)

	// Short declaration
	city := "Pune"
	fmt.Println("City:", city)

	// =========================
	// DATA TYPES
	// =========================

	var i int = 100
	var f float64 = 10.5
	var s string = "Hello"
	var b bool = false

	fmt.Println(i, f, s, b)

	// =========================
	// CONSTANTS
	// =========================

	const PI = 3.14159
	fmt.Println("PI:", PI)

	// =========================
	// ASSIGNMENT
	// =========================

	x := 10
	y := 20

	fmt.Println("Before:", x, y)

	x = y

	fmt.Println("After:", x, y)

	// Multiple assignment

	a, c := 1, 2

	fmt.Println(a, c)

	a, c = c, a

	fmt.Println("Swapped:", a, c)

	// =========================
	// ARITHMETIC OPERATORS
	// =========================

	num1 := 10
	num2 := 3

	fmt.Println("Add:", num1+num2)
	fmt.Println("Sub:", num1-num2)
	fmt.Println("Mul:", num1*num2)
	fmt.Println("Div:", num1/num2)
	fmt.Println("Mod:", num1%num2)

	// =========================
	// COMPARISON OPERATORS
	// =========================

	fmt.Println("==", num1 == num2)
	fmt.Println("!=", num1 != num2)
	fmt.Println(">", num1 > num2)
	fmt.Println("<", num1 < num2)
	fmt.Println(">=", num1 >= num2)
	fmt.Println("<=", num1 <= num2)

	// =========================
	// LOGICAL OPERATORS
	// =========================

	p := true
	q := false

	fmt.Println("AND:", p && q)
	fmt.Println("OR:", p || q)
	fmt.Println("NOT:", !p)

	// =========================
	// INCREMENT / DECREMENT
	// =========================

	counter := 5

	counter++
	fmt.Println(counter)

	counter--
	fmt.Println(counter)

	// =========================
	// INPUT
	// =========================

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")

	userName, _ := reader.ReadString('\n')

	fmt.Println("Hello", userName)

	// Integer Input

	var userAge int

	fmt.Print("Enter age: ")
	fmt.Scan(&userAge)

	fmt.Println("Your age is", userAge)

	// =========================
	// TYPE CONVERSION
	// =========================

	num := 10

	converted := float64(num)

	fmt.Println(converted)

	// =========================
	// ARRAYS
	// =========================

	arr := [3]int{10, 20, 30}

	fmt.Println(arr)

	// =========================
	// SLICES
	// =========================

	slice := []int{1, 2, 3}

	slice = append(slice, 4)

	fmt.Println(slice)

	// =========================
	// MAPS
	// =========================

	marks := map[string]int{
		"Math":    90,
		"Science": 95,
	}

	fmt.Println(marks)

	// =========================
	// IF ELSE
	// =========================

	if userAge >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	// =========================
	// FOR LOOP
	// =========================

	for i := 1; i <= 5; i++ {
		fmt.Println("Loop:", i)
	}

	// While style

	n := 1

	for n <= 3 {
		fmt.Println(n)
		n++
	}

	// =========================
	// SWITCH
	// =========================

	day := 2

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Unknown")
	}
}

// go mod init go-basics
// go run example.go
