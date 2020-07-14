package main

import (
	"fmt"

	"github.com/kts006/dict_example/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	dictionary["a"] = "AA"
	fmt.Println(dictionary)

	defintion, err := dictionary.Search("a")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(defintion)
	}

	//Add
	addErr := dictionary.Add("b", "BBB")
	if addErr != nil {
		fmt.Println("ERROR", addErr)
	} else {
		value, _ := dictionary.Search("b")
		fmt.Println("Found", value)
	}
	addErr2 := dictionary.Add("b", "BBB")
	fmt.Println(addErr2)

	//Update
	fmt.Println("UPDATE")
	updateErr := dictionary.Update("a", "ABCDE")
	if updateErr == nil {
		value, _ := dictionary.Search("a")
		fmt.Println("Found", value)
	} else {
		fmt.Println(updateErr)
	}
	updateErr2 := dictionary.Update("c", "CCCCC")
	fmt.Println(updateErr2)

	//Delete
	fmt.Println("DELETE")
	deleteErr := dictionary.Delete("b")
	if deleteErr == nil {
		_, err := dictionary.Search("b")
		fmt.Println(err)
	} else {
		fmt.Println(deleteErr)
	}
	deleteErr2 := dictionary.Delete("c")
	fmt.Println(deleteErr2)
}
