package main

import "fmt"

func main() {
	codes := map[string]string{"en": "English", "hin": "Hindi"}
	fmt.Printf("Map codes: %s\n", codes)
	fmt.Printf("Length of map: %d\n", len(codes))

	fmt.Printf("2nd element of map using key: %s\n", codes["hin"])

	value, found := codes["hin"]
	fmt.Printf("Value: %s, Found: %t\n", value, found)

	value1, found1 := codes["mine"]
	fmt.Printf("Value: %s, Found: %t\n", value1, found1)

	codes["it"] = "Italian"
	fmt.Printf("Map codes: %s\n", codes)

	codes["en"] = "English Language"
	fmt.Printf("Map codes: %s\n", codes)

	for key, value := range codes {
		fmt.Printf("Key: %s and its value: %s\n", key, value)
	}

	// truncate map
	// using make function
	// codes = make(map[string]string)
	// fmt.Printf("Map codes: %s\n", codes)

	// using for loop
	// for key, _ := range codes {
	// 	delete(codes, key)
	// }
	// fmt.Printf("Map codes: %s\n", codes)

	mymap := make(map[string]int)
	fmt.Printf("My map: %v\n", mymap)

	mymap["A"] = 1
	mymap["B"] = 2
	mymap["C"] = 3
	fmt.Printf("My map: %v\n", mymap)

	delete(mymap, "B")
	fmt.Printf("My map: %v\n", mymap)
}
