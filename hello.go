package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("What is your name?: ")
	fmt.Scanln(&name)
	fmt.Printf("Hello %v!\n", name)
	fmt.Print("What is your age?: ")
	fmt.Scanln(&age)
	
	if age >= 18 {
		fmt.Println("Congrats! You're an adult!")
	} else {
        fmt.Println("You're still a kid")
    }
}
