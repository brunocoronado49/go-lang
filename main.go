package main

/*
	Hello, this is a comment multiple line :)
*/
import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	// This is another comment :D
	fmt.Println("Hello Go!")

	// Variables
	const name string = "Bruce"
	fmt.Printf("Hello %s", name)
	fmt.Println()

	var age int = 26
	fmt.Println(age)
	age = age + 4
	fmt.Println(age)

	fmt.Printf("%s you have %s years old", name, fmt.Sprint(age-4))
	fmt.Println()

	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(age))

	var myFloat float32 = 4.5
	fmt.Println(myFloat)
	fmt.Println(reflect.TypeOf(myFloat))

	fmt.Println(age + int(myFloat))
	fmt.Println(myFloat + float32(age))

	var myBool bool = false
	myBool = true
	fmt.Println(!myBool) // false

	// Variable declarada e inicializada de manera abreviada
	myString := "This is a String"
	fmt.Println(myString)

	const myConstant string = "This is a constant"
	fmt.Println(myConstant)

	// Control de flujo
	newAge := 20
	if newAge > 18 {
		fmt.Println("You are an adult")
	} else if newAge < 0 {
		fmt.Println("You are a baby")
	} else {
		fmt.Println("You are an young man")
	}

	// Data structures

	// Array
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	fmt.Println(myArray)
	fmt.Printf("myArray: %v\n", myArray)
	fmt.Println()

	// Map
	// myMap := map[string]int{"Bruce": 26, "Franco": 23, "Joel": 23}
	myMap := make(map[string]int)
	myMap["Bruce"] = 26
	myMap["Franco"] = 23
	myMap["Joel"] = 23
	fmt.Println(myMap)
	fmt.Println(myMap["Bruce"])

	// List
	myList := list.New()
	myList.PushBack(54)
	myList.PushBack("bruce")
	myList.PushBack(true)
	fmt.Println(myList.Len())
	fmt.Println(myList.Back().Value)
	fmt.Println(myList.Front().Value)
	fmt.Println(myList.Front().Next())

	// Bucles
	for i := 0; i <= 20; i++ {
		fmt.Println(i)
	}

	for i := 0; i <= len(myArray)-1; i++ {
		fmt.Println(myArray[i])
	}

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	for i, value := range myArray {
		fmt.Println(i, value)
	}

	// Functions
	myFunction()
	Grettings()

	// Structure
	myStruct := MyStruct{"Bruce", 26}
	fmt.Println(myStruct)
}

func myFunction() {
	fmt.Println("My function")
}

func Grettings() string {
	return "Hello My Friend"
}

type MyStruct struct {
	name string
	age  int
}
