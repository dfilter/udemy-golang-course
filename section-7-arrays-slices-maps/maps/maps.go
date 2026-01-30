package maps

import "fmt"

/*
Map Notes:
- https://go.dev/blog/maps
- Maps key can be of any comparable type https://go.dev/ref/spec#Comparison_operators
- Maps allow for dynamic access of values using the key. types must be defined at compile
time so those aren't adequate for the use case of dynamic assignment and access of values
in some kind of data structure.
*/

func main() {
	websites := map[string]string{ // map[the type of the key]the type of the value
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)

	fmt.Println(websites["Amazon Web Services"]) // Accessing the value in the map

	websites["LinkedIn"] = "https://linkedin.com" // Adding a new value to the map
	fmt.Println(websites)                         // LinkedIn will now appear in the map

	delete(websites, "Google") // Delete Google key value pair from the map
	fmt.Println(websites)      // Google will no longer appear in the map
}
