package main

import (
	"fmt"
	//"strings"
)

func syntax() {
	//fmt.Println("Hello world")

	// Set a var
	var name = "Foo" // variable type is inferred from assignment
	//var userName string // defining empty variables with types
	//var age int
	fmt.Println(name)

	// concatenate strings
	// Spaces are entered between concatenated values at runtime... annoying
	// fmt.Println("Hello user:", name)

	// Constants
	// const a = 3.14
	// fmt.Println(a)

	// string formatting
	//fmt.Printf("Welcome user: %v\n", name)
	//fmt.Printf("Name type: %T\n", name)

	// var bar string
	//fmt.Print("Enter name: ")
	//fmt.Scan(&bar) // this needs a pointer, that's the & sign
	//fmt.Println(bar)

	//var count uint = 180
	//var otherCount int8 = 127

	//count = count - uint(otherCount) // This converts it to a uint because this
	// is a type mismatch. Or change the instanciation

	// empty array and assignment
	//var my_friends [10]string
	//my_friends[0] = "Tom"
	//my_friends[1] = "Dick"
	//my_friends[2] = "Harry"

	//fmt.Println(my_friends)
	//my_friends[0] = "Lily"
	//fmt.Print(my_friends[0])

	//array of unknown length - slice
	//var friends []string
	//friends = append(friends, "Rosanna Ilaria")
	//friends = append(friends, "Tom Fletcher")
	//friends = append(friends, "Dick Greyson")
	//friends = append(friends, "Harry Potter")
	////fmt.Println(friends)
	//friends[1] = "Lily"
	//fmt.Println(friends)

	////Loops
	//for idx, elem := range friends { // using range is like python's enumerate()
	//    // strings.Fields splits on whitespace
	//    var first_name = strings.Fields(elem)
	//    fmt.Println(idx, first_name)
	//}

	//conditionals

	//for i := 0; i < 10; i++ {
	//    fmt.Println(i)
	//}

	var alive bool = true
	if alive {
		fmt.Println("Alive")
	} else if !alive {
		fmt.Println("He dead, jim")
	} else {
		fmt.Println("Houston, we have a problem")
	}

	//Validation
	//var newName string
	//fmt.Println("Enter name: ")
	//fmt.Scan(&newName)

	//    isValid := len(newName) >= 2 // check length is longer than 2
	//isValidEmail := len(newName) >= 2 && strings.Contains(newName, "@")
	//fmt.Println(isValidEmail)

	city := "London"
	switch city {
	case "New York":
		fmt.Println("New York")
	case " Vancouver":
		fmt.Println("Vancouver")
	case "Hong Kong":
		fmt.Println("Hong Kong")
	default:
		fmt.Println("Other")
	}

	fmt.Println(welcome_user(name))

	// create a map/dict
	var userData = make(map[string]int)
	userData["grade_bio"] = 1
	userData["grade_english"] = 2
	userData["grade_geo"] = 3

	fmt.Print(userData)

	var list_of_maps = make([]map[string]int, 0) // the number is the initial length

	var baz = fmt.Sprintf("%v ages", 5)
	fmt.Println(baz)

}

func welcome_user(user string) string {
	ret := "Welcome:" + user
	return ret
}

type personData struct {
	firstName  string
	lastName   string
	age        int
	is_allowed bool
}
