package main
import "fmt"

func test() int {
    return 1
}

// funcs cannot be compared. This program will not compile if uncommented.

func Test_comparision() {
    fmt.Println("")
//    fmt.Println(test == test)
}
