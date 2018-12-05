package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 32, 11, 2}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Without index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum: ", sum)

	// Range with map
	emails := map[string]string{"mj": "mayank@fueled.com", "bob": "bob@gmail.com"}

	for k, v := range emails {
		fmt.Println(k, ": ", v)
	}

	for k := range emails {
		fmt.Println("Name: ", k)
	}
}
