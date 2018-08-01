package main
import(
  "fmt"
)

func main(){
  var colors[3]string
  colors[0]="Red"
  colors[1]="Green"
  colors[2]="Blue"
  fmt.Println(colors)
  fmt.Println(colors[1])

  var numbers =[5]int{5,3,4,1,2}
  fmt.Println(numbers)
  fmt.Println("number of colors:", len(colors))
  fmt.Println("number of numbers:", len(numbers))

  for i:=range colors{
    fmt.Println(colors[i])
    if colors[i]=="Green"{
      break
    }else if colors[i]=="Red"{
      goto endofprogram
    }
  }
  endofprogram :fmt.Println("End of program")

}
