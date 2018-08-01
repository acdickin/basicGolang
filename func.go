package main
import(
  "fmt"
)
func main(){
  doSomething()

  sum := addValues(23,25)
  fmt.Println("Sum:",sum)

  sum = addAllValues(43,52,11)
  fmt.Println("Sum:",sum)

  name, length := FullName("Andrew", "Dickinson")
  fmt.Printf("Full Name: %v, number of character:%v\n\n", name, length)

  name, length = FullNameNakedReturn("Aurthor", "Dent")
  fmt.Printf("Full Name: %v, number of character:%v\n\n", name, length)
}

func FullName(f,l string)(string, int){
  full := f+" "+l
  length := len(full)
  return full, length
}

func FullNameNakedReturn(f,l string)(full string,length int){
  full = f+" "+l
  length = len(full)
  return
}

//no overloading
func doSomething(){
  fmt.Println("Do something")
}
func addValues(value1, value2 int) int{
  return value1 + value2
}

func addAllValues(values ...int)int{
    sum:=0
    for i:= range values{
      sum+=values[i]
    }
    fmt.Printf("%T\n", values)
    return sum;
}
