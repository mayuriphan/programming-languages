package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	name := bufio.NewScanner(os.Stdin)
	name.Scan()
	fmt.Println("Hello, " + name.Text() + "!")
	var i int = 0
	// this is while style
	for i <= 9 {
		fmt.Println(i)
		i++
	}

	arr := [3]int{1, 2, 3}
	for j := 0; j <= len(arr); j++ {
		fmt.Println(j)
	}

	for range arr {
		// prints ele of arr
	}
}
