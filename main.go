package main

/*
	Hello, this is a comment multiple line :)
*/
import (
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

	age = 26
	switch age {
	case 10:
		fmt.Println("Child")
		break
	case 26:
		fmt.Println("Adult")
		break
	default:
		fmt.Println("Non human")
		break
	}
}
