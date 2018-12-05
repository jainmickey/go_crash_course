package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign kv
	emails["mj"] = "mayank@fueled.com"
	emails["bob"] = "bob@gmail.com"
	emails["mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["mj"])
	fmt.Println(len(emails))

	// Delete from map
	delete(emails, "bob")
	fmt.Println(emails)

	// Declare map and add kv
	emailsShort := map[string]string{"mj": "mayank@fueled.com", "bob": "bob@gmail.com"}
	emailsShort["mike"] = "mike@gmail.com"
	fmt.Println(emailsShort)
}
