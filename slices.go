package main
import(
  "fmt"
  "sort"
)

func main(){
  var colors =[]string{"red","green","blue"}
  fmt.Println(colors)
  colors = append(colors,"purple")
  fmt.Println(colors)
  colors= append(colors[1:len(colors)])
  fmt.Println(colors)
  colors= append(colors[:len(colors)-1])
  fmt.Println(colors)

  numbers := make([]int, 5, 5)
  numbers[0] = 5
  numbers[1] = 58
  numbers[2] = 1
  numbers[3] = 27
  numbers[4] = 9
  fmt.Println(numbers)

  numbers= append(numbers, 235)
  fmt.Println(numbers)
  fmt.Println(cap(numbers))
  sort.Ints(numbers)
  fmt.Println(numbers)
}
