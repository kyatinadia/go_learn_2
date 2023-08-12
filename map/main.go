package main

import "fmt"

func main() {

	//1st way of initializing map
	var person1 map[string]string //declaration
	person1 = map[string]string{} //initialization

	person1["name"] = "Levi"
	person1["age"] = "27"
	person1["address"] = "Shiganshina"

	fmt.Println("name:", person1["name"])
	fmt.Println("age:", person1["age"])
	fmt.Println("address:", person1["address"])

	//@nd way of initializing map
	var person2 = map[string]string{
		"Nama":    "levi ackerman",
		"age":     "27",
		"address": "shiganshina district",
	}

	fmt.Println("name:", person2["Nama"])
	fmt.Println("age:", person2["age"])
	fmt.Println("address:", person2["address"])

	//combining slice with map
}
