package main
import "fmt"

func Test_map() {
	test_map := make(map[string]int)
	test_map["testKey"] = 55
	fmt.Println(test_map)

	empty_map_make := make(map[string]int)
	fmt.Print("Empty Map Make: ")
	fmt.Println(empty_map_make)

	empty_map := map[string]int{}
	fmt.Print("Map: ")
	fmt.Println(empty_map)
}
