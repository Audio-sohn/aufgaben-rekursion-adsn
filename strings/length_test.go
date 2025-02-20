package strings

import "fmt"

func Example_length() {

	fmt.Println(Length("abcde"))
	fmt.Println(Length("abdksutldo"))
	fmt.Println(Length(""))
	fmt.Println(Length("nn"))

	// Output:
	// 5
	// 10
	// 0
	// 2

}
