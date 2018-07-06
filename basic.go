package main

import (
	"fmt"
)

func main() {
	str1 := "The quick brown fox"
	str2 := "jumps over"
	str3 := "the lazy brown dog"
	aNumber := 42
	isTrue := true

	stringLength, _ := fmt.Println(str1, str2, str3)
	// if err == nil{
	fmt.Println("string length: ", stringLength)
	// }

	fmt.Printf("Value of number: %v\n", aNumber)
	fmt.Printf("Value of number: %v\n", isTrue)
	fmt.Printf("Value of number: %.2f\n", float64(aNumber))
	fmt.Printf("Data types: %T, %T, %T, %T and %T\n",
		str1, str2, str3, aNumber, isTrue)
	myString := fmt.Sprintf("Data types as var: %T, %T, %T, %T and %T\n",
		str1, str2, str3, aNumber, isTrue)
	fmt.Print(myString)
}
