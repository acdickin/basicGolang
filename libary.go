package main
import(
  "fmt"
  "stringutil"
)
func main(){
// to call stringutil it needs to be in main src for go projects
  name, length := stringutil.FullName("Andrew", "Dickinson")
  fmt.Printf("Full Name: %v, number of character:%v\n\n", name, length)

  name, length = stringutil.FullNameNakedReturn("Aurthor", "Dent")
  fmt.Printf("Full Name: %v, number of character:%v\n\n", name, length)
}
