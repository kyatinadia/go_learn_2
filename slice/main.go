package main

import (
	"fmt"
	"strings"
)

func main() {
	var animals1 = make([]string, 3)
	animals1[0] = "cat"
	animals1[1] = "bird"
	animals1[2] = "rabbit"

	var animals2 = make([]string, 3)
	animals2[0] = "illama"
	animals2[1] = "otter"
	animals2[2] = "sloth"

	animals1 = append(animals1, "python", "burung", "kelinci")
	fmt.Printf("%#v\n", animals1)

	//slice append with ellipsis
	animals1 = append(animals1, animals2[1])
	fmt.Printf("%#v\n", animals1)

	//slice with copy func
	copyFunc := copy(animals1, animals2)
	fmt.Println("Animals1 ->", animals1)
	fmt.Println("Animals2 ->", animals2)
	fmt.Println("AnimalsCopy ->", copyFunc)

	//slice with slicing
	fruits1 := []string{"apple", "manggo", "orange", "banana", "grape"}

	var fruits2 = fruits1[1:4]
	fmt.Printf("%#v\n", fruits2)

	//slice with backing array
	var lang1 = []string{"python", "ruby", "java", "go", "Javascript"}
	var lang2 = lang1[2:4] //proses slicing
	lang2[0] = "rust"      // assign lang 2
	fmt.Println("Lang 1 ->", lang1)
	fmt.Println("Lang 2 ->", lang2)

	fmt.Println(strings.Repeat("-", 30))

	//slice with cap
	var newFruits1 = []string{"apple", "mango", "durian", "banana"}
	fmt.Println("Fruits2 cap:", cap(newFruits1))
	fmt.Println("Fruits len:", len(newFruits1))
	fmt.Println(strings.Repeat("-", 20))

	var newFruits2 = newFruits1[0:2]
	fmt.Println("Fruits2 cap:", cap(newFruits2))
	fmt.Println("Fruits2 len:", len(newFruits2))
	fmt.Println(strings.Repeat("-", 20))

	var newFruits3 = newFruits1[1:]
	fmt.Println("Fruits2 cap:", cap(newFruits3))
	fmt.Println("Fruits3 len:", len(newFruits3))
	fmt.Println(strings.Repeat("-", 20))
}
