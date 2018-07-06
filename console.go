package main

import (
	"bufio"
	"fmt"
	"os"
  "strconv"
  "strings"
)

func main() {
	// var s string
	// fmt.Scanln(&s)
	// fmt.Println(s)

	reader := bufio.NewReader(os.stdin)
	fmt.Print("Enter text:")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

  fmt.Print("Enter Numer:")
  str, _ := reader.ReadString('\n')

  //trimspace used so we don't get \n
  // 64 is the int type
  f,err := strconv.Parsefloat(strings.TrimSpace(str),64)
  if err != nil{
    fmt.Println(err)
  }else{
    fmt.Println("Value of number:", f)
  }


}
