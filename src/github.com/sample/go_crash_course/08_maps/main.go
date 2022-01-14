package main

import "fmt"

func main() {

	/ //Define a map
	// email := make(map[key]value)

	// emails := make(map[string]string)

	/// Assign Key Value

	// emails["Sunny"] = "s@mail.com"

	// emails["Tobi"] = "t@mail.com"

	// fmt.Println(emails)
	// fmt.Println(emails["Sunny"])
	// fmt.Println(len(emails))

	/// Delete from map

	// delete(emails, "Sunny")
	// fmt.Println(emails)

	/// Declare map and add kv
	emails := map[string]string{"Bob": "bob@mail.com", "John": "john@mail.com", "Mike": "mike@mail.com"}
	fmt.Println(emails)

}
