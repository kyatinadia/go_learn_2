package main

import "fmt"

func main() {
	//array modify through index
	var animals = [3]string{"kucing", "burung", "kelinci"}
	animals[0] = "cat"
	animals[1] = "bird"
	animals[2] = "rabbit"
	fmt.Printf("%#v \n", animals)

	//array loop through element
	for key, animal := range animals {
		fmt.Printf("Cara pertama index: %d, value: %s\n", key, animal)
	}

	for i := 0; i < len(animals); i++ {
		fmt.Printf("Cara kedua, index: %d, value: %s\n", i, animals[i])
	}

	for key := range animals {
		fmt.Printf("Cara ketiga, index %d, value: %s\n", key, animals[key])
	}

	//multidimensional array
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}
	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}
