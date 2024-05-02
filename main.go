package main

func main() {

	// Strings, Runes and Bytes

	// var myString = []rune("résumé")
	// var indexed = myString[1]
	// fmt.Printf("%v, %T\n", indexed, indexed)
	// for i, v := range myString {
	// 	fmt.Println(i, v)
	// }
	// fmt.Printf("\nThe length of 'myString' is %v", len(myString))

	// var myRune = 'a'
	// fmt.Printf("\nmyRune = %v", myRune)

	// var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	// var strBuilder strings.Builder
	// for i := range strSlice {
	// 	strBuilder.WriteString(strSlice[i])
	// }
	// var catStr = strBuilder.String()
	// fmt.Printf("\n%v", catStr)

	// Arrays, Slices, Maps, Loops

	// intArr := [...]int32{1, 2, 3}
	// fmt.Println(intArr)

	// var myMap map[string]uint8 = make(map[string]uint8)
	// fmt.Printf("My Map:%v", myMap)

	// var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 26}
	// fmt.Println(myMap2["Adam"])
	// var age, ok = myMap2["John"]

	// if ok {
	// 	fmt.Printf("The age is %v", age)

	// } else {
	// 	fmt.Println("Invalid Name")
	// }

	// for name, age := range myMap2 {
	// 	fmt.Printf("Name: %v, Age:%v \n", name, age)
	// }

	// Variables and Basic Types

	// var printValue string = "hello world!"
	// printMe(printValue)

	// var numerator int = 11
	// var denominator int = 2
	// var result, remainder, err = intDivision(numerator, denominator)

	// switch {
	// case err != nil:
	// 	fmt.Println(err.Error())
	// case remainder == 0:
	// 	fmt.Printf("The result is %v", result)
	// default:
	// 	fmt.Printf("The result is %v with remainder %v", result, remainder)
	// }

	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("The result is %v", result)

	// } else {

	// 	fmt.Printf("The result is %v with remainder %v", result, remainder)
	// }
}

// func printMe(printValue string) {
// 	fmt.Println(printValue)
// }

// func intDivision(numerator int, denominator int) (int, int, error) {
// 	var err error
// 	if denominator == 0 {
// 		err = errors.New("cannot divide by zero")
// 		return 0, 0, err
// 	}

// 	var result int = numerator / denominator
// 	var remainder int = numerator % denominator
// 	return result, remainder, err
// }
